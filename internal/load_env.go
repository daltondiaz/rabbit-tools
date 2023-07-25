package internal

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ParamAccess struct {
    Url string 
    Pass string
    User string
    Api string
}

func LoadEnvVariables(env string) ParamAccess {

    err := godotenv.Load("./config.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    url := "RABBIT_URL"
    user := "RABBIT_USER"
    pass := "RABBIT_PASS"
    api := "RABBIT_API"
    if env != "" {
        url = env + "_" + url
        user = env + "_" + user 
        pass = env + "_" + pass 
        api = env + "_" + api
    }
    var params ParamAccess
    params.Url = os.Getenv(url)
    params.User = os.Getenv(user)
    params.Pass = os.Getenv(pass) 
    params.Api = os.Getenv(api) 
    return params
}
