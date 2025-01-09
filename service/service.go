package service

import (
	"errors"
)

const (
	invalidWriting = "l'écriture s'est mal passée"
	validWriting   = "l'écriture s'est bien passée"
	MethodGet      = "GET"
	MethodPost     = "POST"
	chmod          = 0770
)

var (
	errInvalidWriting error = errors.New(invalidWriting)
	ErrUnknownMethod  error = errors.New("methode non traitée")
	Macs                    = []string{
		"01:01:01:01:01:01",
		"01:01:01:01:01:02",
		"01:01:01:01:01:03",
	}
)

type MacAddressList struct {
	Addresses []string `json:"addresses"`
}

type MACService interface {
	GetMACAddresses(mac []string) []string
}

type macService struct{}

func NewService() MACService {
	return &macService{}
}

func (macService) GetMACAddresses(macs []string) []string {
	// Pour cet exemple, on retourne une liste d'adresses MAC statiques

	return macs
}
