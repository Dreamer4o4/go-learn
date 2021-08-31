package util

import (
	"encoding/json"
	"log"
	"os"
)

func ReadJsonConfig(path string, conf interface{}) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(conf)
	if err != nil {
		log.Fatal(err)
		return
	}
}
