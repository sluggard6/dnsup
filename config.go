package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type DnsRecord struct {
	Name string
	Type string
}

type Config struct {
	ZoneId      string
	DnsRecordId string
	XAuthKey    string
	XAuthEmail  string
	IpGeter     []string
	DnsRecord   DnsRecord
}

var config = Config{
	// Authkey: "null",
	IpGeter: []string{"https://ip.myfile.live/ip"},
}

func LoadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v\n", err)
		return config
	}
	json.Unmarshal(data, &config)
	return config
}
