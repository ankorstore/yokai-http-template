# Yokai HTTP Template

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go version](https://img.shields.io/badge/Go-1.20-blue)](https://go.dev/)

> HTTP application template based on the [Yokai](https://github.com/ankorstore/yokai) Go framework.

<!-- TOC -->
* [Overview](#overview)
* [Getting started](#getting-started)
  * [Installation](#installation)
    * [With gonew](#with-gonew)
    * [With GitHub](#with-github)
  * [Usage](#usage)
* [Template contents](#template-contents)
  * [Layout](#layout)
  * [Makefile](#makefile)
<!-- TOC -->

## Overview

This template provides:

- a ready to extend [Yokai](https://github.com/ankorstore/yokai) application, with the [fxhttpserver](https://github.com/ankorstore/yokai/tree/main/fxhttpserver) module installed
- a ready to use [dev environment](docker-compose.yaml), based on [Air](https://github.com/cosmtrek/air) (for live reloading)
- some examples of [handler](internal/handler/welcome.go), [service](internal/service/welcome.go) and [tests](internal/handler/welcome_test.go) to get started

See the [Yokai documentation](https://ankorstore.github.io/yokai) for more details.

## Getting started

### Installation

#### With gonew

You can install [gonew](https://go.dev/blog/gonew), and simply run:

```shell
gonew github.com/ankorstore/yokai-http-template github.com/foo/bar
cd bar
make fresh
```

#### With GitHub

You can also create your repository [using the GitHub template](https://github.com/new?owner=ankorstore&template_name=yokai-http-template).

After cloning your repository, simply run:

```shell
make rename to=foo/bar
make fresh
```

### Usage

After a short moment, the application will be available on:
- [http://localhost:8080](http://localhost:8080) for the application
- [http://localhost:8081](http://localhost:8081) for the dashboard

## Template contents

### Layout

This template is following the [standard go project layout](https://github.com/golang-standards/project-layout):

- `cmd/`: entry points
- `configs/`: configuration files
- `internal/`:
  - `handler/`: HTTP handler and test examples
  - `service/`: service and test examples
  - `bootstrap.go`: bootstrap (modules, lifecycles, etc)
  - `routing.go`: routing
  - `services.go`: dependency injection

### Makefile

This template provides a `Makefile`:

```
make up     # start the docker compose stack
make down   # stop the docker compose stack
make logs   # stream the docker compose stack logs
make fresh  # refresh the docker compose stack
make test   # run tests
make lint   # run linter
```
