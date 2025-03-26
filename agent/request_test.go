package agent

import (
	"reflect"
	"testing"
)

const (
	prompt  = "this is prompt"
	content = "this is content"

	modelDeepSeek = "deepseek-r1:14b"
)

func Test_newDefaultOllamaReq(t *testing.T) {
	type args struct {
		prompt  string
		content string
		model   string
	}
	tests := []struct {
		name string
		args args
		want OllamaReq
	}{
		{
			name: "deepseek-r1:14b",
			args: args{
				prompt:  prompt,
				content: content,
				model:   modelDeepSeek,
			},
			want: OllamaReq{
				Messages: []Message{
					{prompt, roleSystem},
					{content, roleUser},
				},
				Model: modelDeepSeek,
				Options: Options{
					FrequencyPenalty: 0,
					PresencePenalty:  0,
					Temperature:      0.5,
					TOPP:             1,
				},
				Stream: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDefaultOllamaReq(tt.args.prompt, tt.args.content, tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDefaultOllamaReq() = %v, want %v", got, tt.want)
			}
		})
	}
}
