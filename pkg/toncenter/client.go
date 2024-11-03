package toncenter

import (
	httpClient "github.com/hedeqiang/toncenter/pkg/toncenter/http"
	"github.com/hedeqiang/toncenter/pkg/toncenter/services"
	"net/http"
	"time"
)

const (
	defaultBaseURL    = "https://toncenter.com/api/v3/"
	defaultTimeout    = 10 * time.Second
	defaultRetryCount = 3
	defaultRetryDelay = 1 * time.Second
)

// HTTPClient defines the interface for making HTTP requests
type HTTPClient interface {
	Request(method, endpoint string, body interface{}) ([]byte, error)
}

// Client is the TON Center API client
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
	retryCount int
	retryDelay time.Duration

	Blockchain *services.BlockchainService
	//Jetton     *services.JettonService
	//NFT        *services.NFTService
	//Wallet     *services.WalletService
}

// NewClient creates a new TON Center API client
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		baseURL: defaultBaseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
		retryCount: defaultRetryCount,
		retryDelay: defaultRetryDelay,
	}

	for _, opt := range opts {
		opt(c)
	}

	httpWrapper := httpClient.NewClient(
		c.baseURL,
		c.apiKey,
		c.httpClient,
		c.retryCount,
		c.retryDelay,
	)

	c.Blockchain = services.NewBlockchainService(httpWrapper)
	//c.Jetton = services.NewJettonService(httpWrapper)
	//c.NFT = services.NewNFTService(httpWrapper)
	//c.Wallet = services.NewWalletService(httpWrapper)

	return c
}
