package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type cache_size struct {
	GB   int64 `json:"GB"`
	MB   int64 `json:"MB"`
	KB   int64 `json:"KB"`
	Byte int64 `json:"Byte"`
}

type cors_config struct {
	AllowCredentials bool     `json:"AllowCredentials"`
	AllowOrigins     []string `json:"AllowOrigins"`
	AllowMethods     []string `json:"AllowMethods"`
	AllowHeaders     []string `json:"AllowHeaders"`
	ExposeHeaders    []string `json:"ExposeHeaders"`
}

type cors struct {
	Use    bool        `json:"use"`
	Config cors_config `json:"config"`
}

type server struct {
	Port    int   `json:"port"`
	Cors    cors  `json:"cors"`
	MaxFSMB int64 `json:"MaxFileSize-MB"`
}

type CDN_CONFIG_JSON struct {
	Cache_Dir  string     `json:"cache-dir"`
	Cache_Size cache_size `json:"cache-size"`
	Server     server     `json:"server"`
}

func NEW_CDN_CONFIG_JSON() *CDN_CONFIG_JSON {
	file, err := os.Open("cdn.config.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	var config *CDN_CONFIG_JSON
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalln(err)
	}
	return config
}

func (s cache_size) Size() int64 {
	var size int64 = 0
	size += s.Byte
	size += s.KB * 1024
	size += s.MB * 1024 * 1024
	size += s.GB * 1024 * 1024 * 1024
	return size
}
