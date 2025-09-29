package main

import (
	"fmt"
	"os"
	"path"
)

const DICT_PATH = "dict"
const JMDICT_FILENAME = "JMdict_e.xml"

func main() {
	data, err := readDictionary(JMDICT_FILENAME)
	if err != nil {
		// TODO: Handle error properly
		panic(err)
	}

	fmt.Println(string(data))
}

func readDictionary(dictName string) ([]byte, error) {
	filename := path.Join(DICT_PATH, dictName)
	data, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}
