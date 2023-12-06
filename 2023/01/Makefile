BINARY_NAME=dayone.out
test: build
	./${BINARY_NAME} test.txt

run: build
	./${BINARY_NAME} input.txt

build:
	go build -o ${BINARY_NAME} main.go

clean:
	go clean
	rm ${BINARY_NAME}
