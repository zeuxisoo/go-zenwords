# Zenwords

An experimental program for providing text filtering services

## Usage

Install the vendor (if not exists)

    dep ensure

Get the tools

    make tools

Generate the related code from protocol buffers

    make generate

Build the application

    go build -o zenwords *.go
