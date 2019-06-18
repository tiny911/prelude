all:
	go build -o cmd/cmd cmd/main.go
clean:
	rm -rf bin/*