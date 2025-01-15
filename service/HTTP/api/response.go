package api

type getMACAddressesResponse struct {
	MACAddresses []string `json:"mac_addresses"`
}

type getAuthorResponse struct {
	Authorization string
}

type postActivityResponse struct {
	Response string `json:"Response"`
}
