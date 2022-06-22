package service

import (
	"github.com/adi221/good-reads/pkg/config"
	"github.com/adi221/good-reads/pkg/db"
	"github.com/adi221/good-reads/pkg/sanitizer"
	"github.com/adi221/good-reads/pkg/scraper"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var instance *Registry

// Registry is the structure definition of service registry
type Registry struct {
	conf       config.Config
	db         db.DB
	logger     zerolog.Logger
	webScraper scraper.WebScraper
	sanitizer  *sanitizer.Sanitizer
}

// Configure the global service registry
func Configure(conf config.Config, database db.DB) error {
	webScraper, err := scraper.NewWebScraper("")
	if err != nil {
		return err
	}
	blockList, err := sanitizer.NewBlockList(conf.General.BlockList)
	if err != nil {
		return err
	}

	instance = &Registry{
		conf:       conf,
		db:         database,
		webScraper: webScraper,
		logger:     log.With().Str("component", "service").Logger(),
		sanitizer:  sanitizer.NewSanitizer(blockList),
	}
	return nil
}

// Lookup returns the global service registry
func Lookup() *Registry {
	if instance != nil {
		return instance
	}
	log.Fatal().Msg("Service registry is not configured")
	return nil
}
