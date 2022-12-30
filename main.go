package main

import (
	"encoding/json"
	"flag"
	"os"
)

func main() {
	configFile := flag.String("c", "config.dev.json", "config file")
	flag.Parse()

	var conf config

	byteData, err := os.ReadFile(*configFile)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(byteData, &conf)

	svc := NewLoggingService(NewCurrencyConverter(conf.Url, conf.Key))
	server := NewJSONAPIServer(conf.ListenAddress, svc)
	server.Run()
}

type config struct {
	Key           string `json:"key"`
	ListenAddress string `json:"address"`
	Url           string `json:"url"`
}
