package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

var home, _ = os.UserHomeDir()
var configFile = home + "/.config/mangodl.conf"

type Config struct {
	Directory string `json:"directory"`
}

func check() {
	if runtime.GOOS == "windows" {
		configFile = home + "/mangodl.conf"
	}
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		os.Create(configFile)
	}
}
func ReadJSON() string {
	check()
	var config Config
	configData, _ := ioutil.ReadFile(configFile)
	json.Unmarshal(configData, &config)
	return config.Directory
}
func WriteJson(data string) {
	check()
	//check if the user has put an '/' which is required, otherwise add it to the string
	if data[len(data)-1] != '/' {
		data += "/"
	}
	_, err := ioutil.ReadFile(configFile)
	var config Config
	config.Directory = data
	newData, _ := json.MarshalIndent(config, "", "	")
	err = ioutil.WriteFile(configFile, newData, 0777)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Set default directory to", config.Directory)
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
	//write to the config struct and then write the struct to the file
	newData, _ := json.MarshalIndent(config, "", "	")
	err = ioutil.WriteFile(configFile, newData, 0777)
	if err != nil {
		log.Println(err)
	}
}
