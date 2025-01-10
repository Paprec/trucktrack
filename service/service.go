package service

import (
	"errors"
)

const (
	invalidWriting = "l'écriture s'est mal passée"
	validWriting   = "l'écriture s'est bien passée"
)

var (
	errInvalidWriting error = errors.New(invalidWriting)
	ErrUnknownMethod  error = errors.New("methode non traitée")
	Macs                    = []string{
		"01:01:01:01:01:01",
		"01:01:01:01:01:02",
		"01:01:01:01:01:03",
	}
	AddMac string = "Adresse ajouté"
)

type MacAddressList struct {
	Addresses []string `json:"addresses"`
}

type MACService interface {
	GetMACAddresses(mac []string) []string

	AddMACAddresses(addmac string) string
}

type macService struct{}

func NewService() MACService {
	return &macService{}
}

func (macService) GetMACAddresses(macs []string) []string {

	return macs
}

func (macService) AddMACAddresses(addmac string) string {

	return addmac
}
