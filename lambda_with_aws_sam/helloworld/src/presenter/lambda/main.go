package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) error {
	fmt.Println("hello world" + os.Getenv("AnyParameterYouWant"))
	return nil
}

func main() {
	lambda.Start(handler)
}
