package readers

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"modules/requests"
)

type Args struct {
	ConfigPath string
}

type config = requests.AppConfig

func ReadSettings(path string) config {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	byteData, _ := ioutil.ReadAll(f)

	var cfg config
	json.Unmarshal(byteData, &cfg)

	return cfg
}

func GetArgs() Args {
	configPath := flag.String("config-path", "", "Path to the JSON file. (Required)")
	flag.Parse()

	var args Args
	if *configPath == "" {
		fmt.Println("config-path is required")
		os.Exit(1)
	} else {
		args.ConfigPath = *configPath
	}

	return args
}
