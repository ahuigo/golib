package llm

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestLiteLLM(t *testing.T) {
	LlmEndpoint := "https://api.litelm.com/v1/chat/completions"
	picPath := "testdata/test.jpg"
	apiKey := os.Getenv("LITELLM_API_KEY")
	if apiKey == "" {
		t.Skip("跳过测试：未设置 LITELLM_API_KEY 环境变量")
	}
	imageData, err := os.ReadFile(picPath)
	if err != nil {
		t.Fatal(err)
	}

	// Encode image to base64
	imgBase64 := base64.StdEncoding.EncodeToString(imageData)

	// Construct request for face validation
	promptText := `判断是否是人脸。 请只回答 "是" 或 "否"`

	requestPayload := APIRequest{
		Model: "gpt-4-turbo",
		Messages: []Message{
			{
				Role: "user",
				Content: []ContentPart{
					{Type: "text", Text: promptText},
					{Type: "image_url", ImageURL: &ImageURL{URL: "data:image/jpeg;base64," + imgBase64}},
				},
			},
		},
		MaxTokens:   50,
		Temperature: 0.1,
	}

	jsonData, err := json.Marshal(requestPayload)
	if err != nil {
		t.Log("JSON编码失败: %w", err)
	}

	// Send API request
	fmt.Println("🔄 正在验证人脸...")
	req, err := http.NewRequest("POST", LlmEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("创建请求失败: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("API请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		t.Fatalf("API请求失败，状态码: %d，响应内容: %s", resp.StatusCode, string(bodyBytes))
	}

	// Parse response
	var apiResponse APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		t.Fatalf("解析API响应失败: %v", err)
	}

	if len(apiResponse.Choices) == 0 || apiResponse.Choices[0].Message.Content == "" {
		t.Fatalf("错误：未收到有效的AI响应")
	}

	result := strings.TrimSpace(strings.ToLower(apiResponse.Choices[0].Message.Content))
	fmt.Printf("🤖 AI人脸验证结果: '%s'\n", result)

	isFase := strings.Contains(result, "是") || strings.Contains(result, "yes")
	t.Logf("人脸验证结果: %v", isFase)
}

// ImageURL 定义了图片URL的数据结构
type ImageURL struct {
	URL string `json:"url"`
}

// ContentPart 定义了消息内容的数据结构，可以是文本或图片
type ContentPart struct {
	Type     string    `json:"type"`
	Text     string    `json:"text,omitempty"`
	ImageURL *ImageURL `json:"image_url,omitempty"`
}

// Message 定义了发送给API的消息结构
type Message struct {
	Role    string        `json:"role"`
	Content []ContentPart `json:"content"`
}

// APIRequest 定义了调用AI API的请求体
type APIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	MaxTokens   int       `json:"max_tokens"`
	Temperature float64   `json:"temperature"`
}

// APIResponse 定义了从AI API接收的响应体
type APIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}
