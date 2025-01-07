package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Paprec/trucktrack/service"
	httptransport "github.com/go-kit/kit/transport/http"
)

const (
	slash           = "/"
	EndPointURL     = "geturl"
	endPointMetrics = "metrics"
)

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return getMACAddressesRequest{}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func MakeHandler(svc service.MACService) http.Handler {
	return httptransport.NewServer(
		makeEndpoint(svc),
		decodeRequest,
		encodeResponse,
	)
}
