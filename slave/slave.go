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
)

func main() {

	rssiAddr := regexp.MustCompile(strenghtMAC)
	MACAddr := regexp.MustCompile(formatMAC)
	cmd := exec.Command(commandName, commandArg, commandValue)

	retour, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la commande type: ", err)
	}

	chanOut := make(chan string)
	errs := make(chan error)

	go listenIO(chanOut, cmd)

	go getIO(retour, chanOut)

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

	resp := getRequest(addr)
	if resp != true {
		//Barriere Fermee
	}
	// Barriere ouverte

	if rssi > -40 {
		postRequest(addr)
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err = <-errs
	log.Println(fmt.Sprintf("service terminated: %s", err))

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

func getRequest(id string) bool {

	req := link + author + authorForm + id
	response, errors := http.Get(req)
	if errors != nil {
		log.Println("Error Get method")
	}
	body, errors := io.ReadAll(response.Body)
	if errors != nil {
		log.Println("Error Read Body")
	}
	response.Body.Close()
	resp := string(body)

	return resp == "OK"
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
	resp := fmt.Sprintf("%s", res)
	log.Printf(resp)

}
