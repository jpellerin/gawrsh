VPATH = $(PATH)

default:
	go build

test:
	go test

start: goreman
	goreman start

goreman:
	go get github/mattn/goreman
