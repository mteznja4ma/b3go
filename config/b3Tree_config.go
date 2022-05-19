package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type B3TreeConfig struct {
	ID          string                  `json:"id"`
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	Root        string                  `json:"root"`
	Properties  map[string]interface{}  `json:"properties"`
	Nodes       map[string]B3NodeConfig `json:"nodes"`
}

func LoadB3Tree(path string) (*B3TreeConfig, bool) {
	var t B3TreeConfig
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("LoadB3Tree Read Error:%v", err)
		return nil, false
	}
	err = json.Unmarshal(f, &t)
	if err != nil {
		log.Fatalf("LoadB3Tree Json Error:%v", err)
		return nil, false
	}

	return &t, true
}
