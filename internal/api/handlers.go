package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Enilsonn/Olhama_Service/internal/service"
	"github.com/Enilsonn/Olhama_Service/internal/utils"
)

type generateRequest struct {
	Message string `json:"message"`
}

type generateResponse struct {
	Output       string `json:"output"`
	TokensUsed   int    `json:"tokens_used"`
	TokensInput  int    `json:"tokens_input"`
	TokensOutput int    `json:"tokens_output"`
}

type Handler struct {
	iaService *service.IAService
}

func NewIAService(s *service.IAService) *Handler {
	return &Handler{
		iaService: s,
	}
}

func (h *Handler) GenerateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req generateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{
			"error":   true,
			"code":    "INVALID_REQUEST",
			"message": fmt.Sprintf("error to decode json request: %v", err),
		})
	}

	resp, err := h.iaService.Generate(context.Background(), req.Message)
	if err != nil {
		utils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{
			"error":   true,
			"code":    "UNSECCESSFULY_GENERATION",
			"message": fmt.Sprintf("error to generate response: %v", err),
		})
	}

	jsonResponse := generateResponse{
		Output:       resp.Output,
		TokensUsed:   resp.TokensUsed,
		TokensInput:  resp.TokensInput,
		TokensOutput: resp.TokensOutput,
	}

	utils.EncodeJson(w, r, http.StatusOK, jsonResponse)
}
