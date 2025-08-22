package main

import (
	"github.com/predictionendpoint/mcp-server/config"
	"github.com/predictionendpoint/mcp-server/models"
	tools_imagepredictionapi "github.com/predictionendpoint/mcp-server/tools/imagepredictionapi"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_imagepredictionapi.CreatePredictimageurlTool(cfg),
		tools_imagepredictionapi.CreatePredictimageurlwithnostoreTool(cfg),
	}
}
