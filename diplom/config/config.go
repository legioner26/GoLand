package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Addr          string `yaml:"addr"`
	SMSFile       string `yaml:"sms_file"`
	MMSAddr       string `yaml:"mms_addr"`
	MMSFile       string `yaml:"mms_file"`
	VoiceCallFile string `yaml:"voice_call_file"`
	EmailFile     string `yaml:"email_file"`
	BillingFile   string `yaml:"billing_file"`
	SupportAddr   string `yaml:"support_addr"`
	SupportFile   string `yaml:"support_file"`
	IncidentAddr  string `yaml:"incident_addr"`
	IncidentFile  string `yaml:"incident_file"`
	WebDir        string `yaml:"web_dir"`
}

var GlobalConfig Config

func NewConfig(file string) Config {
	var config Config

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return GetDefaultConfig()
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println(err.Error())
		return GetDefaultConfig()
	}
	fmt.Println(config)
	return config
}

func GetDefaultConfig() Config {
	fmt.Println("get default config")

	//const dir = "/skillbox/diplom/cmd/"
	const dirsim = "/simulator/"
	const addr = ":9999"

	var config Config

	config.Addr = addr
	config.SMSFile = dirsim + "sms.data"
	config.MMSAddr = "http://" + dirsim + "/mms"
	config.MMSFile = dirsim + "mms.json"
	config.VoiceCallFile = dirsim + "voice.data"
	config.EmailFile = dirsim + "email.data"
	config.BillingFile = dirsim + "billing.data"
	config.SupportAddr = "http://" + addr + "/support"
	config.SupportFile = dirsim + "support.json"
	config.IncidentAddr = "http://" + addr + "/incident"
	config.IncidentFile = dirsim + "incident.json"
	config.WebDir = "/web"
	return config
}

func ForHerokuConfig(config Config) Config {
	port := "8383"
	config.MMSAddr = "http://127.0.0.1:" + port + "/mms"
	config.SupportAddr = "http://127.0.0.1:" + port + "/support"
	config.IncidentAddr = "http://127.0.0.1:" + port + "/incident"

	return config
}
