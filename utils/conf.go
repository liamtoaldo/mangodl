package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

var home, _ = os.UserHomeDir()
var configFile = home + "/.config/mangodl.conf"

type Config struct {
	Directory string `json:"directory"`
	Output    string `json:"output"`
}

func check() {
	if runtime.GOOS == "windows" {
		configFile = home + "/mangodl.conf"
	}
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		os.Create(configFile)
		defaultJson()
	}
}
func ReadJSON() Config {
	check()
	var config Config
	configData, _ := ioutil.ReadFile(configFile)
	json.Unmarshal(configData, &config)
	return config
}
func WriteJson(dir string, out string) {
	check()
	//check if the user has put an '/' which is required, otherwise add it to the string
	if dir[len(dir)-1] != '/' {
		dir += "/"
	}
	_, err := ioutil.ReadFile(configFile)
	var config Config
	config.Directory = dir
	//check if output is one of the available output files, otherwise, put the default (img)
	//TODO if a new one is added, update this:
	if out != "pdf" && out != "cbz" {
		config.Output = "img"
	} else {
		config.Output = out
	}

	newData, _ := json.MarshalIndent(config, "", "	")
	err = ioutil.WriteFile(configFile, newData, 0777)
	if err != nil {
		log.Println(err)
	}
}
func defaultJson() {
	check()
	configData, err := ioutil.ReadFile(configFile)
	var config Config
	json.Unmarshal(configData, &config)
	if config.Directory != home+"/Downloaded Manga/" && config.Directory != "" {
		return
	}
	//if it's windows, then redirect the downloaded manga to the Desktop, for better usability
	if runtime.GOOS == "windows" {
		config.Directory = home + "/Desktop/Downloaded Manga/"
	} else {
		config.Directory = home + "/Downloaded Manga/"
	}
	config.Output = "img"
	//write to the config struct and then write the struct to the file
	newData, _ := json.MarshalIndent(config, "", "	")
	err = ioutil.WriteFile(configFile, newData, 0777)
	if err != nil {
		log.Println(err)
	}
}
