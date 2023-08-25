package main

import (
	"encoding/json"
	"time"

	"github.com/FrangipaneTeam/terraform-analytic-tool/api"
	"github.com/FrangipaneTeam/terraform-analytic-tool/clients"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func consumer(r *clients.RedisClient, i *clients.InfluxDBClient) {
	subscriber := r.Subscribe(ctx, "PubSubCavTerraformAnalytics")
	trackingWriteAPI := i.NewWriteAPI()

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		x := &api.AnalyticRequest{}
		if err := json.Unmarshal([]byte(msg.Payload), x); err != nil {
			panic(err)
		}

		tags := map[string]string{
			"organizationID":       x.OrganizationID,
			"terraformExecutionID": x.TerraformExecutionID,
			"clientVersion":        x.ClientVersion,
			"resourceName":         x.ResourceName,
			"action":               x.Action,
		}
		fields := map[string]interface{}{
			"value":         1,
			"executionTime": x.ExecutionTime,
		}
		point := write.NewPoint("tracking", tags, fields, time.Now())

		trackingWriteAPI.WritePoint(point)
	}
}
