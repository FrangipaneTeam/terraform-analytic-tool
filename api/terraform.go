package api

import (
	"errors"
	"net/http"
)

// TerraformRequest is the request payload for Terraform data model.
type TerraformRequest struct {
	// TerraformExecutionID is the uniq id generated at the beginning of the terraform execution
	TerraformExecutionID string `json:"terraformExecutionId"`
	// ClientToken is the key used to identify the client
	ClientToken string `json:"clientToken"`
	// ClientVersion is the version of the provider
	ClientVersion string `json:"version"`
}

func (a *TerraformRequest) Bind(_ *http.Request) error {
	// TerraformExecutionID is a required field
	if a.TerraformExecutionID == "" {
		return errors.New("missing required terraformExecutionId fields")
	}

	// ClientToken is a required field
	if a.ClientToken == "" {
		return errors.New("missing required clientToken fields")
	}

	// ClientVersion is a required field
	if a.ClientVersion == "" {
		return errors.New("missing required version fields")
	}

	return nil
}

// TerraformResponse is the response payload for the Terraform data model.
type TerraformResponse struct{}
