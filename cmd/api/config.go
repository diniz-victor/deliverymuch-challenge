package main

import (
	"os"
)

//ConfigPuppy something
type ConfigPuppy struct {
	endpoint string
}

//ConfigGiphy something
type ConfigGiphy struct {
	endpoint string
	apiToken string
}

//Config something
type Config struct {
	puppy ConfigPuppy
	giphy ConfigGiphy
}

func setupConfig() Config {
	return Config{
		puppy: ConfigPuppy{
			endpoint: os.Getenv("PUPPY_URL"),
		},
		giphy: ConfigGiphy{
			endpoint: os.Getenv("GIPHY_URL"),
			apiToken: os.Getenv("GIPHY_API_TOKEN"),
		},
	}
}
