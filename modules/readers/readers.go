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
	OutputPath string
}

func ReadSettings(path string) requests.AppConfig {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	byteData, _ := ioutil.ReadAll(f)

	var cfg requests.AppConfig
	json.Unmarshal(byteData, &cfg)

	return cfg
}

func GetArgs() Args {
	configPath := flag.String("config-path", "", "Path to the JSON file. (Required)")
	outputPath := flag.String("output-path", "", "Path to the output folder. (Default is the directory of the config-path.)")
	flag.Parse()

	var args Args
	if *configPath == "" {
		fmt.Println("config-path is required")
		os.Exit(1)
	} else {
		args.ConfigPath = *configPath
	}
	if *outputPath == "" {
		args.OutputPath = getLocatedFolder(*configPath)
	} else {
		args.OutputPath = *outputPath
	}

	return args
}

func getLocatedFolder(path string) string {
	lastSlashIdx := 0
	for i := range path {
		if path[i] == 47 {
			lastSlashIdx = i
		}
	}

	return path[:lastSlashIdx+1]
}
