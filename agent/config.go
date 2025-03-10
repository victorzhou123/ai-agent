package agent

type Config struct {
	Client Client `json:"client" mapstructure:"client"`
	Role   Role   `json:"role" mapstructure:"role"`
}

type Client struct {
	Llm Llm `json:"llm" mapstructure:"llm"`
}

type Llm struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     string `json:"port" mapstructure:"port"`
	Protocol string `json:"protocol" mapstructure:"protocol"`
}

type Role struct {
	Abstract Setting `json:"abstract" mapstructure:"abstract"`
	Polish   Setting `json:"polish" mapstructure:"polish"`
}

type Setting struct {
	Model  string `json:"model" mapstructure:"model"`
	Prompt string `json:"prompt" mapstructure:"prompt"`
}
