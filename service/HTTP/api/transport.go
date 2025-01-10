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
	r.Get("/list", httptransport.NewServer(
		makeEndpoint(svc),
		decodeRequest,
		encodeResponse,
	))

	r.Post("/newaddr", httptransport.NewServer(
		newAddrEndpoint(svc),
		decodeNewAddr,
		encodeNewAddr,
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

func decodeNewAddr(_ context.Context, r *http.Request) (interface{}, error) {

	return AddMACAddressesRequest{}, nil
}

func encodeNewAddr(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "Content-Type")
	return json.NewEncoder(w).Encode(response)
}
