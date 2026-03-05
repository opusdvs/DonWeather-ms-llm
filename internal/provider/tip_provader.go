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
	apiAgentUrl string
	promptId    string
	projectId   string
	apiKey      string
	client      *http.Client
}

func NewTipProvider(apiAgentUrl string, promptId string, projectId string, apiKey string) *TipProvider {
	return &TipProvider{
		apiAgentUrl: apiAgentUrl,
		promptId:    promptId,
		projectId:   projectId,
		apiKey:      apiKey,
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
		Вероятность снега %f процентов.
		Текущая температура: %.1f градусов.
		Важно: при минусовой температуре (ниже 0 °C) не рекомендуй брать зонт — осадки скорее в виде снега. Зонт рекомендуй только при плюсовой температуре и вероятности дождя.
		Сформируй короткий совет на русском языке.
		Проанализируй погодные показатели и выдай краткий практический совет по одежде и активности.
		`, prediction.TempDelta, prediction.RainProbability, prediction.WindProbability, prediction.SnowProbability, prediction.Temperature)
	data := domain.TipRequest{
		Prompt: domain.Prompt{Id: p.promptId},
		Input:  inputString,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("failed to marshal data: %w", err)
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s", p.apiAgentUrl), bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("failed to create request: %w", err)
		return nil, err
	}
	req.Header.Set("x-Project-Id", p.projectId)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.apiKey))
	req.Header.Set("Content-Type", "application/json")
	log.Println("reqData", string(jsonData))
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
