package agent

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/victorzhou123/ai-agent/common/log"
)

type AgentService interface {
	Abstract(string) (string, error)
	Polish(string) (string, error)
}

type agent struct {
	agentCfg Config
}

func NewAgentService(cfg Config) AgentService {
	return &agent{
		agentCfg: cfg,
	}
}

type reqOpt struct {
	host     string
	port     string
	protocol string
	input    string
	prompt   string
	model    string
}

func (s *agent) requestAgent(opt *reqOpt) (AgentResp, error) {

	req, err := json.Marshal(newDefaultOllamaReq(opt.prompt, opt.input, opt.model))
	if err != nil {
		return AgentResp{}, fmt.Errorf("func genReq error, prompt is: %s", opt.prompt+opt.input)
	}

	// call service
	resp, err := http.Post(
		fmt.Sprintf("%s://%s:%s/api/chat", opt.protocol, opt.host, opt.port),
		"application/json",
		bytes.NewBuffer(req))
	if err != nil {
		return AgentResp{}, errors.New("cannot connect service")
	}
	defer resp.Body.Close()

	// unmarshal result
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("ReadAll failed, err is %s, body is %s", err, string(body))

		return AgentResp{}, errors.New("ReadAll failed")
	}

	log.Infof("response body is: %s", string(body))

	var agentResp AgentResp
	if err := json.Unmarshal(body, &agentResp); err != nil {
		return AgentResp{}, fmt.Errorf("unmarshal in requestAgent failed, body is: %s", string(body))
	}

	return agentResp, nil
}

func (s *agent) Abstract(input string) (string, error) {
	opt := reqOpt{
		host:     s.agentCfg.Client.Llm.Host,
		port:     s.agentCfg.Client.Llm.Port,
		protocol: s.agentCfg.Client.Llm.Protocol,
		input:    input,
		prompt:   s.agentCfg.Role.Abstract.Prompt,
		model:    s.agentCfg.Role.Abstract.Model,
	}

	resp, err := s.requestAgent(&opt)
	if err != nil {
		return "", err
	}

	if !resp.Done {
		log.Errorf("respone of agent not done, content is %s, opt is %s", resp.Message.Content, opt)

		return "", errors.New("response of agent not done, there may some problem")
	}

	return s.postProcess(resp.Message.Content, s.agentCfg.Role.Polish.Model), nil
}

func (s *agent) Polish(input string) (string, error) {
	opt := reqOpt{
		host:     s.agentCfg.Client.Llm.Host,
		port:     s.agentCfg.Client.Llm.Port,
		protocol: s.agentCfg.Client.Llm.Protocol,
		input:    input,
		prompt:   s.agentCfg.Role.Polish.Prompt,
		model:    s.agentCfg.Role.Polish.Model,
	}

	resp, err := s.requestAgent(&opt)
	if err != nil {
		return "", err
	}

	if !resp.Done {
		log.Errorf("respone of agent not done, content is %s, opt is %s", resp.Message.Content, opt)

		return "", errors.New("response of agent not done, there may some problem")
	}

	return s.postProcess(resp.Message.Content, s.agentCfg.Role.Polish.Model), nil
}

// postProcess: post process the response of agent
func (s *agent) postProcess(resp, model string) string {

	if strings.Contains(model, "deepseek") {
		// delete all content between double "\n\u003c/think\u003e\n"
		respArr := strings.Split(resp, "\n\u003c/think\u003e\n\n")
		if len(respArr) < 2 {
			return resp
		}

		log.Infof("response of deepseek agent is %s", respArr[1])

		return respArr[1]
	}

	return resp
}
