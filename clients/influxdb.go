package clients

import (
	influxdb "github.com/influxdata/influxdb-client-go/v2"
	influxdbapi "github.com/influxdata/influxdb-client-go/v2/api"
)

type InfluxDBConfig struct {
	Address string `yaml:"address"` // host:port address.
	Token   string `yaml:"token"`   // Optional password.
	Org     string `yaml:"org"`
	Bucket  string `yaml:"bucket"`
}

type InfluxDBClient struct {
	influxdb.Client
	config InfluxDBConfig
}

func NewInfluxDBClient(i InfluxDBConfig) *InfluxDBClient {
	// Check if all fields are set
	if i.Address == "" {
		panic("InfluxDB address is not set")
	}
	if i.Token == "" {
		panic("InfluxDB token is not set")
	}
	if i.Org == "" {
		panic("InfluxDB org is not set")
	}
	if i.Bucket == "" {
		panic("InfluxDB bucket is not set")
	}

	client := influxdb.NewClient(i.Address, i.Token)

	return &InfluxDBClient{
		client,
		i,
	}
}

// GetOrg returns the organization name.
func (i *InfluxDBClient) GetOrg() string {
	return i.config.Org
}

// GetBucket returns the bucket name.
func (i *InfluxDBClient) GetBucket() string {
	return i.config.Bucket
}

// NewWriteAPI returns the write api.
func (i *InfluxDBClient) NewWriteAPI() influxdbapi.WriteAPI {
	return i.WriteAPI(i.config.Org, i.config.Bucket)
}
