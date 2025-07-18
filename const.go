package creem

const (
	BaseURL     = "https://api.creem.io"
	TestBaseURL = "https://test-api.creem.io"
)

type Mode string

const (
	TestMode    Mode = "test"
	ProdMode    Mode = "prod"
	SandboxMode Mode = "sandbox"
)
