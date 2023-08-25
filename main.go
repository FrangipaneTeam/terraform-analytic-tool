package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/FrangipaneTeam/terraform-analytic-tool/api"
	"github.com/FrangipaneTeam/terraform-analytic-tool/clients"
)

var ctx = context.Background()

func main() {
	var (
		configFile   string
		onlyAPI      bool
		onlyConsumer bool
	)

	// Read flags passed to the program
	flag.StringVar(&configFile, "config", "", "Path to the configuration file")
	flag.BoolVar(&onlyAPI, "only-api", false, "Run only the API")
	flag.BoolVar(&onlyConsumer, "only-consumer", false, "Run only the consumer")
	flag.Parse()

	if onlyAPI && onlyConsumer {
		panic("only-api and only-consumer flags cannot be used together. Unset all flags to run API and consumer")
	}

	// Load configuration
	cfg, err := readConfig(configFile)
	if err != nil {
		panic(err)
	}

	// Setup Redis Client
	redisClient := clients.NewRedisClient(clients.RedisConfig{
		Address:      cfg.Redis.Address,
		Password:     cfg.Redis.Password,
		MaxRetries:   cfg.Redis.MaxRetries,
		DialTimeout:  cfg.Redis.DialTimeout,
		ReadTimeout:  cfg.Redis.ReadTimeout,
		WriteTimeout: cfg.Redis.WriteTimeout,
	})

	influxdbClient := clients.NewInfluxDBClient(clients.InfluxDBConfig{
		Address: cfg.InfluxDB.Address,
		Token:   cfg.InfluxDB.Token,
		Org:     cfg.InfluxDB.Org,
		Bucket:  cfg.InfluxDB.Bucket,
	})

	if !onlyAPI && !onlyConsumer {
		onlyAPI = true
		onlyConsumer = true
	}

	if onlyConsumer {
		go consumer(redisClient, influxdbClient)
	}

	if onlyAPI {
		api.New(api.Settings{
			Address: cfg.API.Address,
			Port:    cfg.API.Port,
			Token:   cfg.API.Token,

			RedisClient:    redisClient,
			InfluxDBClient: influxdbClient,
		})
	}

	// Wait forever to keep the program running and catch CTRL+C signal
	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	// TODO close redis and influxdb connections

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.

	log.Default().Println("shutting down")
	os.Exit(0)
}
