package config

type Configuration struct {
	Port int

	DBHostname string
	DBPort int
	DBUser string
	DBPassword string
	DbName string
}

var Config = new(Configuration)