package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config []Transfer

type Transfer struct {
	ReaderSettings ReaderSettings `json:"reader"`
	WriterSettings WriterSettings `json:"writer"`
}

type WriterSettings struct {
	Name       string            `json:"name"`
	Properties map[string]string `json:"properties"`
}

type ReaderSettings struct {
	WriterSettings
	Delay string `json:"delay"`
}

func (s *WriterSettings) GetProperty(name string) (string, error) {
	if v, ok := s.Properties[name]; ok {
		return v, nil
	}
	return "", fmt.Errorf("Missing property: %s", name)
}

func LoadConfig(configFile string) (*Config, error) {
	bt, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(bt, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
