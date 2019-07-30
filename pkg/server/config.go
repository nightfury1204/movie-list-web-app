package server

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

const (
	OMDBAPITokenEnv = "OMDB_API_TOKEN"
	DefaultPort     = "8443"
	OMDBAPIUrl      = "http://www.omdbapi.com/"
)

type Config struct {
	// Specifies the server listening port
	Port string

	// Specifies the url of the OMDB api
	OMDBAPIUrl string

	// Specifies the api toke for OMDB
	OMDBAPIToken string

	// This channel is used for graceful server shutdown
	StopCh chan os.Signal
}

func NewConfig() *Config {
	return &Config{
		Port:         DefaultPort,
		OMDBAPIToken: os.Getenv(OMDBAPITokenEnv),
		OMDBAPIUrl:   OMDBAPIUrl,
		StopCh:       make(chan os.Signal, 1),
	}
}

func (c *Config) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.Port, "server-port", c.Port, "Specifies the server listening port.")
	fs.StringVar(&c.OMDBAPIUrl, "omdb-api-url", c.OMDBAPIUrl, "Specifies the url of the OMDB api.")
	fs.StringVar(&c.OMDBAPIToken, "omdb-api-token", c.OMDBAPIToken, "Specifies the api toke for OMDB.Or you can set it by OMDB_API_TOKEN environment variable.")
}

func (o *Config) Validate() error {
	if o.Port == "" {
		return errors.New("server port must be non empty")
	}
	if o.OMDBAPIUrl == "" {
		return errors.New("OMDB api url must be non empty")
	}
	if o.OMDBAPIToken == "" {
		return errors.New("OMDB api token must be non empty")
	}
	return nil
}
