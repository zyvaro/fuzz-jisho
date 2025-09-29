package main

import (
	"fmt"
)

// CAREFUL!! Some function in other files
// depend on this global variables
const DICT_PATH = "dict"
const JMDICT_FILENAME = "JMdict_e.xml"

func main() {
	jmdict := decodeJmdict()

	fmt.Println(jmdict)
}
