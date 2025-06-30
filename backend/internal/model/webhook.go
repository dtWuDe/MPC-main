package model

import (
	"time"

	"github.com/google/uuid"
)

type Webhook struct {
	ID             uuid.UUID `json:"id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	Name           string    `json:"name"`
	URL            string    `json:"url"`
	Events         []string  `json:"events"`
	Secret         string    `json:"secret"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type WebhookEvent struct {
	ID        uuid.UUID              `json:"id"`
	WebhookID uuid.UUID              `json:"webhook_id"`
	Event     string                 `json:"event"`
	Data      map[string]interface{} `json:"data"`
	Status    string                 `json:"status"`
	Attempts  int                    `json:"attempts"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

type CreateWebhookRequest struct {
	Name   string   `json:"name" validate:"required"`
	URL    string   `json:"url" validate:"required,url"`
	Events []string `json:"events" validate:"required,min=1"`
}

type WebhookResponse struct {
	Webhook Webhook `json:"webhook"`
}
