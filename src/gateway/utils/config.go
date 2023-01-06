package utils

type Configuration struct {
	LogFile            string `json:"log_file"`
	Port               uint16 `json:"port"`
	RawJWKS            string `json:"raw-jwks"`
	FlightsEndpoint    string `json:"flights-endpoint"`
	TicketsEndpoint    string `json:"tickets-endpoint"`
	PrivilegesEndpoint string `json:"privileges-endpoint"`
}

var (
	Config Configuration
)

// TODO: returnable error
func InitConfig() {
	Config = Configuration{
		"logs/server.log",
		8080,
		`{"keys":[{"kty":"RSA","alg":"RS256","kid":"MeFQFgrQ4H20rObty3HDo2U-mAAD0dPydKrXOJ9zGAc","use":"sig","e":"AQAB","n":"quqU1buEQMDreTIXabUD491R05xrBpTkn5mf9JUtRWjtFp1qj5mQ7fpagYrs0nxbnJtHESbdTnoF1bsUT4qmXnldOC7VrZZr4mW3fhlNjF176yF4mFSjqCcRaj3uELBc2vbpEn-xasS0oyjr-pQ9n5MGQWkHCUzDm1yigunTYqIALnRFLBLTesXWzKyFHggvTeIjgVt-kPDPjn8bzwQrZC4MC0s-gmgHXZnY7wQMCJ33satSzrbe_XikoJsyKEUfeU3SKjd_MVhuvvvWSv9BUJWsgUzxySnBSGxIlydYPqVdLB6YN4sEItRBbLC0_0m3uYyAQpew7IaHda7yQoIW9Q"}]}`,

		"http://flights-service:8060",
		"http://tickets-service:8070",
		"http://privileges-service:8050",
	}
}
