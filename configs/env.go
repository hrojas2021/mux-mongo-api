package configs

import (
	"os"
)

func EnvMongoURI() string {
	mongouri := os.Getenv("MONGOURI")
	if mongouri == "" {
		mongouri = "mongodb://localhost:27017"
	}
	return mongouri
}
