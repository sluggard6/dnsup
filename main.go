package main

import (
	"fmt"
	"io"
	"net/http"
)

var client = &http.Client{}

func main() {
	s := "cmd %s"
	fmt.Println(fmt.Sprintf(s, "start"))
	var config = LoadConfig("conf/cfg.json")
	ip, err := getIp(config.IpGeter)
	if err != nil {
		fmt.Printf("can't get ip, %v", err)
		return
	}
	ns := GetCloudflareServer(config, client)
	ns.Update(UpdateInfo{
		Content: ip,
		Comment: "dnsup update record!",
	})
	// updateDns(config.Authkey, ip)

}

func getIp(url []string) (ip string, err error) {
	var data []byte
	for _, url := range config.IpGeter {
		var resp *http.Response
		resp, err = http.Get(url)
		if err != nil {
			fmt.Printf("get ip failed, %v\n", err)
			continue
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("read error, %v\n", err)
			continue
		}
		ip = string(data)
	}
	if len(ip) == 0 {
		return ip, err
	}
	return ip, nil
}
