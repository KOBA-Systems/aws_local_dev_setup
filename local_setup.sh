#/bin/bash

S3_BUCKET="local-aws-demo"
LAMBDA_NAME="local-test"

alias laws="aws --endpoint-url=http://localhost:4566"

# Create S3 bucket
laws s3 mb s3://${S3_BUCKET}

# Build lambda deployment package
GOOS=linux go build -o ./main .
zip deployemnt.zip main; rm main

# Deploy lambda
laws lambda create-function \
    --region us-east-1 \
    --function-name ${LAMBDA_NAME} \
    --runtime go1.x \
    --handler main \
    --memory-size 128 \
    --zip-file fileb://deployemnt.zip \
    --role arn:aws:iam::000000000000:role/irrelevant \
    --environment \
        "{\"Variables\":
            {
                \"AWS_REGION\": \"us-east-1\",
                \"AWS_ENDPOINT\": \"http://awsservices:4566\",
                \"S3_BUCKET\": \"${S3_BUCKET}\"
            }
        }"

# Running a lambda
# aws --endpoint-url=http://localhost:4566 lambda invoke --function-name local-test --payload '{"id": "00001", "name": "Koba Systems"}' output.json
# aws --endpoint-url=http://localhost:4566 s3 ls s3://local-aws-demo/