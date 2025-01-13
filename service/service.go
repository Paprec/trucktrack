package service

import (
	"errors"
)

const (
	invalidWriting = "l'écriture s'est mal passée"
	invalidReading = "la lecture s'est mal passée"
	validWriting   = "l'écriture s'est bien passée"
	Ack            = "OK"
	errorBodyRead  = "impossible de lire le body"
	errorAuthReq   = "impossible de determiner l'autorisation"
)

var (
	errInvalidWriting error = errors.New(invalidWriting)
	errInvalidReading error = errors.New(invalidReading)
	ErrUnknownMethod  error = errors.New("methode non traitée")
	ErrBodyRead       error = errors.New(errorBodyRead)
	ErrAuthreq        error = errors.New(errorAuthReq)
	Macs                    = [3]string{
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

	AuthorId(mac string) string
}

type macService struct{}

func NewService() MACService {
	return &macService{}
}

func (macService) GetMACAddresses(macs []string) []string {

	return macs
}

func (macService) AuthorId(mac string) string {

	for i := 0; i < len(Macs); i++ {
		switch mac != Macs[i] {
		case true:
			return ""
		case false:
			return mac
		}
	}
	return ""
}
