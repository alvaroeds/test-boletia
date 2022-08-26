package config

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
)

const (
	defaultHttpPort = "8080"
)

// Config is the central setting.
type Config struct {
	CurrencyApiURL      string
	CurrencyApiKEy      string
	CurrencyApiInterval int
	CurrencyApiTimeout  int
	HttpPort            string
	PostgresURL         string
}

func GetConfig() *Config {
	httpport := os.Getenv("HTTP_PORT")
	if httpport == "" {
		httpport = defaultHttpPort
	}

	postgresURL := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_DATABASE"),
	)

	conf := &Config{
		CurrencyApiURL:      getEnvString("CURRENCY_API_URL", ""),
		CurrencyApiKEy:      getEnvString("CURRENCY_API_KEY", ""),
		CurrencyApiInterval: getEnvNum("INTERVAL_MINUTES_CURRENCY_API", 1),
		CurrencyApiTimeout:  getEnvNum("TIMEOUT_CURRENCY_API", 10),
		HttpPort:            getEnvString("HTTP_PORT", "8000"),
		PostgresURL:         postgresURL,
	}

	return conf
}

func getEnvString(key, defvalue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defvalue
}

func getEnvNum(key string, defvalue int) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return defvalue
	}

	return val
}
