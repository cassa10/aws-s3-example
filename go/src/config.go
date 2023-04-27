package main

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Aws AwsConfig
}

type AwsConfig struct {
	Bucket    string `split_words:"true" required:"true"`
	Region    string `split_words:"true" required:"true"`
	SecretId  string `split_words:"true" required:"true"`
	SecretKey string `split_words:"true" required:"true"`
}

func GetConfig() Config {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		panic(err.Error())
	}
	return config
}
