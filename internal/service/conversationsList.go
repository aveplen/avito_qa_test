package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aveplen/avito_qa_test/internal/models"
)

type ConversationService struct {
	token string
}

func NewConversationService(token string) *ConversationService {
	return &ConversationService{
		token: token,
	}
}

func (cs *ConversationService) GetConversationsList() ([]models.ConvWrapper, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://slack.com/api/conversations.list", nil)
	if err != nil {
		return nil, fmt.Errorf("can't form request to slack api: %w", err)
	}

	request.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("Bearer %s", cs.token)},
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("can't send request to slack api: %w", err)
	}
	defer response.Body.Close()

	statusCode := response.StatusCode
	if statusCode/100 != 2 {
		return nil, fmt.Errorf("api returned bad status code: %d", statusCode)
	}

	var res struct {
		Channels []models.ConvWrapper `json:"channels"`
	}

	if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("can't decode response body %w", err)
	}

	return res.Channels, nil
}
