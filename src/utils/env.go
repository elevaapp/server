package utils

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	workingDirectory, _ := os.Getwd()
	basePath := strings.Split(workingDirectory, "src")[0]

	envMap, _ := godotenv.Read(basePath + "/.env")

	for _, key := range os.Environ() {
		variable := strings.SplitN(key, "=", 2)
		envMap[variable[0]] = variable[1]
	}

	return envMap[key]
}
