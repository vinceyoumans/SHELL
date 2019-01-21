package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

const conf = "./shell.json"

//  this will retrieve the default from
//  shell.json file.
func GetShellDefaults() Config {
	var config Config
	configFile, err := os.Open(conf)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	config.Boom.BoomTime = config.Boom.BoomTime * 1000 * 1000 * 1000

	fmt.Println(config)
	return config

}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
