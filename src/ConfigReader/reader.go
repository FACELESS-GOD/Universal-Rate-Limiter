package ConfigReader

type RateLimit struct {
	Unit            string `yaml:"unit"`
	RequestsPerUnit int64  `yaml:"requests_per_unit"`
}

type Descriptor struct {
	Key        string    `yaml:"key"`
	Value      string    `yaml:"value"`
	Algorithm  string    `yaml:"algorithm"`
	Rate_Limit RateLimit `yaml:"rate_limit"`
}

type RateLimiterConfiguration struct {
	Domain      string       `yaml:"domain"`
	Descriptors []Descriptor `yaml:"descriptors"`
}
