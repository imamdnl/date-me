<p align="center">
    <img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" align="center" width="30%">
</p>
<p align="center"><h1 align="center">DATE-ME</h1></p>
<p align="center">
	<em><code>Dating app</code></em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/last-commit/imamdnl/date-me?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/imamdnl/date-me?style=default&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/imamdnl/date-me?style=default&color=0080ff" alt="repo-language-count">
</p>
<p align="center"><!-- default option, no dependency badges. -->
</p>
<p align="center">
	<!-- default option, no dependency badges. -->
</p>
<br>

##  Table of Contents

- [ Overview](#-overview)
- [ Features](#-features)
- [ Project Structure](#-project-structure)
- [ Getting Started](#-getting-started)
  - [ Prerequisites](#-prerequisites)
  - [ Installation](#-installation)
  - [ Usage](#-usage)

---

##  Overview

<code>❯ Date me is a dating app</code>

---

##  Features

<code>❯ Register User</code>\
<code>❯ Login User</code>

---

##  Project Structure

```sh
└── date-me/
    ├── Date Me.postman_collection.json
    ├── README.md
    ├── cmd
    │   └── main.go
    ├── domain
    │   ├── constants
    │   │   └── data.go
    │   ├── dto
    │   │   └── user.go
    │   ├── entity
    │   │   └── user.go
    │   ├── repository
    │   │   ├── user_repo.go
    │   │   └── wrapper.go
    │   ├── usecase
    │   │   ├── user.go
    │   │   └── wrapper.go
    │   └── value_object
    │       └── user_gender.go
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── delivery
    │   │   ├── http
    │   │   ├── middleware
    │   │   └── routes.go
    │   ├── repository
    │   │   ├── get_user_by_email.go
    │   │   ├── init.go
    │   │   ├── mapper
    │   │   ├── model
    │   │   └── store_user.go
    │   └── usecase
    │       ├── init.go
    │       ├── login_user.go
    │       └── register_user.go
    └── pkg
        ├── common
        │   ├── base.go
        │   └── jwt.go
        ├── config
        │   ├── database.go
        │   ├── environment.go
        │   └── postgres.go
        ├── exceptions
        │   └── customerror.go
        └── utils
            └── util_http.go
```




---
##  Getting Started

###  Prerequisites

Before getting started with date-me, ensure your runtime environment meets the following requirements:

- **Programming Language:** Go 1.23
- Postman
- Docker with Docker compose


###  Installation

Install date-me using one of the following methods:

**Build from source:**

1. Clone the date-me repository:
```sh
git clone https://github.com/imamdnl/date-me
```

2. Navigate to the project directory:
```sh
cd date-me
```

3. Install the project dependencies:


**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
go mod tidy
go mod vendor
```

4. Build Docker image:
```sh
docker build -t date-me:latest -f .deploy/Dockerfile .
```


###  Usage
Run date-me using the following command:
**Using `docker compose`**

```sh
docker compose -f .deploy/docker-compose.yml up -d
```

to stop the service, run command:
```sh
docker compose -f .deploy/docker-compose.yml down -v
```


###  Testing
Import the `Date Me.postman_collection.json` Postman collection and run the collection


### Env
| ENV Name               | Required | Default Value | Purpose                   |
|------------------------|----------|---------------|---------------------------|
| `PORT`                 | Y        |               | Port REST API application |
| `DB_POSTGRES_HOST`     | Y        |               | DB PostgreSQL Config      |
| `DB_POSTGRES_PORT`     | Y        |               | DB PostgreSQL Config      |
| `DB_POSTGRES_USERNAME` | Y        |               | DB PostgreSQL Config      |
| `DB_POSTGRES_PASSWORD` | Y        |               | DB PostgreSQL Config      |
| `DB_POSTGRES_NAME`     | Y        |               | DB PostgreSQL Config      |
| `DB_POSTGRES_SCHEMA`   | Y        |               | DB PostgreSQL Config      |
| `JWT_SECRET`           | Y        |               | JWT Config                |
| `JWT_EXPIRY`           | Y        |               | JWT Config                |

