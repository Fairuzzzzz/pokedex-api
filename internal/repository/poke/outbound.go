package poke

import (
	"github.com/Fairuzzzzz/pokedex-api/internal/configs"
	"github.com/Fairuzzzzz/pokedex-api/pkg/httpclient"
)

type outbound struct {
	cfg    *configs.Config
	client httpclient.HTTPClient
}

func NewPokeOutbound(cfg *configs.Config, client httpclient.HTTPClient) *outbound {
	return &outbound{
		cfg:    cfg,
		client: client,
	}
}
