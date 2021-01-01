package config

type BaseCfg struct {
	Mode     string `yaml:"mode"`
	Property string `yaml:"property"`
	Method   string `yaml:"method"`
	Resp     string `yaml:"resp"`
	Type     string `yaml:"type"`
	Split    string `yaml:split`
}

type ServerCfg struct {
	Port    string   `yaml:"port"`
	Domains []string `yaml:"domains"`
}

type Configuration struct {
	Normal Normal    `yaml:"normal"`
	Server ServerCfg `yaml:"server"`
}

type Normal map[string]BaseCfg
