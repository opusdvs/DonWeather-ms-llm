package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/opusdvs/DonWeather-ms-ollama/internal/domain"
)

type TipProvider struct {
	apiURL string
	model  string
	client *http.Client
}

func NewTipProvider(apiURL string, model string) *TipProvider {
	return &TipProvider{
		apiURL: apiURL,
		model:  model,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (p *TipProvider) GetTip(ctx context.Context, prediction domain.Prediction) (*domain.Tip, error) {
	var tip domain.Tip
	inputString := fmt.Sprintf(`Температура изменится на %f градусов.
		Вероятность дождя %f процентов.
		Вероятность ветра %f процентов.
		Сформируй короткий совет на русском языке.
		Проанализируй погодные показатели и выдай краткий практический совет по одежде и активности.
		`, prediction.TempDelta, prediction.RainProbability, prediction.WindProbability)
	data := domain.TipRequest{
		Promt: domain.Prompt{Id: "fvt65721dad78sop0dhj"},
		Input: inputString,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("failed to marshal data: %w", err)
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/generate", p.apiURL), bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("failed to create request: %w", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := p.client.Do(req)
	if err != nil {
		log.Println("failed to do request: %w", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to generate tip: %s", resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&tip)
	if err != nil {
		log.Println("failed to decode tip: %w", err)
		return nil, err
	}
	return &tip, nil
}
