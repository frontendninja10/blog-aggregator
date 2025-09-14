package app

import (
	"github.com/frontendninja10/blog-aggregator/internal/config"
	"github.com/frontendninja10/blog-aggregator/internal/database"
)

type State struct {
	Config *config.Config
	DB *database.Queries
}

func NewState(cfg *config.Config, db *database.Queries) *State {
	return &State{
		Config: cfg,
		DB: db,
	}
}