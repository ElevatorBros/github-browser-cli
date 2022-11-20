BINARY_NAME=main.out

build:
	go build -o ${BINARY_NAME} src/*.go

run:
	go build -o ${BINARY_NAME} src/*.go
	./${BINARY_NAME} $(arg)

clean:
	go clean
	rm ${BINARY_NAME}

deps:
	go get github.com/ktr0731/go-fuzzyfinder 
