package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Paprec/trucktrack/service"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-zoo/bone"
)

const (
	slash           = "/"
	EndPointURL     = "geturl"
	endPointMetrics = "metrics"
)

func MakeHandler(svc service.MACService) http.Handler {
	r := bone.New()
	// Server -> Master
	r.Get("/list", httptransport.NewServer(
		listEndpoint(svc),
		decodeListRequest,
		encodeListResponse,
	))

	// Master <-> Slave
	// En param : ID - MAC addresse
	r.Get("/author", httptransport.NewServer(
		authorEndpoint(svc),
		decodeAuthorRequest,
		encodeAuthorResponse,
	))

	// Reception JSON : ID - MAC Addresse, Time, I/O,
	r.Post("/activity", httptransport.NewServer(
		activityEndpoint(svc),
		decodeActivityRequest,
		encodeActivityResponse,
	))

	return r
}

func decodeListRequest(_ context.Context, r *http.Request) (interface{}, error) {

	return getMACAddressesRequest{}, nil
}

func encodeListResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func decodeAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := r.URL.Query()
	macAddress := req.Get("ID")
	return getAuthorRequest{ID: macAddress}, nil
}

func encodeAuthorResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "text/plain")
	ack := response.(getAuthorResponse)
	if ack.Authorization != "OK" {
		w.WriteHeader(http.StatusUnauthorized)
	}

	_, err := w.Write([]byte(fmt.Sprintf("%s\n", response)))
	return err

}

func decodeActivityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body, errr := io.ReadAll(r.Body)
	if errr != nil {
		return nil, errr
	}
	defer r.Body.Close()

	message := string(body)
	if message == "" {
		return nil, fmt.Errorf("le message est vide")
	}
	log.Printf("Message reçu : %s\n", message)

	return postActivityRequest{Activity: message}, nil
}

func encodeActivityResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte(fmt.Sprintf("%s\n", response)))
	return err
}
