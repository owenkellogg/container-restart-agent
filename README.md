# Container Restart Agent

small executable that polls SQS for messages and issues
docker restart commands when messages are received

## Configuration

The following environment variables must be set:

- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY

