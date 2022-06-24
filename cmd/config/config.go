package config

import (
	"fmt"
	"github.com/jinzhu/configor"
)

// Settings Some of this field should be loaded from an environment file (.env)
type Settings struct {
	// APP environments
	Port string `default:"8080"`

	// DB environments
	DBName       string `default:"test"`
	DBIP         string `default:"localhost"`
	DBPort       string `default:"5435"`
	DBUser       string `default:"tester"`
	DBPass       string `default:"mySuperPass"`
	DBRetryCount int    `default:"1"`

	// Google Books API environments
	// GoogleAPIKey should be injected from environment variables, But as you requested, It can be taken from a query param
	GoogleHost   string `default:"https://www.googleapis.com/books/v1"`
	GoogleAPIKey string `default:"<YourGoogleApiKey>"`

	// KEY for auth
	SecretKey string `default:"myhash123"`

	AllowInsecureCert bool `default:"true"`

	DebugMode bool `default:"true"`
}

var Config = Settings{}

func init() {
	if err := configor.Load(&Config, "config.yml"); err != nil {
		fmt.Println("Error trying to load configuration", err.Error())
	}
}
