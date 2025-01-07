package api

import (
	"context"

	"github.com/Paprec/trucktrack/service"
	"github.com/go-kit/kit/endpoint"
)

func makeEndpoint(svc service.MACService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		macs, err := svc.GetMACAddresses(ctx)
		if err != nil {
			return nil, err
		}
		return getMACAddressesResponse{MACAddresses: macs}, nil
	}
}
