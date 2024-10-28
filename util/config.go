package util

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/magiconair/properties"
)

type Config struct {
	prop *properties.Properties
}

const CONFIG_FILE = "config.conf"

func NewConfig() *Config {
	c := new(Config)

	configPath := os.Getenv("HOME")
	if configPath != "" && !strings.HasSuffix(configPath, "/") {
		configPath = configPath + "/" + CONFIG_FILE
	} else {
		configPath = configPath + CONFIG_FILE
	}
	log.Printf("Loading config file from %s", configPath)
	c.prop = properties.MustLoadFile(configPath, properties.UTF8)
	log.SetPrefix("Config : ")
	log.Println("flight_info_request=" + c.FlightInfoRequest())
	log.Println("mqtt_host=" + c.MQTTAddr())
	log.Println("mqtt_topic=" + c.MQTTTopic())
	log.Println("flight_info_key=" + c.FlightInfoKey())
	log.Println("redis_db_addr=" + c.RedisDBAddr())
	log.Println("rest_service_address=" + c.RestServiceAddress())
	log.Println("flight_fetch_interval_days=" + strconv.Itoa(c.RestFlightFetchIntervalDays()))

	return c
}

func (config *Config) FlightInfoRequest() string {
	return config.prop.GetString("flight_info_request", "http://localhost/flights")
}

func (config *Config) MQTTAddr() string {
	return config.prop.GetString("mqtt_addr", "localhost:1883")
}

func (config *Config) MQTTTopic() string {
	return config.prop.GetString("mqtt_topic", "gofly")
}

func (config *Config) FlightInfoKey() string {
	return config.prop.GetString("flight_info_key", "no_key")
}

func (config *Config) RedisDBAddr() string {
	return config.prop.GetString("redis_db_addr", "localhost:6379")
}

func (config *Config) RestServiceAddress() string {
	return config.prop.GetString("rest_service_address", ":8000")
}

func (config *Config) RestFlightFetchIntervalDays() int {
	return config.prop.GetInt("flight_fetch_interval_days", 1)
}
