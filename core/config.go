package core

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Config []Transfer

type Transfer struct {
	ReaderSettings Settings `json:"reader"`
	WriterSettings Settings `json:"writer"`
}

type Settings struct {
	Name string `json:"name"`
	Delay int `json:"delay"`
	Properties map[string]string `json:"properties"`
}

func (s *Settings) GetProperty(name string) (string, error) {
	if v, ok := s.Properties[name]; ok {
		return v, nil
	}
	return "", fmt.Errorf("Missing property: %s", name)
}

func LoadConfig(configFile string) (*Config, error) {
	bt, err := ioutil.ReadFile(configFile)
	if (err != nil) {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(bt, &config)
	if (err != nil) {
		return nil, err
	}
	return &config, nil
}
