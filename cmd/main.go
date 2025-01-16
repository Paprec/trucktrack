package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	authorForm   = "?ID="
)

type Addresses struct {
	MacAddresses []string `json:"mac_addresses"`
}

func main() {
	//var test Addresses
	logger := logg.NewLogfmtLogger(os.Stderr)
	svc := newService(logger)

	// cmd := exec.Command(commandName, commandArg, commandValue)

	// retour, err := cmd.StdoutPipe()
	// if err != nil {
	// 	fmt.Println("Erreur lors de l'exécution de la commande type: ", err)
	// }

	// chanOut := make(chan string)
	errs := make(chan error)

	// go listenIO(chanOut, cmd)

	// go getIO(retour, chanOut)

	go startHTTPServer(api.MakeHandler(svc), port, errs)

	res := getRequest("list")
	log.Println("res:", string(res))
	if string(res) == "OK" {
		log.Println("Barriere UP")
	}
	// errjsonparse := json.Unmarshal(res, &test)

	// if errjsonparse != nil {
	// 	log.Println("Error JSON Parse")
	// }

	// log.Printf("JSON Parse -> %v\n", test)

	// startAddr := regexp.MustCompile(strenghtMAC)
	// MACAddr := regexp.MustCompile(formatMAC)

	// for adresse := range chanOut {

	// 	addr := MACAddr.FindString(adresse)
	// 	start := startAddr.FindString(adresse)

	// 	if addr != "" && start != "" {
	// 		fmt.Println(addr, start)
	// 	}
	// }

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err := <-errs
	log.Printf("%s", fmt.Sprintf("service terminated: %s", err))
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

// func listenIO(chanOut chan string, cmd *exec.Cmd) {

// 	chanError := cmd.Start()
// 	if chanError != nil {
// 		log.Printf("Error: %s", chanError)
// 	}

// 	chanError = cmd.Wait()
// 	if chanError != nil {
// 		log.Printf("Error: %s", chanError)
// 	}

// 	close(chanOut)

// }

// func getIO(retour io.ReadCloser, chanOut chan string) {
// 	scanner := bufio.NewScanner(retour)

// 	for scanner.Scan() {
// 		chanOut <- scanner.Text()
// 	}

// 	channError := scanner.Err()
// 	if channError != nil {
// 		log.Printf("Error: %s", channError)
// 	}
// }

func startHTTPServer(handler http.Handler, port string, errs chan error) {

	p := fmt.Sprintf(":%s", port)

	log.Printf(fmt.Sprintf("Service started using http on port: %s", port))

	errs <- http.ListenAndServe(p, handler)
	log.Println("errs:", errs)
	if errs != nil {
		fmt.Println("Erreur lors de l'exécution de la commande type: ", errs)
	}

}

// func getRequest(addrend string) []byte {

// 	req := link + addrend
// 	response, errors := http.Get(req)
// 	if errors != nil {
// 		log.Println("Error comm HTTP")
// 	}
// 	body, errors := io.ReadAll(response.Body)
// 	if errors != nil {
// 		log.Println("Error Read Body")
// 	}
// 	response.Body.Close()
// 	// resp := fmt.Sprintf("%s", body)
// 	// log.Printf(resp)
// 	return body
// }

func getRequest(id string) []byte {

	req := link + author + authorForm + "01:01:01:01:01:01"
	response, errors := http.Get(req)
	if errors != nil {
		log.Println("Error Get method")
	}
	body, errors := io.ReadAll(response.Body)
	if errors != nil {
		log.Println("Error Read Body")
	}
	response.Body.Close()

	return body
}
