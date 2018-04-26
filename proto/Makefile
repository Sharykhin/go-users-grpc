#!make
.PHONY: default generate

default: generate

generate:
	protoc -I . --go_out=plugins:. user.proto