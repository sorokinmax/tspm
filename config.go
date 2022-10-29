package main

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// StickersData struct
type StickersData struct {
	Packtitle string `yaml:"packtitle"`
	Packname  string `yaml:"packname"`
	Thumbnail struct {
		File string `yaml:"file"`
	} `yaml:"thumbnail"`
	Headsticker struct {
		File   string `yaml:"file"`
		Emojis string `yaml:"emojis"`
	} `yaml:"headsticker"`
	Stickers []struct {
		File   string `yaml:"file"`
		Emojis string `yaml:"emojis"`
	} `yaml:"stickers"`
}

// Config struct
type Config struct {
	BotToken       string `split_words:"true"`
	OwnerID        int64  `split_words:"true"`
	PathToStickers string `split_words:"true"`
}

func ReadDataFile(cfg *StickersData, file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadConfigEnv(cfg *Config) {
	err := envconfig.Process("TSPM", cfg)
	if err != nil {
		log.Fatal(err)
	}
}
