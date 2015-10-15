package main

import(
  "github.com/goamz/goamz/aws"
  "github.com/goamz/goamz/sqs"
  "os"
  "os/exec"
  "fmt"
  "time"
)

var(
    queueName = "registry-container-restart-messages"
    topicArn  = "arn:aws:sns:ap-southeast-1:356003847803:registry-cpu-utilization"
    containerName = "validator-registry-api"
    auth = aws.Auth{
        AccessKey: os.Getenv("AWS_ACCESS_KEY_ID"),
        SecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
    }
    conn = sqs.New(auth, aws.APSoutheast)
)

func restartDockerContainer(containerName string) {
  args := []string{"docker", "restart", containerName}

  if err := exec.Command("sudo", args...).Run(); err != nil {
    fmt.Println(os.Stderr, err)
    os.Exit(1)
  }
  fmt.Println("successfully called docker")
}

func messageReceived(queue sqs.Queue, message *sqs.Message) {
    fmt.Println("Message received")
    fmt.Println(message)
    fmt.Println("Deleting message")
    _, err := queue.DeleteMessage(message)
    if err != nil {
      os.Exit(1)
    }
    restartDockerContainer("validatorsripplecom_web")
    restartDockerContainer("validatorsripplecom_nginx")
}

func main() {

    queue, err := conn.GetQueue(queueName)

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    for {
        resp, err := queue.ReceiveMessage(1)

        if err != nil {
          fmt.Println(err)
          os.Exit(1)
        }

        if len(resp.Messages) > 0 {
            messageReceived(*queue, &resp.Messages[0])
        }

        time.Sleep(1 * time.Second)
    }
}



