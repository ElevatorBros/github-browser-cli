BINARY_NAME=main.out
FLAGS=-ldflags "-w"

build:
	go build ${FLAGS} -o ${BINARY_NAME} src/*.go 

run:
	make build
	./${BINARY_NAME} $(arg)

clean:
	go clean
	rm ${BINARY_NAME}

deps:
	go get github.com/ktr0731/go-fuzzyfinder 
	go get github.com/PuerkitoBio/goquery
