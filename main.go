package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/toshi1127/sql-swagger-generator/swagger"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Service   swagger.ServiceInfo
	Resources map[string]swagger.Resource
}

func main() {
	configFilePath := flag.String("conf", "", "config file path")
	sqlFilePath := flag.String("sql", "", "sql file path")
	outputDirPath := flag.String("outputDir", "", "output directory path")

	flag.Parse()

	if *configFilePath == "" {
		log.Fatal("No config file path given. Use the -conf flag.")
	}
	if *sqlFilePath == "" {
		log.Fatal("No sql file path given. Use the -sql flag.")
	}
	if *outputDirPath == "" {
		log.Fatal("No output directory path given. Use the -outputDir flag.")
	}

	configBytes, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	sqlBytes, err := ioutil.ReadFile(*sqlFilePath)
	if err != nil {
		log.Fatal(err)
	}

	conf := Config{}
	err = yaml.Unmarshal(configBytes, &conf)
	if err != nil {
		log.Fatal(err)
	}

	file_content := string(sqlBytes)
	lines := strings.Split(file_content, "\n")
	generator := swagger.Generator{
		ServiceInfo: conf.Service,
		Resources:   conf.Resources,
		Queries:     lines,
	}

	swag, err := generator.Generate(filepath.Dir(*outputDirPath))
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(*outputDirPath)
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(swag)
}
