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

const SampleConfig = `{
	"Port": ,     // application port

  	"DBHostname": "", // mongo hostname
  	"DBPort": ,       // mongo port
  	"DBUser": "",     // mongo user
  	"DBPassword": "", // mongo password
  	"DbName": ""      // mongo database name
}`