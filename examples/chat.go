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

	clientID := "4214125"
	apiKey := "6271bcbb-4ada-410d-a43c-f4f1a6a44497"

	client, err := goozon.NewClient(clientID, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	req := models.ChatListRequest{
		Limit: 20,
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

	body := models.ChatHistoryRequest{
		ChatId: "7e2eae0c-cbe1-4b0d-a819-22c030aa6c73",
	}
	msg, err := client.GetHistory(ctx, body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)

}
