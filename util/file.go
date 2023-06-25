package util

import (
	"encoding/json"
	"fmt"
	"hackaton/store"
	"io/ioutil"
	"log"
	"os"
)

const (
	filename = "state.json"
)

func SaveToFile(state *store.Store) {
	stateJSON, err := json.Marshal(state)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filename, stateJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully saved to file")
}

func GetFromFile(state *store.Store) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(state)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully get from file")
}
