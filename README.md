# form3interview
Interview with Form3, golang client for AccountsAPI - by Peter Hooper


## Description

This repository contains the code to complete the [interview question by Form3](https://github.com/form3tech-oss/interview-accountapi).
In summary it is a `go` client for the accounts API [specified here](https://api-docs.form3.tech/api.html#organisation-accounts).
This is the first time I've written any commercial grade code in golang... I'm also used to paired/mob programming so this has been a fun and useful learning exercise - looking forward for feedback.

## Quick Start

After cloning the repository and changing your working directory to that folder,
the complete tests suite of unit tests and integration tests (BDD-style) can be run using either:

```{bash}
docker-compose up
```
or
```{bash}
make start
```

## Usage

### In you own application

This package is designed to be used in another `go` application,
to this end there is an [example.go](https://github.com/diversemix/form3interview/blob/master/example.go) in the root of the folder.
This can be build and run as is designed to be used as a reference implementation.

The code below, shows the minimum code required to create the client. The `accountclient` package also exposes a `New()` method to allow the client to use a custom logger that uses the [ClientLogger](https://github.com/diversemix/form3interview/blob/master/accountclient/client.go) interface and a custom repository that uses [Repository](https://github.com/diversemix/form3interview/blob/master/accountclient/repository.go) interface. In the implementation below it uses the `NewRestClient()` method to create the default REST implementation of the client.

```{go}
package main

import (
	"log"
	"github.com/diversemix/form3interview/accountclient"
)

func main() {
	logger := log.New(log.Writer(), "[client] ", log.Flags())
	client := accountclient.NewRestClient(logger, "http://localhost:8080")
}
```

### During Development

The [Makefile](https://github.com/diversemix/form3interview/blob/master/Makefile) has been designed to help development. Running the command `make` will list all the commands and a brief description as a reminder. They are explained in detail below:

- `make start` - This is equivalent to `docker-compose up`, but with the added step of stopping any currently running services beforehand. Therefore this is recommended over using docker-compose directly.
- `make start-services` - This just starts the supporting services for running the BDD tests. Handy to leave these running when you are developing the BDD tests.
- `make stop` - This will stop any running containers that are managed by the docker-compose file.
- `make bdd-tests` - This runs the integration tests, it depends on having the supporting services running and the environment variable `URL_UNDER_TEST` set to the accounts service you are testing, typically `http://localhost:8080`
- `make unit-tests` - This runs the unit tests, it has no dependents and is designed to run quickly to allow *"failing fast"* as a part of normal development.
- `make test` - This runs all the tests and so has the same requirements as the `bdd-tests` command. This is used as the start up command for testing container in docker-compose.

I've also found installing [looper](http://github.com/nathany/looper) has been beneficial during development.
```
go get -u github.com/nathany/looper
```

## Approach

### Issues

1. Initially it was found that the container `accountapi` could not connect to the database. It was found there was a startup script missing - this was added in `scripts/db/init.sql` ***

2. Later it was found that the container `accountapi` sometimes panicked on startup. To avoid this happening, I extracted the `entrypoint.sh` and modified it, this script is in `scripts/new-entrypoint.sh` ***

3. I have commented out the following scenarios - I would expect them to work but they do not. Maybe its my misunderstanding of how the server works?

```

  # Scenario: Create an Account in GB without a Bic
  #   Given the country is "GB"
  #   And the Bic is ""
  #   Then the Account should be created in memory
  #   And the Account should not be created over the API

  # Scenario: Create an Account in GB without a BankID
  #   Given the country is "GB"
  #   And the BankID is ""
  #   Then the Account should be created in memory
  #   And the Account should not be created over the API
```

### Design Thinking

1. I was keen to ensure that the client that was to be created would be as idiomatic and easy to use for the developer consuming it. I started with the file `example.go` and have attempted to make it as easy to use as possible. Having learnt about BDD along the way, I can see the advantage of starting here - I am now a convert :)

2. For the purposes of the *"Single Responsibility Principle"* the following
    - [RestRepository](https://github.com/diversemix/form3interview/blob/master/accountclient/rest-repository.go) is responsible for all network transport (the only file using `net/http`)
    - [entities/account.go](https://github.com/diversemix/form3interview/blob/master/accountclient/entities/account.go) is responsible for representing the entity that the API deals with. Validation has been used within this type given the documentation of the API. Again this is to *fail fast* so the validation can fail client-side before having to round-trip with all the network latency involved with that.
    - [data-mapper.go](https://github.com/diversemix/form3interview/blob/master/accountclient/data-mapper.go) - is responsible for (un)marshalling the `Account` objects to/from `JSON` - this was a term borrowed from DDD, where effectively JSON is being used for the DTO.
    - [client.go](https://github.com/diversemix/form3interview/blob/master/accountclient/client.go) has the responsibility solely for orchestration between the `Interface` and the `Repository` selected on creation.

3. Interfaces, there are the following interfaces to use the SIP interface-segregation principle.

- ClientInterface - implemented by the main package "accountclient"
- ClientLogger - implemented by the standard "log" package, but abstracted for easy of testing.
- Response - implemented by `RestResponse`
- Repository - implemented by `RestRepository`

### Testing Strategy

Normally I would ensure there are the three level of tests are complete before a product goes out the door. These are:
1. Unit Tests
2. Integration Tests
3. End-to-End Tests

Without a staging environment or equivalent it is impossible to do (3). I have completed what I think is a good "first pass" at the unit tests for this interview - obviously I could carry on! In terms of integration tests, as we have the container available to us, I have attempted to do these in BDD style using godog. This is the first time I have used this technique, again would appreciate any feedback. Thanks!

## Still to do

- Validation of bic and bank_id in the entity?
- Improve unit tests - there are many tests that need more cases, particularly failure cases.
- setup up CI/CD with github actions.
- Improve BDD tests - more reading first... I'm a bit worried at the size of the file `integration_tests.go`
- Improve Account to include all the other possible attributes it can contain (see docs).
