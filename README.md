# Yokai HTTP Template

> Yokai HTTP application template.

<!-- TOC -->
* [Overview](#overview)
* [Documentation](#documentation)
  * [Template usage](#template-usage)
  * [Template contents](#template-contents)
    * [Layout](#layout)
    * [Modules](#modules)
    * [Endpoints](#endpoints)
    * [Makefile](#makefile)
<!-- TOC -->

## Overview

This template provides:

- ready to extend [Fx](https://uber-go.github.io/fx/) application, with the [fxhttpserver](https://github.com/ankorstore/yokai/tree/main/fxhttpserver) module pre-configured to offer:
  - http server
  - automatic and correlated logs and traces
  - automatic healthcheck, metrics and debug endpoints
- ready to use [dev environment](docker-compose.yaml) based on [Air](https://github.com/cosmtrek/air) (live reloading)
- examples of [handler](internal/handler/welcome.go), [service](internal/service/info.go) and [functional tests](internal/handler/welcome_test.go) to get started

## Documentation

Check the [Go Fx HTTP application tutorial](https://go-fx-doc.ankorstore.io/tutorials/http-app/) to learn how to use this template to build HTTP applications.

### Template usage

First, create your repository by [using this template](https://github.com/new?owner=ankorstore&template_name=go-fx-http-app-template&template_owner=ankorstore).

Then, in your repository (for example `ankorstore/your-repo`) root directory:

- create the `.env` file from `.env.example`:
  ```shell
  cp .env.example .env
  ```
- adapt the `.env` file with your [GitHub access token](https://github.com/settings/tokens) in the `GH_ACCESS_TOKEN` variable
- rename the go.mod and imports with:
  ```shell
  make rename to=ankorstore/your-repo
  ```
- start the application:
  ```shell
  make up
  ```

You can now access your application on [http://localhost:8080](http://localhost:8080).

### Template contents

#### Layout

This template is following the [standard go project layout](https://github.com/golang-standards/project-layout):

- `cmd/`: entry points
- `configs/`: configuration files
- `internal/`:
  - `handler/`: HTTP handler and test examples
  - `service/`: service and test examples
  - `bootstrap.go`: bootstrap (modules, lifecycles, etc)
  - `routing.go`: routing
  - `services.go`: dependency injection

#### Modules

This template comes with the following modules preloaded and configured:

- [fxbootstrap](https://github.com/ankorstore/yokai/tree/main/fxbootstrap) module, preloading:
  - [fxconfig](https://github.com/ankorstore/yokai/tree/main/fxconfig) module
  - [fxlog](https://github.com/ankorstore/yokai/tree/main/fxlog) module
  - [fxtrace](https://github.com/ankorstore/yokai/tree/main/fxtrace) module
  - [fxhealthcheck](https://github.com/ankorstore/yokai/tree/main/fxhealthcheck) module
- [fxgenerate](https://github.com/ankorstore/yokai/tree/main/fxgenerate) module
- [fxhttpserver](https://github.com/ankorstore/yokai/tree/main/fxhttpserver) module

#### Endpoints

The example [WelcomeHandler](internal/handler/welcome.go) and the pre-configured [fxhttpserver](https://github.com/ankorstore/yokai/tree/main/fxhttpserver) module offer out of the box the following endpoints:

| Endpoint                                                    | Description                                                               |
|-------------------------------------------------------------|---------------------------------------------------------------------------|
| [GET /](http://localhost:8080)                              | [WelcomeHandler](internal/handler/welcome.go) example endpoint            |
| [GET /_health](http://localhost:8080/_health)               | Health check endpoint (enabled if `modules.http.server.healthcheck=true`) |
| [GET /_metrics](http://localhost:8080/_metrics)             | Metrics endpoint (enabled if `modules.http.server.metrics=true`)          |
| [GET /_debug/config](http://localhost:8080/_debug/config)   | Debug config endpoint (enabled if `app.debug=true`)                       |
| [GET /_debug/routes](http://localhost:8080/_debug/routes)   | Debug routes endpoint (enabled if `app.debug=true`)                       |
| [GET /_debug/version](http://localhost:8080/_debug/version) | Debug version endpoint (enabled if `app.debug=true`)                      |
| [GET /_pprof/index](http://localhost:8080/_pprof/index)     | Pprof index dashboard endpoint (enabled if `app.pprof=true`)              |

#### Makefile

This template provides a `Makefile`:

```
make up     # start the docker compose stack
make down   # stop the docker compose stack
make logs   # stream the docker compose stack logs
make fresh  # refresh the docker compose stack
make test   # run tests
make lint   # run linter
```
