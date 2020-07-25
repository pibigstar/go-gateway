package config

var Cluster clusterConfig

type clusterConfig struct {
	Ip      string `json:"ip"`
	Port    int    `json:"port"`
	SSLPort int    `json:"ssl_port"`
}
