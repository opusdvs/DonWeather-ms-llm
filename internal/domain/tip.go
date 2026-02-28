package domain

type Tip struct {
	ID                 string                 `json:"id"`
	CreatedAt          float64                `json:"created_at"`
	Error              interface{}            `json:"error"`
	IncompleteDetails  interface{}            `json:"incomplete_details"`
	Instructions       string                 `json:"instructions"`
	Metadata           map[string]interface{} `json:"metadata"`
	Model              string                 `json:"model"`
	Object             string                 `json:"object"`
	Output             []TipsOutputItem       `json:"output"`
	ParallelToolCalls  bool                   `json:"parallel_tool_calls"`
	Temperature        float64                `json:"temperature"`
	ToolChoice         string                 `json:"tool_choice"`
	Tools              []interface{}          `json:"tools"`
	TopP               float64                `json:"top_p"`
	Background         bool                   `json:"background"`
	MaxOutputTokens    int                    `json:"max_output_tokens"`
	MaxToolCalls       interface{}            `json:"max_tool_calls"`
	PreviousResponseID interface{}            `json:"previous_response_id"`
	Prompt             TipsPrompt             `json:"prompt"`
	Reasoning          interface{}            `json:"reasoning"`
	ServiceTier        string                 `json:"service_tier"`
	Status             string                 `json:"status"`
	Text               interface{}            `json:"text"`
	TopLogprobs        interface{}            `json:"top_logprobs"`
	Truncation         interface{}            `json:"truncation"`
	Usage              TipsUsage              `json:"usage"`
	User               interface{}            `json:"user"`
	Valid              bool                   `json:"valid"`
	OutputParsed       interface{}            `json:"output_parsed"`
	OutputText         string                 `json:"output_text"`
}

type TipsOutputItem struct {
	ID      string            `json:"id"`
	Content []TipsContentItem `json:"content"`
	Role    string            `json:"role"`
	Status  string            `json:"status"`
	Type    string            `json:"type"`
	Valid   bool              `json:"valid"`
}

type TipsContentItem struct {
	Annotations []interface{} `json:"annotations"`
	Text        string        `json:"text"`
	Type        string        `json:"type"`
	Logprobs    interface{}   `json:"logprobs,omitempty"`
	Valid       bool          `json:"valid"`
	Parsed      interface{}   `json:"parsed"`
}

type TipsPrompt struct {
	ID        string      `json:"id"`
	Variables interface{} `json:"variables"`
	Version   interface{} `json:"version"`
	Valid     bool        `json:"valid"`
}

type TipsUsage struct {
	InputTokens         int                    `json:"input_tokens"`
	InputTokensDetails  map[string]interface{} `json:"input_tokens_details"`
	OutputTokens        int                    `json:"output_tokens"`
	OutputTokensDetails map[string]interface{} `json:"output_tokens_details"`
	TotalTokens         int                    `json:"total_tokens"`
	Valid               bool                   `json:"valid"`
}

type Prediction struct {
	TempDelta       float64 `json:"temp_delta"`
	RainProbability float64 `json:"rain_probability"`
	WindProbability float64 `json:"wind_probability"`
}

type TipRequest struct {
	Prompt Prompt `json:"prompt"`
	Input  string `json:"input"`
}

type Prompt struct {
	Id string `json:"id"`
}
