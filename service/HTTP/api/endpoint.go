package api

import (
	"context"

	"github.com/Paprec/trucktrack/service"
	"github.com/go-kit/kit/endpoint"
)

const (
	ack     = "OK"
	nack    = "Not OK"
	empty   = ""
	errNack = "unauthorized vehicule"
)

// func listEndpoint(svc service.MACService) endpoint.Endpoint {
// 	return func(_ context.Context, request interface{}) (interface{}, error) {

// 		macs := svc.GetMACAddresses(service.Macs)
// 		if len(macs) == 0 {
// 			return nil, service.ErrUnknownMethod
// 		}
// 		return getMACAddressesResponse{MACAddresses: macs}, nil
// 	}
// }

func authorEndpoint(svc service.MACService) endpoint.Endpoint {

	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getAuthorRequest)
		id := svc.AuthorId(req.ID)

		// if err != nil {
		// 	return nil, service.ErrAuthreq
		// }

		switch id != empty {
		case true:
			return getAuthorResponse{Authorization: ack, Error: empty}, nil
		case false:
			return getAuthorResponse{Authorization: nack, Error: errNack}, nil
		}
		return getAuthorResponse{Authorization: empty, Error: empty}, nil
	}

}

// func activityEndpoint(svc service.MACService) endpoint.Endpoint {
// 	return func(_ context.Context, request interface{}) (interface{}, error) {
// 		body, err := io.ReadAll() // Mettre parametre pour ReadAll
// 		if err != nil {
// 			return nil, service.ErrBodyRead
// 		}
// 		log.Println(body)

// 	}
// }
