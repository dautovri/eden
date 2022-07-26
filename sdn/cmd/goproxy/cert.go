package main

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/elazarl/goproxy"
)

var caCert = []byte(`-----BEGIN CERTIFICATE-----
MIIFIjCCAwqgAwIBAgIILnV2rF6rRj8wDQYJKoZIhvcNAQEMBQAwLzELMAkGA1UE
BhMCVVMxEDAOBgNVBAoTB2xmLWVkZ2UxDjAMBgNVBAMTBXByb3h5MB4XDTIyMDcw
ODE0MTY0MloXDTMyMDcwNTE0MTY0MlowLzELMAkGA1UEBhMCVVMxEDAOBgNVBAoT
B2xmLWVkZ2UxDjAMBgNVBAMTBXByb3h5MIICIjANBgkqhkiG9w0BAQEFAAOCAg8A
MIICCgKCAgEAmz4kI8FwvqQKZ+bcXB9Elme3B1hG6fo0gU7Ej1JpR0grfkiea1Kn
g06RiGYjUgl5zQ3MmyE9FQs6SSqbohoWfZv5FabnbqWYy6zjHz4cNeFvfV5kHfe4
UnNUNwLYAYni1InP3iqdVKhCKHS6+5FjvB8iwN5SesBf6yqHKli8+Lm54YIqZRFx
9yMJyM3qCquuqiQJiKibx+76UIUWuf9Whf64p1NLaAlpbq3tbNmzV32BCzn8Otf9
Mv+wnGvxzQDsPRTfgBskptsPF2K28932iSLMudZTnuXfl6ydaHpYNK6SI/3GmkJa
ViZzIGNsjAz31QTqd/06VTAVL3597fSIwBnXSG3NryjKe1qhulk+7hhXiVui32c0
jvkwgTSrWb1FGuzkgUqWXfdUIiDTT/0rDBIbRq23rlonYOnMPJI9G4PQwmOZTFXi
kLg8qq28pz/jHTpn8VqKF/XMDUxf/0EFc/vejk8cgzDDAlqkqkvEGee483PyuPi2
etBX9/+ngFoYDSZqCnPgAShYqq9qroIjtg7/cbdr8KiMO4Dj6yHM2OzTXQ6THjCq
V0As2i64YGDDaMhsavvwB2geznjBXf/extQiVshLEm08RQ8HViRzm0G1p+7nHQDb
VsY25yg9ETnCRMukiBzUXF6uV8z3JauSj1eIZGqgL0wQ+bIpcX1BaNkCAwEAAaNC
MEAwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAQYwHQYDVR0OBBYEFGVF
4X8j3OGihtnFiC3BZ33zqvfjMA0GCSqGSIb3DQEBDAUAA4ICAQATNvl+IgAwz3Gx
i+6WHiqsVwRWKMifZKY952lcusviq0m6Aa/48ifZ4fc7nOTJ/pEXHDJKF/0ObYYH
83j8AenkAp5lHEHZXfX9138QEhmMaBFqSS7IH1Vt8rvr3ZUJdq9rRNLLbnmnewdK
C+YKwFyuqbdjdPMRQJiBWb2WyiWLydn/fMvU8Tsdcriyn4bcJdu2+4iPsb4e1YvA
/ubGBS5Yt/v7iulNGEu9jp5nxfaBaQrc63HXcvyM4f9Q4kMNuOe+hJ/fDtUVSX67
GTfXQGwEWUUaT4nI8mbk55KDdCy1Li4Ky5YweVR7rPeBLZ7Z+LW6pN51JYUzD/LQ
YPhwM/mgEE+cum5jlMOBKkDwqrLpwerThfWJ2Ry+eCeLPb6qxjO7jWZcsL9yhQFX
yBiq6F+zUKN6kRp4kOKvEM52aN0lpY4WY9xvRfehog1NS6YADQqm2zNdiwzu1RGg
IwLzYJF6lJfTS8vPK1DNeZS9rvchc+v1ABfaLlyo4eQ2xtl6T4ynDDuRnFGppq7u
5ZK/cM21U+CmMLs3l1yAuXoCoD+XT1P5kzJPtyjKaImSvNJHA9nKiWamrwmeTNPS
0J7/B8Zv4EKn8mYTe4Okzn2GmOUF8djEyxWuOyfPdROJG5/oNhOJQtOd16xhnX+4
i5dPDNhKEZtP3KeY2vQRoldysRSbFA==
-----END CERTIFICATE-----`)

var caKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`)

func setCA(caCert, caKey []byte) error {
	goproxyCa, err := tls.X509KeyPair(caCert, caKey)
	if err != nil {
		return err
	}
	if goproxyCa.Leaf, err = x509.ParseCertificate(goproxyCa.Certificate[0]); err != nil {
		return err
	}
	goproxy.GoproxyCa = goproxyCa
	goproxy.OkConnect = &goproxy.ConnectAction{Action: goproxy.ConnectAccept, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.MitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.HTTPMitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectHTTPMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.RejectConnect = &goproxy.ConnectAction{Action: goproxy.ConnectReject, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	return nil
}