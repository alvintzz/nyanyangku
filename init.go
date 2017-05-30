package main

import (
	"fmt"
	"flag"
	"log"

	"gopkg.in/gcfg.v1"

	"github.com/alvintzz/nyanyangku/common/database"
	"github.com/alvintzz/nyanyangku/common/render"
)

var environment string
var config Configs
var masterDB *database.Db

func init() {
	var err error

	//Setting flags used for the apps.
	//Environment will determine which config file will be used. The default is development
	flag.StringVar(&environment, "env", "development", "Set Environment of apps. Default is development.")
	flag.Parse()

	//Read config based on environment provided on deployment
	config, err = readConfig(environment)
	if err != nil {
		log.Fatal("Failed to read config. Error:", err)
	}

	//Set log level to [2009/01/23 01:23:23] and short file name
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//Initialize database connection
	masterDB, err = database.Connect(config.Databases.Type, config.Databases.Conn)
	if err != nil {
		log.Fatal("Failed to connect to database. Error:", err)
	}

	//Initialize Templating Engine
	err = render.Init("main", config.Settings.TemplateDir)
	if err != nil {
		log.Fatal("Failed to create main templating.")
	}
}

type Configs struct {
	Settings  ConfigSetting
	Databases ConfigDB
}
type ConfigSetting struct {
	SelfURL     string
	SelfPort    string
	PublicDir   string
	TemplateDir string
}
type ConfigDB struct {
	Conn string
	Type string
}

func readConfig(env string) (Configs, error) {
	//File name and list of config location. Production and Development may have different location
	configName := fmt.Sprintf("%s.ini", env)
	location := []string{
		"files/etc/nyanyangku/",
		"/etc/nyanyangku/",
	}

	//Loop each location and read the file. will stop after it found the file and successfully parsed it
	var success bool
	var config Configs
	for _, loc := range location {
		configLocation := fmt.Sprintf("%s%s", loc, configName)
		err := gcfg.ReadFileInto(&config, configLocation)
		if err != nil {
			log.Printf("Failed to read config in %s. Error: %s", configLocation, err.Error())
			continue
		}
		success = true
		break
	}

	//Return error if no config found or success parsed
	if success == false {
		return config, fmt.Errorf("Failed to read config. No matching log found.")
	}

	return config, nil
}
