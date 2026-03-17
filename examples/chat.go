package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	goozon "github.com/whatcrm/go-ozon"
	"github.com/whatcrm/go-ozon/models"
)

func main() {

	clientID := ""
	apiKey := ""

	client, err := goozon.NewClient(clientID, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	req := models.ChatListRequest{
		Limit: 1,
	}

	chats, err := client.GetList(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	out, err := json.MarshalIndent(chats, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
