package main

import (
	"github.com/hedeqiang/toncenter/pkg/toncenter"
	"log"
	"time"
)

func main() {
	// Create a new client with custom options
	client := toncenter.NewClient(
		"YOUR_API_KEY",
		toncenter.WithTimeout(5*time.Second),
		toncenter.WithRetryCount(3),
	)

	info, err := client.Blockchain.GetBlocks(map[string]interface{}{
		"limit":  1,
		"offset": 0,
		"sort":   "desc",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(info.Blocks)

}
