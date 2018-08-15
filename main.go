package main

import (
	"flag"
	"alerter/log"
	"os"
	"fmt"
	"github.com/BurntSushi/toml"
)

type Flags struct {
	ConfigPath string
}

var (
	flags         Flags
	config        Config
	configOutputs []string
)

func parseFlag() {
	flag.StringVar(&flags.ConfigPath, "config", "./config.toml", "-config /path/to/config.toml")
	flag.Parse()
}

func parseConfig() {
	if _, err := toml.DecodeFile(flags.ConfigPath, &config); err != nil {
		fmt.Println(err.Error())
	}
	configOutputs = append(configOutputs, fmt.Sprintf("[config] app mail notify: %s", config.App.MailNotify))
	configOutputs = append(configOutputs, fmt.Sprintf("[config] api key: %s", config.App.ApiKey))
	configOutputs = append(configOutputs, fmt.Sprintf("[config] mail server addr: %s", config.App.MailServerAddr))
	configOutputs = append(configOutputs, fmt.Sprintf("[config] server account: %s", config.App.ServerAccount))
	configOutputs = append(configOutputs, fmt.Sprintf("[config] mail list: %s", config.Users.MailList))
}

func init() {
	parseFlag()
	parseConfig()
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel("debug")

	for _, conf := range configOutputs {
		log.Info(conf)
	}

	server := New()
	server.OnInit()
	server.Run()
}
