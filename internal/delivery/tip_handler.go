package delivery

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/opusdvs/DonWeather-ms-ollama/internal/domain"
	"github.com/opusdvs/DonWeather-ms-ollama/internal/usecase"
)

type TipHandler struct {
	tipService usecase.TipService
}

func NewTipHandler(tipService usecase.TipService) *TipHandler {
	return &TipHandler{tipService: tipService}
}

func (th *TipHandler) GetTip(w http.ResponseWriter, r *http.Request) {
	var prediction domain.Prediction
	if err := json.NewDecoder(r.Body).Decode(&prediction); err != nil {
		log.Println("failed to decode prediction: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tip, err := th.tipService.GetTip(r.Context(), prediction)
	if err != nil {
		log.Println("failed to get tip: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tip)
	w.WriteHeader(http.StatusOK)
}
