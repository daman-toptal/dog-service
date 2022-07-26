# Service Localhost Deployment

## Overview

This project can be deployed in a single machine (localhost) using docker compose in order to know the behavior of
microservices.

## Index

- [Service Localhost Deployment](#service-localhost-deployment)
  - [Overview](#overview)
  - [Index](#index)
  - [Requirements](#requirements)
  - [Starting services](#starting-services)
  - [Testing services](#testing-services)
    - [Requirements](#requirements-1)
  - [Stoping services](#stoping-services)

## Requirements

* Docker Engine 20.10.11
* Docker Compose 1.29.2
  
## Starting services

Use the following command to deploy all services in your local environment.

```bash
$ cd dog-service #go to project directory
$ docker-compose build #build image
$ docker-compose create #create containers
$ docker-compose start dog-service-internal #start internal service in background
$ docker-compose run dog-service-external #run commands to interact with external service
```

The following command is an example of how to get the photo :

```bash
$ get photo hound

n02088094_3944.jpg saved at ./images/hound_n02088094_3944.jpg

$ get photo hound 2

n02088094_4314.jpg saved at ./images/hound_n02088094_4314.jpg
n02091244_4221.jpg saved at ./images/hound_n02091244_4221.jpg

```

## Testing services

### Requirements

* Go 1.17.x

```bash
$ cd dog-service/internal/service && go test # unit testing

PASS
ok      dog-service/internal/service    2.581s

$ cd dog-service/internal/command && go test -v # unit testing #integration testing, requires dog-service-internal to be running
=== RUN   TestCommandInvalidCommand
--- PASS: TestCommandInvalidCommand (0.00s)
=== RUN   TestCommandGetPhoto
n02089973_4185.jpg saved at ./images/hound_n02089973_4185.jpg
--- PASS: TestCommandGetPhoto (1.88s)
=== RUN   TestCommandGetPhotoBlankName
breed name is required
--- PASS: TestCommandGetPhotoBlankName (0.00s)
=== RUN   TestCommandGetPhotoInvalidCount
invalid count
--- PASS: TestCommandGetPhotoInvalidCount (0.00s)
=== RUN   TestCommandGetPhotoInvalidParam
too many params
--- PASS: TestCommandGetPhotoInvalidParam (0.00s)
PASS
ok      dog-service/internal/command    2.279s
```

## Stoping services

```bash
$ ^C or exit #stop the external service run
$ docker-compose stop dog-service-internal #stop the internal service

Stopping dog-service_dog-service-internal_1    ... done
```

