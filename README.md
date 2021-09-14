# Go WhatsApp REST API Implementation Using Fiber And Swagger
Package cooljar/go-whatsapp-fiber Implements the WhatsApp Web API using [Fiber](https://github.com/gofiber/fiber) web framework,
and also [Swag](https://github.com/gofiber/fiber) to generate Swagger Documentation (2.0).
<br>This repository contains example of implementation [Rhymen/go-whatsapp](https://github.com/Rhymen/go-whatsapp) package.

## Getting Started
These instructions will get you a copy of the project up and running on docker container and on your local machine.

### Prerequisites
Prequisites package:
* [Docker](https://www.docker.com/get-started) - for developing, shipping, and running applications (Application Containerization).
* [Go](https://golang.org/) - Go Programming Language
* [Swag](https://github.com/swaggo/swag) - converts Go annotations to [Swagger Documentation 2.0](https://swagger.io/docs/specification/2-0/basic-structure/)
* [Make](https://www.gnu.org/software/make/manual/make.html) - Automated Execution using Makefile

Optional package:
* [gosec](https://github.com/securego/gosec) Golang Security Checker. Inspects source code for security problems by scanning the Go AST

## Config 
1. Rename `.env.example` to `.env` 
2. fill it environment with your make setting.

### Running On Docker Container
- Run project by using following command:
```bash
$ make run

# Process:
#   - Generate API docs by Swagger
#   - Build and run Docker containers
```
- Stop application by using following command:
```bash
$ make stop

# Process:
#   - Stop and remove app container
#   - remove image
```

- Run Container without Rebuid by using following command:
```bash
$ make docker_run_image

# Process:
#   - Run Docker containers, if image not found will build image
```

### Running On Local Machine
Below is the instructions to run this project on your local machine:
1. Open new `terminal`.
2. Set `run.sh` file permission.
```bash
$ chmod +x ./run.sh
```
3. Run application from terminal by using following command:
```bash
$ ./run.sh
```

### API Access
Go to your API Docs page: [127.0.0.1:3000/swagger/index.html](http://127.0.0.1:3000/swagger/index.html) and fill Basic Auth with `BASIC_AUTH_USERNAME` and `BASIC_AUTH_PASSWORD`, let `BASIC_AUTH_USERNAME` and `BASIC_AUTH_PASSWORD` empty if you want without basic auth
<br>
API Docs page will be look like:
<br><img src="https://raw.githubusercontent.com/cooljar/go-whatsapp-fiber/main/sc.png" width="500">

Below is the instructions to perform messaging:
* Make sure your computer is connected to the internet.
* Prepare your smartphone and make sure the internet is active.
* Hit the [Login](http://127.0.0.1:3000/swagger/index.html#/Whatsapp/post_v1_whatsapp_login) endpoint, you will see a QR Code if request was success.
  <br><img src="https://raw.githubusercontent.com/cooljar/go-whatsapp-fiber/main/qr.png" width="250">
  <br>Check your `.env` environment setting if an error occurred.
* Scant it, and done.
Now you can perforn all endpoint to send a message.

## Testing
- Inspects source code for security problems using [gosec](https://github.com/securego/gosec). You need to install it first.
- Execute unit test by using following command:
```bash
$ make test
```

## Built With
* [Go](https://golang.org/) - Go Programming Languange
* [Go Modules](https://github.com/golang/go/wiki/Modules) - Go Dependency Management System
* [Make](https://www.gnu.org/software/make/) - GNU Make Automated Execution
* [Docker](https://www.docker.com/) - Application Containerization

## Authors
* **Fajar Rizky** - *Initial Work* - [cooljar](https://github.com/cooljar)
* **Ahmad Yusri** - *Reconfig* - [ahmadyusri](https://github.com/ahmadyusri)

## More
-------
