package config

import (
	"fmt"
	"os"
	"strconv"
)

const Version = "0.0.1"

var (
	Port            int
	Env             string = "dev"
	Dsn             string
	Host            string
	TokenSecret     string
	TokenDuration   int64
	RefreshSecret   string
	RefreshDuration int64
)

func Init() {
	Port = stringToInt(os.Getenv("PORT"))
	Env = os.Getenv("ENV")
	Dsn = os.Getenv("DSN")
	Host = os.Getenv("HOST")
	TokenSecret = os.Getenv("TOKEN_SECRET")
	TokenDuration = stringToInt64(os.Getenv("TOKEN_DURATION"))
	RefreshSecret = os.Getenv("REFRESH_SECRET")
	RefreshDuration = stringToInt64(os.Getenv("REFRESH_DURATION"))
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to int")
	}
	return i
}

func stringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Error converting string to int64")
	}
	return i
}

func Addr() string {
	return fmt.Sprintf("%s:%d", Host, Port)
}

func IsDev() bool {
	return Env == "dev"
}
