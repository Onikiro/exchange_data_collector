package config

import "os"

type Config struct {
	InfluxBucket string
	InfluxOrg    string
	InfluxToken  string
	InfluxUrl    string
}

func LoadConfig() Config {
	bucket, ok := os.LookupEnv("INFLUX_BUCKET")
	if !ok {
		bucket = ""
	}

	org, ok := os.LookupEnv("INFLUX_ORG")
	if !ok {
		org = ""
	}

	token, ok := os.LookupEnv("INFLUX_TOKEN")
	if !ok {
		token = ""
	}

	url, ok := os.LookupEnv("INFLUX_URL")
	if !ok {
		url = ""
	}

	config := Config{
		InfluxBucket: bucket,
		InfluxOrg:    org,
		InfluxToken:  token,
		InfluxUrl:    url,
	}

	return config
}
