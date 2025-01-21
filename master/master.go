package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strconv"
	"syscall"

	"github.com/Paprec/trucktrack/service"
	"github.com/Paprec/trucktrack/service/HTTP/api"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	logg "github.com/go-kit/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const (
	commandName  = "bluetoothctl"
	commandArg   = "scan"
	commandValue = "on"
	port         = "9090"
	strenghtMAC  = `RSSI: -([0-9]{2})`
	formatMAC    = `([0-9A-Fa-f]{2}(:[0-9A-Fa-f]{2}){5})`
	link         = "http://localhost:9090/"
	author       = "author"
	activity     = "activity"
	authorForm   = "?ID="
	in           = "entre"
	out          = "sort"
	authorIn     = "attend a l'entree"
	authorOut    = "attend a la sortie"
)

var (
	addr    string
	rssi    int
	errconv error
	res     []byte
)

func main() {

	logger := logg.NewLogfmtLogger(os.Stderr)
	svc := newService(logger)

	cmd := exec.Command(commandName, commandArg, commandValue)

	retour, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la commande type: ", err)
	}

	chanOut := make(chan string)
	errs := make(chan error)

	go listenIO(chanOut, cmd)

	go getIO(retour, chanOut)

	go getAddrRssi(chanOut)

	go startHTTPServer(api.MakeHandler(svc), port, errs)

	go getRequest("list")

	rssiCase()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err = <-errs
	log.Printf("%s\n", fmt.Sprintf("service terminated: %s", err))
}

func newService(logger logg.Logger) service.MACService {
	svc := service.NewService()
	svc = api.LoggingMiddleware(svc, logger)
	svc = api.MetricsMiddleware(
		svc,
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "service",
			Subsystem: "api",
			Name:      "request_count",
			Help:      "Number of requests received",
		}, []string{"method"}),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "service",
			Subsystem: "api",
			Name:      "request_latency_in_microseconds",
			Help:      "Total duration of requests in microseconds",
		}, []string{"method"}),
	)
	return svc
}

func listenIO(chanOut chan string, cmd *exec.Cmd) {

	chanError := cmd.Start()
	if chanError != nil {
		log.Printf("Error: %s", chanError)
	}

	chanError = cmd.Wait()
	if chanError != nil {
		log.Printf("Error: %s", chanError)
	}

	close(chanOut)

}

func getIO(retour io.ReadCloser, chanOut chan string) {
	scanner := bufio.NewScanner(retour)

	for scanner.Scan() {
		chanOut <- scanner.Text()
	}

	channError := scanner.Err()
	if channError != nil {
		log.Printf("Error: %s", channError)
	}
}

func startHTTPServer(handler http.Handler, port string, errs chan error) {

	//p := fmt.Sprintf(":%s", port)

	log.Printf(fmt.Sprintf("Service started using http on port: %s", port))

	errs <- http.ListenAndServe(port, handler)

	if errs != nil {
		fmt.Println("Erreur lors de l'exécution de la commande type: ", errs)
	}
}

func getRequest(addrend string) []byte {

	req := link + addrend
	response, errors := http.Get(req)
	if errors != nil {
		log.Println("Error comm HTTP")
	}
	body, errors := io.ReadAll(response.Body)
	if errors != nil {
		log.Println("Error Read Body")
	}
	response.Body.Close()
	resp := fmt.Sprintf("%s", body)
	log.Printf(resp)

	return body

}

func getAddrRssi(chanOut chan string) {
	rssiAddr := regexp.MustCompile(strenghtMAC)
	MACAddr := regexp.MustCompile(formatMAC)

	for adresse := range chanOut {

		addr = MACAddr.FindString(adresse)
		rssi, errconv = strconv.Atoi(rssiAddr.FindString(adresse))
		if errconv != nil {
			fmt.Println("Erreur lors de la conversion string to int: ", errconv)
		}
		if addr != "" && rssi != 0 {
			fmt.Println(addr, rssi)
		}
	}
}

func postRequest(id string) {

	reqIn := link + activity + "Le camion" + id + in
	// reqOut := link + activity + "Le camion" + id + out
	// reqAuthorIn := link + activity + "Le camion" + id + authorIn
	// reqAuthorOut := link + activity + "Le camion" + id + authorOut

	body := bytes.NewBuffer([]byte(id))
	response, errors := http.Post(reqIn, "application/json", body)
	if errors != nil {
		log.Println("Error POST method")
	}

	res, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error Read Body")
	}

	response.Body.Close()
	resp := string(res)
	log.Println(resp)

}

func rssiCase() {
	if rssi > -70 {
		res = getRequest(addr)
		log.Println("res:", string(res))

		if string(res) != "OK" {
			log.Println("Barrière fermée")
		}
	}

	if rssi > -40 && string(res) == "OK" {
		postRequest(addr)
		log.Println("Barrière ouverte")
	}
}
