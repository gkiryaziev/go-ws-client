package conf

import "testing"

func TestLoad(t *testing.T) {
	// read and parse yaml file
	config, err := NewConfig("../config.yaml").Load()
	if err != nil {
		t.Fatal(err)
	}

	// check parameters
	switch {
	case config.Debug != true && config.Debug != false:
		t.Error("Error, Debug =", config.Debug)
	case config.Server.Address == "":
		t.Error("Error, Address is empty.")
	}
}
