package config

import "os"

type Config struct {
	InfluxBucket string
	InfluxOrg    string
	InfluxToken  string
	InfluxUrl    string
}

func LoadConfig() *Config {
	bucket, ok := os.LookupEnv("INFLUX_BUCKET")
	if !ok {
		bucket = "datacollector"
	}

	org, ok := os.LookupEnv("INFLUX_ORG")
	if !ok {
		org = "kabalov"
	}

	token, ok := os.LookupEnv("INFLUX_TOKEN")
	if !ok {
		token = "_-sZiVzOaxL2jYbQwxmCQx4Pc66OWanWZzvp2es5W8V5aj-qEv8-JAtSB-57ntpWg5jywkURKqzQjAm1wmSDlg=="
	}

	url, ok := os.LookupEnv("INFLUX_URL")
	if !ok {
		url = "http://localhost:8086"
	}

	config := Config{
		InfluxBucket: bucket,
		InfluxOrg:    org,
		InfluxToken:  token,
		InfluxUrl:    url,
	}

	return &config
}
