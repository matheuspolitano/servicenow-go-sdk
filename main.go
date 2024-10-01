package main

import (
	"context"
	"os"
	"time"

	"github.com/matheuspolitano/servicenow-go-sdk/config"
	"github.com/matheuspolitano/servicenow-go-sdk/snow"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	ctx := context.Background()

	vConfig, err := config.LoadConfig(".")
	if err != nil{
		log.Fatal().Err(err)
	}
	query := snow.Query{
		snow.FieldQuery{
			Name: "name",
			Value: "auto_close",
			Operator: snow.Equal,
		},
	}
	access := snow.NewAccessConfig(vConfig.SnowUsername, vConfig.SnowPassword, vConfig.Endpoint)
	snow, err := snow.NewSnowClient(*access)
	if err != nil{
		log.Fatal().Err(err)
	}
	body, err := snow.ExecuteQuery(ctx, "em_event_type", query)
	if err != nil{
		log.Fatal().Err(err)
	}
	log.Info().Fields(body).Msg("Snow response")

}
