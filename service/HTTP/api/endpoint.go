package api

import (
	"context"
	"strings"

	"github.com/Paprec/trucktrack/service"
	"github.com/go-kit/kit/endpoint"
)

const (
	ack     = "OK"
	nack    = "Not OK"
	empty   = ""
	errNack = "unauthorized vehicule"
)

func listEndpoint(svc service.MACService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		macs := svc.GetMACAddresses(service.Macs)
		if len(macs) == 0 {
			return nil, service.ErrUnknownMethod
		}
		return getMACAddressesResponse{MACAddresses: macs}, nil
	}
}

func authorEndpoint(svc service.MACService) endpoint.Endpoint {

	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getAuthorRequest)
		id := svc.AuthorId(req.ID)

		switch id != empty {
		case true:
			return getAuthorResponse{ack}, nil
		case false:
			return getAuthorResponse{nack}, nil
		}
		return getAuthorResponse{empty}, nil
	}

}

func activityEndpoint(svc service.MACService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postActivityRequest)
		if strings.Contains(req.Activity, "attend") {
			addr := req.Activity[10:27]

			id := svc.AuthorId(addr)

			if id != empty {
				return postActivityResponse{Response: "Ouvrir"}, nil
			}
			return postActivityResponse{Response: "Fermer"}, nil
		}

		return postActivityResponse{Response: "OK"}, nil

	}
}
