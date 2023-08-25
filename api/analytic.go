package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

type AnalyticRequest struct {
	*TerraformRequest

	// OrganizationID is the id of the cloudavenue organization
	OrganizationID string `json:"organizationId"`
	// ResourceName is the name of the resource
	ResourceName string `json:"resourceName"`
	// Action is the action performed on the resource (create, update, delete, read or import)
	Action string `json:"action"`

	// ExecutionTime is the time in ms to execute the action
	ExecutionTime int64 `json:"executionTime"`

	// Data is the interface containing extra data
	Data map[string]interface{} `json:"data,omitempty"`
}

func (a *AnalyticRequest) Bind(r *http.Request) error {
	// TerraformRequest is a required field
	if a.TerraformRequest == nil {
		return errors.New("missing required terraformRequest fields")
	} else if err := a.TerraformRequest.Bind(r); err != nil {
		return err
	}

	return nil
}

func (a AnalyticRequest) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

// AnalyticResponse is the response payload for the Analytic data model.
type AnalyticResponse struct{}
