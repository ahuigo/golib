package gpt4

import (
	"context"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"

	"github.com/spf13/viper"
)

type (
	Conf struct {
		Key      string
		EndPoint string
	}
)

var conf Conf

func init() {
	viper.AddConfigPath(os.Getenv("HOME") + "/.mom")
	viper.SetConfigName("gpt")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

}

func NewGptClient() (*azopenai.Client, error) {
	// keyCredential, err := azcore.NewKeyCredential(conf.Key)
	keyCredential := azcore.NewKeyCredential(conf.Key)
	// keyCredential, err := azidentity.NewDefaultAzureCredential(nil)
	// if err != nil {
	// 	return nil, err
	// }
	client, err := azopenai.NewClientWithKeyCredential(conf.EndPoint, keyCredential, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func TestGptAsk(t *testing.T) {
	qa := "golang 的uint8 与[]byte 区别是"    // 问题
	rule := "你是一个后端开发专家，请你为用户提供简洁、有效的回答" // 规则
	system := ""                         //一些上下方知识, doc, markdown文档等
	// messages := []azopenai.ChatMessage{
	// 	{Role: to.Ptr(azopenai.ChatRoleSystem), Content: to.Ptr(system + "\n" + rule)},
	// 	{Role: to.Ptr(azopenai.ChatRoleUser), Content: to.Ptr(qa)},
	// }
	messages := []azopenai.ChatRequestMessageClassification{
		&azopenai.ChatRequestSystemMessage{
			Content: to.Ptr(rule + "\n" + system),
		},
		&azopenai.ChatRequestSystemMessage{
			Content: to.Ptr(qa),
		},
	}
	gptClient, err := NewGptClient()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := gptClient.GetChatCompletions(context.TODO(), azopenai.ChatCompletionsOptions{
		Messages:       messages,
		DeploymentName: to.Ptr("gpt-4"),
	}, nil)

	if err != nil {
		t.Fatal(err)
	}
	msg := *resp.Choices[0].Message.Content
	t.Log(msg)

}
