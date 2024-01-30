package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type UpdateInfo struct {
	Content string
	Comment string
}

type NameServer interface {
	Update(info UpdateInfo)
}

type Auth interface {
}

type CloudflareServer struct {
	zoneId      string
	dnsRecordId string
	XAuthEmail  string
	XAuthKey    string
	DnsRecord   DnsRecord
	client      *http.Client
}

const reqBodyTemplate = `{"content": "%s","name": "%s","proxied": false,"type": "%s","comment": "%s","tags": [],"ttl": 1}`

func (s CloudflareServer) Update(info UpdateInfo) {
	reqBody := fmt.Sprintf(reqBodyTemplate, info.Content, cs.DnsRecord.Name, cs.DnsRecord.Type, info.Comment)
	url := fmt.Sprintf(baseUrl, cs.zoneId, cs.dnsRecordId)
	fmt.Println(url)
	fmt.Println(reqBody)
	req, err := http.NewRequest("PUT", url, strings.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("X-Auth-Email", s.XAuthEmail)
	// req.Header.Add("X-Auth-Key", s.XAuthKey)
	req.Header.Add("Authorization", "Bearer "+s.XAuthKey)
	if err != nil {
		fmt.Println(err.Error())
	}
	resp, err := s.client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	fmt.Println(string(data))
}

var cs CloudflareServer

const baseUrl = "https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s"

func GetCloudflareServer(config Config, client *http.Client) NameServer {
	cs = CloudflareServer{
		zoneId:      config.ZoneId,
		XAuthKey:    config.XAuthKey,
		XAuthEmail:  config.XAuthEmail,
		dnsRecordId: config.DnsRecordId,
		client:      client,
		DnsRecord:   config.DnsRecord,
	}
	// url := fmt.Sprintf(baseUrl, cs.zoneId, cs.dnsRecordId)
	// fmt.Println(url)
	return cs
}
