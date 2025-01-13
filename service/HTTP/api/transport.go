package api

import (
	"context"
	"encoding/json"
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
	// r.Get("/list", httptransport.NewServer(
	// 	listEndpoint(svc),
	// 	decodeListRequest,
	// 	encodeListResponse,
	// ))

	// Master <-> Slave
	// En param : ID - MAC addresse
	r.Get("/author", httptransport.NewServer(
		authorEndpoint(svc),
		decodeAuthorRequest,
		encodeAuthorResponse,
	))

	// // Reception JSON : ID - MAC Addresse, Time, I/O,
	// r.Post("/activity", httptransport.NewServer(
	// 	activityEndpoint(svc),
	// 	decodeActivityRequest,
	// 	encodeActivityResponse,
	// ))

	return r
}

// func decodeListRequest(_ context.Context, r *http.Request) (interface{}, error) {

// 	return getMACAddressesRequest{}, nil
// }

// func encodeListResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
// 	w.Header().Set("Content-Type", "application/json")
// 	return json.NewEncoder(w).Encode(response)
// }

func decodeAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := r.URL.Query()
	macAddress := req.Get("ID")
	return getAuthorRequest{ID: macAddress}, nil
}

func encodeAuthorResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	ack := response.(getAuthorResponse)
	if ack.Authorization != "OK" {
		w.WriteHeader(http.StatusUnauthorized)
	}
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(ack)

}

// func decodeActivityRequest(_ context.Context, r *http.Request) (interface{}, error) {

// 	return getMACAddressesRequest{}, nil
// }

// func encodeActivityResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
// 	w.Header().Set("Content-Type", "application/json")
// 	return json.NewEncoder(w).Encode(response)
// }
