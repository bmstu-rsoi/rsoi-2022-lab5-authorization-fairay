package utils

type DBConfiguration struct {
	Type string `json:"type"`
	Name string `json:"name"`

	User     string `json:"user"`
	Password string `json:"password"`

	Host string `json:"host"`
	Port string `json:"port"`
}

type Configuration struct {
	DB      DBConfiguration `json:"db"`
	LogFile string          `json:"log_file"`
	Port    uint16          `json:"port"`
	RawJWKS string          `json:"raw-jwks"`
}

var (
	Config Configuration
)

// TODO: returnable error
func InitConfig() {
	Config = Configuration{
		DBConfiguration{
			"postgres",
			"tickets",
			"program",
			"test",
			"postgres-service",
			"5432",
		},
		"logs/server.log",
		8070,
		`{"keys":[{"kty":"RSA","alg":"RS256","kid":"MeFQFgrQ4H20rObty3HDo2U-mAAD0dPydKrXOJ9zGAc","use":"sig","e":"AQAB","n":"quqU1buEQMDreTIXabUD491R05xrBpTkn5mf9JUtRWjtFp1qj5mQ7fpagYrs0nxbnJtHESbdTnoF1bsUT4qmXnldOC7VrZZr4mW3fhlNjF176yF4mFSjqCcRaj3uELBc2vbpEn-xasS0oyjr-pQ9n5MGQWkHCUzDm1yigunTYqIALnRFLBLTesXWzKyFHggvTeIjgVt-kPDPjn8bzwQrZC4MC0s-gmgHXZnY7wQMCJ33satSzrbe_XikoJsyKEUfeU3SKjd_MVhuvvvWSv9BUJWsgUzxySnBSGxIlydYPqVdLB6YN4sEItRBbLC0_0m3uYyAQpew7IaHda7yQoIW9Q"}]}`,
	}
}
