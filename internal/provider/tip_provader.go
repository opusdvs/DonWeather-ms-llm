package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func (p *TipProvider) GetTip(ctx context.Context, input domain.Prediction) (*domain.Tip, error) {
	var tip domain.Tip
	promt := fmt.Sprintf(`Температура изменится на %f градусов.
		Вероятность дождя %f процентов.
		Вероятность ветра %f процентов.
		Сформируй короткий совет на русском языке.
		Ответь только советом, без других комментариев.`, input.TempDelta, input.RainProbability, input.WindProbability)
	data := domain.TipRequest{
		Promt:  promt,
		Model:  p.model,
		Stream: false,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/generate", p.apiURL), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to generate tip: %s", resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&tip)
	if err != nil {
		return nil, err
	}
	return &tip, nil
}
