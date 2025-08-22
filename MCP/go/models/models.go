package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ImageUrl represents the ImageUrl schema from the OpenAPI specification
type ImageUrl struct {
	Url string `json:"Url,omitempty"`
}

// ImagePredictionResultModel represents the ImagePredictionResultModel schema from the OpenAPI specification
type ImagePredictionResultModel struct {
	Iteration string `json:"Iteration,omitempty"`
	Predictions []ImageTagPredictionModel `json:"Predictions,omitempty"`
	Project string `json:"Project,omitempty"`
	Created string `json:"Created,omitempty"`
	Id string `json:"Id,omitempty"`
}

// ImageTagPredictionModel represents the ImageTagPredictionModel schema from the OpenAPI specification
type ImageTagPredictionModel struct {
	Probability float32 `json:"Probability,omitempty"`
	Tag string `json:"Tag,omitempty"`
	Tagid string `json:"TagId,omitempty"`
}
