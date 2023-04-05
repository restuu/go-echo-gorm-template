# Golang Webserver

This is a simple webserver written in Golang using Echo as the web framework and GORM as the ORM.

## Installation

To run this webserver, you need to have Go installed on your system. You can download it from https://golang.org/dl/.

Then, clone this repository using:

```bash
git clone https://github.com/restuu/go-echo-gorm-template.git

# install dependencies
go mod download

# setup env

# run
go run cmd/webserver/ --config=.env
```