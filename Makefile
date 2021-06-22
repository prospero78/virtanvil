run:
	clear
	go build -race -o ./bin/gvar ./cmd/gva/main.go
	cd ./bin && \
	./gvar
build:
	clear
	go build -o ./bin/gva ,/cmd/gva/main.go
mod:
	go mod tidy
	go mod vendor
