package main

import (
	"encoding/xml"
	"os"
	"path"
)

type JMDict struct {
	XMLName xml.Name      `xml:"JMdict"`
	Entry   []JMDictEntry `xml:"entry"`
}

type JMDictEntry struct {
	XMLName    xml.Name             `xml:"entry"`
	Seq        int                  `xml:"ent_seq"`
	KanjiEle   JMDictKanjiElement   `xml:"k_ele"`
	ReadingEle JMDictReadingElement `xml:"r_ele"`
	Sense      JMDictSense          `xml:"sense"`
}

type JMDictKanjiElement struct {
	XMLName xml.Name `xml:"k_ele"`
	Keb     string   `xml:"keb"`
}

type JMDictReadingElement struct {
	XMLName xml.Name `xml:"r_ele"`
	Reb     string   `xml:"reb"`
}

type JMDictSense struct {
	XMLName xml.Name `xml:"sense"`
	Gloss   []string `xml:"gloss"`
}

// This function decode the JMDict xml file inside dict
// directory. It does depends on the global variables
// inside main file.
func decodeJmdict() *JMDict {
	file, err := os.Open(path.Join(DICT_PATH, JMDICT_FILENAME))
	if err != nil {
		panic(err)
	}
	decoder := xml.NewDecoder(file)
	decoder.Strict = false
	jmdict := JMDict{}
	err = decoder.Decode(&jmdict)
	if err != nil {
		panic(err)
	}

	return &jmdict
}
