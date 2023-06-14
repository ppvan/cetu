package main

import (
	"flag"
	"fmt"
)

type Config struct {
	Domain string
	Port   string
	DSN    string
}

// ParseConfig parses the command line arguments and returns a Config struct.
func ParseConfig() Config {
	domain := flag.String("domain", "localhost", "Domain name")
	port := flag.String("port", "8080", "Port number")
	dsn := flag.String("dsn", "cetu:cetu@/cetu?parseTime=true", "MySQL data source name")
	flag.Parse()

	return Config{
		Domain: *domain,
		Port:   *port,
		DSN:    *dsn,
	}
}

func (config *Config) BaseURL() string {
	return fmt.Sprintf("%s:%s", config.Domain, config.Port)
}
