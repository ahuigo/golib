package main

import (
	"fmt"
	"testing"

	"gopkg.in/yaml.v3"

	"log"
)

func TestYamlAnchor(test *testing.T) {
	type Config struct {
		Main struct {
			Host string `yaml:"host"`
		} `yaml:"main"`
		Pg struct {
			Host   string `yaml:"host"`
			Dbname string `yaml:"dbname"`
		} `yaml:"pg"`
		Pg2 struct {
			Host   string `yaml:"host"`
			Dbname string `yaml:"dbname"`
		} `yaml:"pg2"`
	}

	data := `
main: &main-keto
  host: "pg-inner.com"
pg: 
  <<: *main-keto
  dbname: "dbname1"
pg2: *main-keto
`
	config := Config{}

	err := yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("Host: %s\n", config.Pg.Host)
	fmt.Printf("Host: %v\n", config.Pg2)
}
