package internal

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() ParamAccess {

    err := godotenv.Load("./config.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    var params ParamAccess
    params.Url = os.Getenv("RABBIT_URL")
    params.User = os.Getenv("RABBIT_USER")
    params.Pass = os.Getenv("RABBIT_PASS") 
    return params
}
