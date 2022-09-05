package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Rpc        string `json:"rpc"`
	ChainID    int    `json:"chainId"`
	PrivateKey string `json:"privateKey"`
}

func NewConfig(path string) (Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var c Config
	err = json.Unmarshal(data, &c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
