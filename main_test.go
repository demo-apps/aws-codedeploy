package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestConfig(t *testing.T) {
	contentBytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		t.Errorf("The config file missing")
	}

	var c config
	if err = json.Unmarshal(contentBytes, &c); err != nil {
		t.Errorf("Error while parsing the config file")
	}
	
	if c.Greeting == "" {
		t.Errorf("The config is missing the Greeting key")
	}
}
