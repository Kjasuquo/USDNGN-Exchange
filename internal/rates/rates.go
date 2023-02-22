package rates

import (
	"github.com/kjasuquo/usdngn-exchange/config"
	"net/http"
)

type Client struct {
	Http   http.Client
	Config config.Config
}

func NewRatesClient(conf config.Config) *Client {
	return &Client{Config: conf}
}
