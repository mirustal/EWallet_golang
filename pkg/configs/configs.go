package configs

type ConfigFiber struct {
	Type string  `yaml:"type"`
}

type ConfigMongoDB struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Database string `json:"database"`
	Collection string `json:"collection"`
}