package api

type getMACAddressesResponse struct {
	MACAddresses []string `json:"mac_addresses"`
}

type addMACAddressesResponse struct {
	ADDMacAddresses string
}
