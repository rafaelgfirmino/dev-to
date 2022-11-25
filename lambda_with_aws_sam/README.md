The tutorial in [dev.to](https://dev.to/rafaelgfirmino/aws-sam-and-go-2fn0)

In the last year, I  worked creating lambda functions for many parts of core businesses.
When I started this journey with lambda I looked about the serverless, but the documentation is partitioned in many plugins repositories, because this point I decided to use AWS SAM.  We do not have the intention the changing our cloud.
AWS Serverless Application Model is a good tool for creating your lambda and some resources like SQS, SNS, S3, and etc.

**Where GO enter?**
Go is compiled and your binary contains only the library needed for work very well, this approach creates a small binary.
Go has very fast performance and your memory usage is very low.
Another good point is the pipeline. The pipeline in GO is very smaller and easy to understand and mantein.
Let me show you how easy is create a serverless function using AWS SAM

[Link of the repository example
](https://github.com/rafaelgfirmino/aws-lambda-series)
# The Golang Function

```go
package main

import (
  "context"
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
```

```yaml
# This is the SAM template that represents the architecture of your serverless application
# https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-template-basics.html

# The AWSTemplateFormatVersion identifies the capabilities of the template
# https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/format-version-structure.html
AWSTemplateFormatVersion: 2010-09-09
Description: >-
    Any description you want
# Transform section specifies one or more macros that AWS CloudFormation uses to process your template
# https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-section-structure.html
Transform:
  - AWS::Serverless-2016-10-31

# Resources declares the AWS resources that you want to include in the stack
# https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resources-section-structure.html

# More info about Globals: https://github.com/awslabs/serverless-application-model/blob/master/docs/globals.rst
Globals:
  Function:
    Timeout: 200
    Runtime: go1.x
    Environment:
      Variables:
        AnyParameterYouWant: !Ref AnyParameterYouWant #this parameter is injected in all lambdas in this file

Parameters:
  stage:
    Type: String
    Default: homologation
  AnyParameterYouWant:
    Type: String
    Default:  this-stage-is-${stage}

Resources:

  HelloWorldFunction:
    Type: AWS::Serverless::Function
    Properties:
      FunctionName: !Sub helloWorld-${stage}
      PackageType: Zip
      CodeUri: src/main
      Handler: lambda
      Tracing: Active # https://docs.aws.amazon.com/lambda/latest/dg/lambda-x-ray.html
      MemorySize: 1024
      Events:
        Schedule:
          Type: Schedule
          Properties:
            Schedule: 'rate(5 minutes)'
            Name: !Sub HelloWorld-${stage} #this is a name of schedule this approach change by stage name
            Description: HelloWorldFunction
            Enabled: true

  HelloWorldLogGroup:
    Type: AWS::Logs::LogGroup
    Properties:
      LogGroupName: !Sub /aws/lambda/${HelloWorldFunction}
      RetentionInDays: 30
    DependsOn: HelloWorldFunction
```
```yaml
name: hello - World
on:
  push:
    branches: [ main ]

jobs:
  build:
    name: hello Wold Build
    runs-on: ubuntu-latest
    defaults:
      run:
        working-directory: go/conciliation
    env:
      ENV_APP: 'HelloWorldTest'
    steps:
      - uses: actions/checkout@v2
      - uses: aws-actions/setup-sam@v1

      - name: aws login
        uses: aws-actions/configure-aws-credentials@v1
        with:
          aws-access-key-id: ${{ secrets.AWS_ACCESS_ID_SAVED_IN_GITHUB_SECRETS }}
          aws-secret-access-key: ${{ secrets.AWS_SECRETS_SAVED_IN_GITHUB_SECRETS }}
          aws-region: 'sa-east-1'

      - name: setup go
        uses: actions/setup-go@v2
        with:
          go-version: '1.19'

      - name: download dependencies
        run: go mod tidy -go=1.19

      - name: Build app with SAM
        run: sam build

      - name: Upaload sam files for S3
        run: |
          sam package --template-file .aws-sam/build/template.yaml \
          --s3-bucket lambda-resources-wieidjdh \
          --output-template-file packaged.yaml
      - name: Deploy sam
        run: |
          sam deploy --template-file packaged.yaml \
          --parameter-overrides \
            Stage=${{ env.ENV_APP }} \
            AnyParameterYouWant="this is value created in pipeline" \
          --stack-name hello-world-${{ env.ENV_APP }} \
          --capabilities CAPABILITY_IAM \
          --no-confirm-changeset
```
**Before running the pipeline you need create the bucket ex: lambda-resources-wieidjdh**

Runing Pipeline

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/d14e8can5ypr9k6kbsz6.png)

Cloudformation stack running after success github actions pipeline finished 

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/9p891icx63a7xa43pctw.png)

The EventBrige has been attached with success in lambda.

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/sjngaakafnknj9yjz4q5.png)

CloudWatch Logs Group has been created with retention for 30 days.

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/7s992v9m6chchmrx711z.png)

Lambda result after running.

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/gspwldyhrwq2nq7kf96s.png)


Aws SAM and Go is a good choice for create a lambda services very fast in aws. 