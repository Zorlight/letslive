package config

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"sen1or/letslive/transcode/pkg/discovery"
	"sen1or/letslive/transcode/pkg/logger"

	"gopkg.in/yaml.v3"
)

type Service struct {
	Name            string `yaml:"name"`
	Hostname        string `yaml:"hostname"`
	APIPort         int    `yaml:"apiPort"`
	RtmpBindAddress string `yaml:"rtmpBindAddress"`
	Port            int    `yaml:"port"`
}

type RTMP struct {
	Port int `yaml:"port"`
}

type IPFS struct {
	Enabled           bool     `yaml:"enabled"`
	Gateway           string   `yaml:"gateway"` // the gateway address, it is used to generate the final url to the ipfs file
	SubGateways       []string `yaml:"subGateways"`
	BootstrapNodeAddr string   `yaml:"bootstrapNodeAddr"`
}

type MinIO struct {
	Enabled    bool   `yaml:"enabled"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	BucketName string `yaml:"bucketName"`
	ReturnURL  string `yaml:"returnURL"`
}

type Transcode struct {
	PublicHLSPath  string `yaml:"publicHLSPath"`
	PrivateHLSPath string `yaml:"privateHLSPath"`
	FFMpegSetting  struct {
		FFMpegPath     string `yaml:"ffmpegPath"`
		MasterFileName string `yaml:"masterFileName"`
		HLSTime        int    `yaml:"hlsTime"`
		CRF            int    `yaml:"crf"`
		Preset         string `yaml:"preset"`
		HlsListSize    int    `yaml:"hlsListSize"`
		HlsMaxSize     int    `yaml:"hlsMaxSize"`
		Qualities      []struct {
			Resolution string `yaml:"resolution"`
			MaxBitrate string `yaml:"maxBitrate"`
			FPS        int    `yaml:"fps"`
			BufSize    string `yaml:"bufSize"`
		} `yaml:"qualities"`
	} `yaml:"ffmpegSetting"`
}

type Config struct {
	Service   `yaml:"service"`
	RTMP      `yaml:"rtmp"`
	Transcode `yaml:"transcode"`
	IPFS      `yaml:"ipfs"`
	MinIO     `yaml:"minio"`
	Webserver struct {
		Port int `yaml:"port"`
	} `yaml:"webserver"`
}

func RetrieveConfig(registry discovery.Registry) *Config {
	config, err := retrieveConfig(registry)
	if err != nil {
		logger.Panicf("failed to get config: %s", err)
	}

	return config
}

func retrieveConfig(registry discovery.Registry) (*Config, error) {
	configserverURL, err := registry.ServiceAddress(context.Background(), "configserver")
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(
		"%s/%s-%s.yml",
		configserverURL,
		"transcode_service",
		os.Getenv("CONFIG_SERVER_PROFILE"),
	)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, fmt.Errorf("error while creating request: %s", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to config server: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(body, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return &config, nil
}
