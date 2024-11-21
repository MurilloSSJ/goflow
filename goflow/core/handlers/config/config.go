package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)
func LoadConfig() Config {
	f, err := os.OpenFile("env.yml", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	var cfg Config
	decoder := yaml.NewDecoder(f)
	errDecode := decoder.Decode(&cfg)
	if errDecode != nil {
		log.Fatal(err)
		panic(errDecode)
	}
	return cfg
}

