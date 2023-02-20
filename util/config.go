package util

import (
	"github.com/magiconair/properties"
)

type Config struct {
	prop *properties.Properties
}

const CONFIG_FILE = "config.conf"

func NewConfig() *Config {
	c := new(Config)
	c.prop = properties.MustLoadFile("${HOME}/"+CONFIG_FILE, properties.UTF8)
	return c
}

func (config *Config) FlightInfoRequest() string {
	return config.prop.GetString("flight_info_request", "http://localhost/flights")
}

func (config *Config) MQTTAddr() string {
	return config.prop.GetString("mqtt_host", "localhost:1883")
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
