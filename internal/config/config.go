package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct{
	Db_url              string    `json:"db_url"`         
	Current_user_name   string    `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) error{
	cfg.Current_user_name = userName
	return write(*cfg)
}

func Read()(Config, error){
	fullPath, err := getConfigFilePath()
	if err!=nil{
		return Config{},err
	}

	file, err := os.Open(fullPath)
	if err!= nil{
		return Config{},err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err!=nil{
		return cfg, err
	}
	return cfg,nil
}

func getConfigFilePath() (string, error){
	home, err := os.UserHomeDir()
	if err!=nil{
		return "",err
	}
	fullPath := filepath.Join(home,configFileName)
	return fullPath,nil
}

func write(cfg Config) error{
	fullPath, err := getConfigFilePath()
	if err!=nil{
		return err
	}
	file,err := os.Create(fullPath)
	if err!= nil{
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err!=nil{
		return err
	}
	return nil
}

