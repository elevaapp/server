package utils

import "github.com/joho/godotenv"

func GetEnv(key string) string {
	envMap, _ := godotenv.Read(".env")
	return envMap[key]
}
