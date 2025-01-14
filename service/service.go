package service

import (
	"errors"
	"fmt"
	"log"
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

	AuthorId(mac string) string

	PostActivity(message string) string
}

type macService struct{}

func NewService() MACService {
	return &macService{}
}

func (macService) GetMACAddresses(macs []string) []string {

	return macs
}

func (macService) AuthorId(mac string) string {
	macs := ""

	for i := 0; i < len(Macs); i++ {
		if mac == Macs[i] {
			return mac
		}
		log.Println("Valeur de Macs", Macs[i])
	}
	return macs
}

func (macService) PostActivity(message string) string {
	mesg := fmt.Sprintf("Message reçu : %s", message)
	log.Println("mesg:", mesg)
	return mesg
}
