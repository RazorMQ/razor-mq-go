package client

type RazorConfig struct {
	Host string
	Port string
}

type RazorMQClient struct {
	Config RazorConfig
}
