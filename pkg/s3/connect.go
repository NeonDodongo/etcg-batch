package s3

// ConnInfo holds the connection data required to interface with S3
type ConnInfo struct {
	PublicKey string `json:"publicKey"`
	SecretKey string `json:"secretKey"`
	Endpoint  string `json:"endpoint"`
	Region    string `json:"region"`
	Token     string `json:"token"`
}

// DecodeConnInfo not implemented yet
func DecodeConnInfo(ci ConnInfo) ConnInfo {
	// TODO: Implement decoding encoded secrets (base64 probably)

	return ConnInfo{}
}
