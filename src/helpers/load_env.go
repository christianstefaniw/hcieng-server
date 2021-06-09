package helpers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(RootDir() + "/.env"); err != nil {
		fmt.Println(err)
	}
}
