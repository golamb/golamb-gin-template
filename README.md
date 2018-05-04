# Go Lambda Gin Template for GoLambda

## Init project

```sh
golamb init gin my-project
```

## Require

* [dep](https://github.com/golang/dep)
* [aws-lambda-go](https://github.com/aws/aws-lambda-go/)
* [gin-gonic](https://github.com/gin-gonic/gin)
* [aws-sam-local](https://github.com/awslabs/aws-sam-local)

## First Step

* Create function in AWS Lambda
* Choose Runtime to Go 1.x
* Create role or choose existing role

## Atfer create

* edit Handler to "main"
* save and deploy

## Dev

```sh
dep ensure
./build.sh && sam local start-api
```


