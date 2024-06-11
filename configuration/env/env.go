package env

import (
	"os"

	"github.com/joho/godotenv"
)

type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	DBUrl         string
}

func (c *conf) setDBUrl() {
	c.DBUrl = "host=" + c.DBHost + " port=" + c.DBPort + " user=" + c.DBUser + " password=" + c.DBPassword + " dbname=" + c.DBName + " sslmode=disable"
}

func LoadConfig(path string) (*conf, error) {
	if err := godotenv.Load(path); err != nil {
		return nil, err
	}
	c := &conf{}

	c.DBDriver = os.Getenv("DB_DRIVER")
	c.DBHost = os.Getenv("DB_HOST")
	c.DBPort = os.Getenv("DB_PORT")
	c.DBUser = os.Getenv("DB_USER")
	c.DBPassword = os.Getenv("DB_PASSWORD")
	c.DBName = os.Getenv("DB_NAME")
	c.WebServerPort = os.Getenv("WEB_SERVER_PORT")

	c.setDBUrl()

	return c, nil
}
