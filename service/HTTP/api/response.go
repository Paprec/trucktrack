package api

// type getMACAddressesResponse struct {
// 	MACAddresses []string `json:"mac_addresses"`
// }

type getAuthorResponse struct {
	Authorization string `json:"ACK"`
	Error         string `json:"error"`
}

type postActivityResponse struct {
	Response string `json:"Response"`
}
