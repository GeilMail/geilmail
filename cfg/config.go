package cfg

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	SQLite SQLiteConfig
	IMAP   IMAPConfig `yaml:"IMAP"`
	SMTP   SMTPConfig
	TLS    TLSConfig
	HTTP   HTTPConfig
}

type IMAPConfig struct {
	ListenIP string
	Port     int
}

type SMTPConfig struct {
	ListenIP string
	Port     int
}

type TLSConfig struct {
	CertPath string
	KeyPath  string
}

type HTTPConfig struct {
	Listen string
}

func ReadConfig(cfgPath string) *Config {
	cfg := Config{}
	log.Println("Reading Configuration file...")
	buf, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
