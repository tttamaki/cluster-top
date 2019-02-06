package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Tick       int    `yaml:"tick"`        // tick (in seconds) between receiving data
	Timeout    int    `yaml:"timeout"`     // threshold (in seconds) after a node is considered/displayed as offline
	RouterIp   string `yaml:"router_ip"`   // ip of cluster-smi-server
	MaxDisplay int    `yaml:"max_display"` // maximum number of display processes
	Ports      struct {
		Nodes   string `yaml:"nodes"`   // port of cluster-smi-server, which nodes send to
		Clients string `yaml:"clients"` // port of cluster-smi-server, where clients subscribe to
	} `yaml:"ports"`
	MinUsage  float32  `yaml:"minusage"`     // minimum cpu usage (%) to show process if usage is larger than MinUsage
}

func LoadConfig() Config {

	c := Config{}

	c.RouterIp = "10.30.81.234"
	c.Tick = 3
	c.MaxDisplay = 5
	c.Timeout = 180
	c.Ports.Nodes = "9082"
	c.Ports.Clients = "9083"
	c.MinUsage = 5

	if os.Getenv("CLUSTER_TOP_CONFIG_PATH") != "" {
		fn := os.Getenv("CLUSTER_TOP_CONFIG_PATH")

		yamlFile, err := ioutil.ReadFile(fn)
		if err != nil {
			log.Fatalf("Error: %v ", err)
		}

		err = yaml.Unmarshal(yamlFile, &c)
		if err != nil {
			log.Fatalf("Error in %s %v", fn, err)
		}
	}

	return c
}

func (c Config) Print() {

	if os.Getenv("CLUSTER_TOP_CONFIG_PATH") != "" {
		log.Println("Read configuration from", os.Getenv("CLUSTER_TOP_CONFIG_PATH"))
	} else {
		log.Println("use default configuration")
	}

	log.Println("  Tick:", c.Tick)
	log.Println("  Timeout:", c.Timeout)
	log.Println("  RouterIp:", c.RouterIp)
	log.Println("  Ports:")
	log.Println("    Nodes:", c.Ports.Nodes)
	log.Println("    Clients:", c.Ports.Clients)

}
