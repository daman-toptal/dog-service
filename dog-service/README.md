# Dog Fetcher

## Overview

This is an example project which demonstrates the use of microservices for a command line application. The backend is
powered by 2 microservices, all of which happen to be written in Go and using Docker
to isolate and deploy the ecosystem.

* External Service : Parses user commands, communicates with gRPC and save response as file
* Internal Service : Makes call to Dog API for downloading images given breed name and count

## Deployment

The application can be deployed in **local machine**. You can find the appropriate documentation in the following link:

* [local machine (docker compose)](./docs/localhost.md)

## Available Commands

```bash
$ get photo {breed_name} {count} # use get photo to download photos, count is optional and default is 1
$ exit # exit to stop the run
```
