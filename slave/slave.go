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
	message      = "Camion"
)

func main() {

	cmd := exec.Command(commandName, commandArg, commandValue)

	retour, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la commande type: ", err)
	}

	chanOut := make(chan string)
	errs := make(chan error)

	go listenIO(chanOut, cmd)

	go getIO(retour, chanOut)

	go getRequest("list")

	go postRequest("activity", message)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err = <-errs
	log.Println(fmt.Sprintf("service terminated: %s", err))
	startAddr := regexp.MustCompile(strenghtMAC)
	MACAddr := regexp.MustCompile(formatMAC)

	for adresse := range chanOut {

		addr := MACAddr.FindString(adresse)
		start := startAddr.FindString(adresse)

		if addr != "" && start != "" {
			fmt.Println(addr, start)
		}
	}
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

func getRequest(addrend string) {

	req := link + addrend
	response, errors := http.Get(req)
	if errors != nil {
		log.Println("Error Get method")
	}
	body, errors := io.ReadAll(response.Body)
	if errors != nil {
		log.Println("Error Read Body")
	}
	response.Body.Close()
	resp := fmt.Sprintf("%s", body)
	log.Printf(resp)

}

func postRequest(addrend string, message string) {

	req := link + addrend
	body := bytes.NewBuffer([]byte(message))
	response, errors := http.Post(req, "application/json", body)
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
