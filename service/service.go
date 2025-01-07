package service

import (
	"context"
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
)

type MacAddressList struct {
	Addresses []string `json:"addresses"`
}

type MACService interface {
	GetMACAddresses(ctx context.Context) ([]string, error)
}

type macService struct{}

func NewService() MACService {
	return &macService{}
}

func (macService) GetMACAddresses(ctx context.Context) ([]string, error) {
	// Pour cet exemple, on retourne une liste d'adresses MAC statiques
	macs := []string{
		"01:01:01:01:01:01",
		"01:01:01:01:01:02",
		"01:01:01:01:01:03",
	}
	return macs, nil
}
