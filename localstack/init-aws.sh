#!/bin/bash
set -xe

echo "Creating initial resources..."

awslocal dynamodb create-table \
    --table-name w3labs \
    --billing-mode PAY_PER_REQUEST \
    --attribute-definitions \
        AttributeName=PK,AttributeType=S \
        AttributeName=SK,AttributeType=S \
        AttributeName=GSI_PK_1,AttributeType=S \
        AttributeName=GSI_SK_1,AttributeType=S \
        AttributeName=GSI_PK_2,AttributeType=S \
        AttributeName=GSI_SK_2,AttributeType=S \
    --key-schema AttributeName=PK,KeyType=HASH AttributeName=SK,KeyType=RANGE \
    --global-secondary-indexes \
        "[
            {\"IndexName\": \"GSI1\", \"KeySchema\": [{\"AttributeName\":\"GSI_PK_1\",\"KeyType\":\"HASH\"},{\"AttributeName\":\"GSI_SK_1\",\"KeyType\":\"RANGE\"}], \"Projection\": {\"ProjectionType\":\"ALL\"}, \"ProvisionedThroughput\": {\"ReadCapacityUnits\":5, \"WriteCapacityUnits\":5}} 
            ,{\"IndexName\": \"GSI2\", \"KeySchema\": [{\"AttributeName\":\"GSI_PK_2\",\"KeyType\":\"HASH\"},{\"AttributeName\":\"GSI_SK_2\",\"KeyType\":\"RANGE\"}], \"Projection\": {\"ProjectionType\":\"ALL\"}, \"ProvisionedThroughput\": {\"ReadCapacityUnits\":5, \"WriteCapacityUnits\":5}} 
        ]"



