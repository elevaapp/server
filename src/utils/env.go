package utils

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	envMap, _ := godotenv.Read(".env")

	for _, key := range os.Environ() {
		variable := strings.SplitN(key, "=", 2)
		envMap[variable[0]] = variable[1]
	}

	return envMap[key]
}
