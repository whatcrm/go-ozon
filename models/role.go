package models

import "time"

type RoleResponse struct {
	ExpiresAt time.Time `json:"expires_at"`
	Roles     []struct {
		Name    string   `json:"name"`
		Methods []string `json:"methods"`
	} `json:"roles"`
}
