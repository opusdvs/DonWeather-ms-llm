package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/opusdvs/DonWeather-ms-ollama/internal/delivery"
	"github.com/opusdvs/DonWeather-ms-ollama/internal/provider"
	"github.com/opusdvs/DonWeather-ms-ollama/internal/usecase"
)

func main() {
	appCtx, appCancel := context.WithCancel(context.Background())
	defer appCancel()
	ollamaApiUrl := os.Getenv("OLLAMA_API_URL")
	if ollamaApiUrl == "" {
		log.Fatal("OLLAMA_API_URL is not set")
	}
	ollamaModel := os.Getenv("OLLAMA_MODEL")
	if ollamaModel == "" {
		log.Fatal("OLLAMA_MODEL is not set")
	}
	yandexApiKey := os.Getenv("YANDEX_API_KEY")
	if yandexApiKey == "" {
		log.Fatal("YANDEX_API_KEY is not set")
	}
	yandexApiAgentUrl := os.Getenv("YANDEX_API_AGENT_URL")
	if yandexApiAgentUrl == "" {
		log.Fatal("YANDEX_API_AGENT_URL is not set")
	}
	yandexPromptId := os.Getenv("YANDEX_PROMPT_ID")
	if yandexPromptId == "" {
		log.Fatal("YANDEX_PROMPT_ID is not set")
	}
	yandexProjectId := os.Getenv("YANDEX_PROJECT_ID")
	if yandexProjectId == "" {
		log.Fatal("YANDEX_PROJECT_ID is not set")
	}

	tipProvider := provider.NewTipProvider(yandexApiAgentUrl, yandexPromptId, yandexProjectId, yandexApiKey)
	tipService := usecase.NewTipService(tipProvider)
	tipHandler := delivery.NewTipHandler(*tipService)

	tipsMux := http.NewServeMux()
	tipsMux.HandleFunc("/api/v1/tips/get-tips", tipHandler.GetTip)

	mainMux := http.NewServeMux()
	mainMux.Handle("/", tipsMux)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mainMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func(server *http.Server) {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
			return
		}
		fmt.Println("Server listen and serve success")
	}(server)

	fmt.Print("Server stardet on port 8080")
	<-appCtx.Done()

	fmt.Println("Server stopped")
	gfCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(gfCtx); err != nil {
		log.Fatal("Filed to shutdown server: %w", err)
	}
	fmt.Println("Server shutdown successfuly")

}
