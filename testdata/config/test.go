package config

import (
	"fmt"
	"os"
)

type UserData struct {
	Host string `envconfig:"USER_DATA_SERVICE_HOST" `
	Port string `envcofig:"USER_DATA_SERVICE_PORT" default:"8080"`
}

const (
	testEnvKey = "key"
)

func getENVFund() {
	val := os.Getenv(testEnvKey)
	fmt.Println(val)
}
