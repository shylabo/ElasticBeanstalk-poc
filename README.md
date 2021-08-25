# ElasticBeanstalk-poc

This repository is for the test of deploying an application to AWS ElasticBeanstalk without CodeCommit.

Notice that you have to prepare your aws account and set up IAM Role and Security groups before executing eb command below.

In main.go, there should be at least 1 endpoint which is listening and serving.

## install aws-elasticbeanstalk-cli tool

```
$ brew install awsebcli
```

## initialize eb

```
$ eb init
```

- Do you wish to continue with CodeCommit? -> No.
- Do you want to set up SSH for your instance -> No.

```
$ eb create
```

- Select a load balancer type -> classic
- Would you like to enable Spot Fleet request for this environment? -> No.

## go build

```
$ GOOS=linux GOARCH=amd64 go build -o bin/application
```

If this command fails, you might install go modules.

```
$ go mod init
```

## Let's deploy

```
$ eb deploy
```

```
$ eb open
```

Check below.

http://elasticbeanstalk-poc-dev.ap-northeast-1.elasticbeanstalk.com/
