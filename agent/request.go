package agent

const (
	roleSystem = "system"
	roleUser   = "user"
)

// request
type OllamaReq struct {
	Messages []Message `json:"message"`
	Model    string    `json:"model"`
	Options  Options   `json:"options"`
	Stream   bool      `json:"stream"`
}

func newDefaultOllamaReq(prompt, content, model string) OllamaReq {
	promptMsg := Message{
		Content: prompt,
		Role:    roleSystem,
	}

	userMsg := Message{
		Content: content,
		Role:    roleUser,
	}

	options := Options{}
	options.setDefault()

	return OllamaReq{
		Messages: []Message{promptMsg, userMsg},
		Model:    model,
		Options:  options,
		Stream:   false,
	}
}

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type Options struct {
	FrequencyPenalty float32 `json:"frequency_penalty"`
	PresencePenalty  float32 `json:"presence_penalty"`
	Temperature      float32 `json:"temperature"`
	TOPP             float32 `json:"top_p"`
}

func (opt *Options) setDefault() {
	opt.FrequencyPenalty = 0
	opt.PresencePenalty = 0
	opt.Temperature = 0.5
	opt.TOPP = 1
}

// response
type AgentResp struct {
	Message Message `json:"message"`
	Done    bool    `json:"done"`
}
