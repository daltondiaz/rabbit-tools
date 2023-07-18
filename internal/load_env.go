package internal

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(env string) ParamAccess {

    err := godotenv.Load("./config.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    url := "RABBIT_URL"
    user := "RABBIT_USER"
    pass := "RABBIT_PASS"
    if env != "" {
        url = env + "_" + url
        user = env + "_" + user 
        pass = env + "_" + pass 
    }
    var params ParamAccess
    params.Url = os.Getenv(url)
    params.User = os.Getenv(user)
    params.Pass = os.Getenv(pass) 
    return params
}
