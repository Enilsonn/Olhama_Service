package service

import (
	"context"
	"fmt"

	"github.com/Enilsonn/Olhama_Service/internal/olhama"
)

type GenerateResponse struct {
	Output       string
	TokensUsed   int
	TokensInput  int
	TokensOutput int
}

type IAService struct {
	olhamaClient *olhama.Client
}

func NewIAService(oc *olhama.Client) *IAService {
	return &IAService{
		olhamaClient: oc,
	}
}

func (s *IAService) Generate(ctx context.Context, message string) (*GenerateResponse, error) {
	olhamaResponse, err := s.olhamaClient.Generate(ctx, message)
	if err != nil {
		return nil, fmt.Errorf("erro to generate request from the model")
	}

	// tokens que serão discontados no bd
	tokensUsed := olhamaResponse.PromptEvalCount + olhamaResponse.EvalCount

	return &GenerateResponse{
		Output:       olhamaResponse.Response,
		TokensUsed:   tokensUsed,
		TokensInput:  olhamaResponse.PromptEvalCount,
		TokensOutput: olhamaResponse.EvalCount,
	}, nil
}
