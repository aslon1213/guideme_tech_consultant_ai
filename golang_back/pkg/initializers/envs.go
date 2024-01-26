package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvs() error {
	fmt.Println("Loading envs...")
	return godotenv.Load()
}
