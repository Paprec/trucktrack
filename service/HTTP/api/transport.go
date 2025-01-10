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
	r.Get("/list", httptransport.NewServer(
		makeEndpoint(svc),
		decodeRequest,
		encodeResponse,
	))

	// Master <-> Slave
	// En param : ID - MAC addresse
	r.Get("/author", httptransport.NewServer(
		makeEndpoint(svc),
		decodeRequest,
		encodeResponse,
	))

	// Reception JSON : ID - MAC Addresse, Time, I/O,
	r.Post("/activity", httptransport.NewServer(
		makeEndpoint(svc),
		decodeRequest,
		encodeResponse,
	))

	return r
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {

	return getMACAddressesRequest{}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
