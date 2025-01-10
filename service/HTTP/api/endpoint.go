package api

import (
	"context"

	"github.com/Paprec/trucktrack/service"
	"github.com/go-kit/kit/endpoint"
)

func makeEndpoint(svc service.MACService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		macs := svc.GetMACAddresses(service.Macs)
		if len(macs) == 0 {
			return nil, service.ErrUnknownMethod
		}
		return getMACAddressesResponse{MACAddresses: macs}, nil
	}
}

func newAddrEndpoint(svc service.MACService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		addmac := svc.AddMACAddresses(service.AddMac)

		return AddMACAddressesResponse{addmac}, nil
	}
}
