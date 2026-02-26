package domain

import "time"

type Tip struct {
	ID              string    `json:"id"`
	Model           string    `json:"model"`
	Response        string    `json:"response"`
	Done            bool      `json:"done"`
	DoneReason      string    `json:"done_reason"`
	Context         []int     `json:"context"`
	TotalDuration   int       `json:"total_duration"`
	LoadDuration    int       `json:"load_duration"`
	PromptEvalCount int       `json:"prompt_eval_count"`
	EvalCount       int       `json:"eval_count"`
	EvalDuration    int       `json:"eval_duration"`
	CreatedAt       time.Time `json:"created_at"`
}

type Prediction struct {
	TempDelta       float64 `json:"temp_delta"`
	RainProbability float64 `json:"rain_probability"`
	WindProbability float64 `json:"wind_probability"`
}

type TipRequest struct {
	Promt Prompt `json:"promt"`
	Input string `json:"input"`
}

type Prompt struct {
	Id string `json:"id"`
}
