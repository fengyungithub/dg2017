#!/usr/bin/make -f

GO=api/src/github.com/nickschuch/articles
PHP=clients/php

generate:
	rm -fR $(GO) $(PHP)
	mkdir -p $(GO) $(PHP)
	protoc -I . articles.proto --go_out=plugins=grpc:$(GO) --php_out=plugins=grpc:$(PHP)
