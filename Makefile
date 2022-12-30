build:
	go build -o bin/moneyconverter

run: build
	./bin/moneyconverter -c config.prod.json