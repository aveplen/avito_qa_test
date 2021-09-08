package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aveplen/avito_qa_test/internal/models"
)

func ReadJSON(filename string) (*models.JSONMessages, error) {
	temp, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("can't read file %s: %w", filename, err)
	}

	var res *models.JSONMessages
	if err := json.Unmarshal(temp, &res); err != nil {
		return nil, fmt.Errorf("can't unmarshal temp: %w", err)
	}

	return res, nil
}
