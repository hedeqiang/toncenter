package services

import (
	"encoding/json"
	"fmt"
	toncenterHTTP "github.com/hedeqiang/toncenter/pkg/toncenter/http"
	"github.com/hedeqiang/toncenter/pkg/toncenter/models"
	"net/http"
)

type BlockchainService struct {
	client toncenterHTTP.Client
}

func NewBlockchainService(client toncenterHTTP.Client) *BlockchainService {
	return &BlockchainService{
		client: client,
	}
}

func (s *BlockchainService) GetBlocks(params map[string]interface{}) (*models.BlocksResponse, error) {
	endpoint := buildQueryString("blocks", params)

	respBody, err := s.client.Request(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result models.BlocksResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &result, nil
}
