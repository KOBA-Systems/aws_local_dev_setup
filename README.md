# aws_local_dev_setup
Example to show how to setup AWS local development environment using golang aws-sdk-v2 and localstack.
In this example, there are 2 AWS components, S3 and Lambda. When the Lambda function is invoked, it will create a result file in S3 bucket. Both components are run locally without communicating with the external AWS services.  

## 1.Prerequisites
Need to install:
  - Docker and Docker Compose
  - AWS CLI v2
  - Golang v1.15


## 2.How to run
- Start docker compose: `docker-compose up`
- After docker compose starts up,run the setup script to create S3 bucket, build and deploy lambda function: `./local_setup.sh`
- Run the following command to invoke the lambda function: `aws --endpoint-url=http://localhost:4566 lambda invoke --function-name local-test --payload '{"id": "00001", "name": "Koba Systems"}' output.json`
- Check if the file is created in S3: `aws --endpoint-url=http://localhost:4566 s3 ls s3://local-aws-demo/`
