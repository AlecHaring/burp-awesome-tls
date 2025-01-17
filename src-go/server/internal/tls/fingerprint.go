package tls

import (
	"strings"

	utls "github.com/refraction-networking/utls"
)

type Fingerprint string

func (fingerprint Fingerprint) ToClientHelloId() *utls.ClientHelloID {
	if fingerprint == "Default" {
		return nil
	}

	parts := strings.Split(string(fingerprint), " ")

	clientHelloID := &utls.ClientHelloID{}
	if len(parts) >= 1 {
		clientHelloID.Client = parts[0]
	}
	if len(parts) >= 2 {
		clientHelloID.Version = parts[1]
	}

	return clientHelloID
}
