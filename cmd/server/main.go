package main

import (
	"context"
	"log"
	"w3labs/internal/adapters"
	"w3labs/internal/config"

	"github.com/caarlos0/env/v11"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {

	ctx := context.Background()

	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to parse configs. %+v", err)
	}

	logger := adapters.NewSlog(cfg.Logger)
	logger.Debug("APPLICATION HAS STARTED", nil)

	awsCfg, err := awsConfig.LoadDefaultConfig(ctx, func(o *awsConfig.LoadOptions) error {
		if cfg.AwsEndpoint != "" {
			o.BaseEndpoint = cfg.AwsEndpoint
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to connecto to aws. %+v", err)
	}

	dynamodbClient := dynamodb.NewFromConfig(awsCfg)
	tables, err := dynamodbClient.ListTables(ctx, &dynamodb.ListTablesInput{})
	if err != nil {
		log.Fatalf("Failed to list tables. %+v", err)
	}

	logger.Info("Tables: ", map[string]any{
		"tables": tables,
	})
}
