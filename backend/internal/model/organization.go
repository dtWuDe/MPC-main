package model

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID                     uuid.UUID              `json:"id"`
	Name                   string                 `json:"name"`
	Domain                 *string                `json:"domain,omitempty"`
	APIKeyPrefix           string                 `json:"api_key_prefix"`
	Status                 string                 `json:"status"`
	Plan                   string                 `json:"plan"`
	MaxUsers               int                    `json:"max_users"`
	MaxAPIRequestsPerMonth int64                  `json:"max_api_requests_per_month"`
	Settings               map[string]interface{} `json:"settings,omitempty"`
	CreatedAt              time.Time              `json:"created_at"`
	UpdatedAt              time.Time              `json:"updated_at"`
}

type OrganizationMember struct {
	ID             uuid.UUID `json:"id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	UserID         uuid.UUID `json:"user_id"`
	Role           string    `json:"role"`
	Permissions    []string  `json:"permissions"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type APIKey struct {
	ID             uuid.UUID  `json:"id"`
	OrganizationID uuid.UUID  `json:"organization_id"`
	Name           string     `json:"name"`
	KeyHash        string     `json:"key_hash"`
	Permissions    []string   `json:"permissions"`
	Status         string     `json:"status"`
	LastUsedAt     *time.Time `json:"last_used_at,omitempty"`
	ExpiresAt      *time.Time `json:"expires_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type APIUsage struct {
	ID             uuid.UUID  `json:"id"`
	OrganizationID uuid.UUID  `json:"organization_id"`
	APIKeyID       *uuid.UUID `json:"api_key_id,omitempty"`
	Endpoint       string     `json:"endpoint"`
	Method         string     `json:"method"`
	StatusCode     int        `json:"status_code"`
	ResponseTimeMs *int       `json:"response_time_ms,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
}

// Request/Response models
type CreateOrganizationRequest struct {
	Name                   string                 `json:"name" validate:"required"`
	Domain                 *string                `json:"domain,omitempty"`
	Plan                   string                 `json:"plan" validate:"required,oneof=basic professional enterprise"`
	MaxUsers               int                    `json:"max_users" validate:"min=1,max=1000"`
	MaxAPIRequestsPerMonth int64                  `json:"max_api_requests_per_month" validate:"min=1000"`
	Settings               map[string]interface{} `json:"settings,omitempty"`
}

type UpdateOrganizationRequest struct {
	Name                   *string                 `json:"name,omitempty"`
	Domain                 *string                 `json:"domain,omitempty"`
	Status                 *string                 `json:"status,omitempty" validate:"omitempty,oneof=active suspended inactive"`
	Plan                   *string                 `json:"plan,omitempty" validate:"omitempty,oneof=basic professional enterprise"`
	MaxUsers               *int                    `json:"max_users,omitempty" validate:"omitempty,min=1,max=1000"`
	MaxAPIRequestsPerMonth *int64                  `json:"max_api_requests_per_month,omitempty" validate:"omitempty,min=1000"`
	Settings               *map[string]interface{} `json:"settings,omitempty"`
}

type CreateAPIKeyRequest struct {
	Name        string     `json:"name" validate:"required"`
	Permissions []string   `json:"permissions" validate:"required,min=1"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
}

type OrganizationResponse struct {
	Organization Organization `json:"organization"`
	MemberCount  int          `json:"member_count"`
	APIKeyCount  int          `json:"api_key_count"`
	UsageStats   UsageStats   `json:"usage_stats"`
}

type UsageStats struct {
	TotalRequests       int64   `json:"total_requests"`
	RequestsThisMonth   int64   `json:"requests_this_month"`
	AverageResponseTime float64 `json:"average_response_time"`
}

type OrganizationListResponse struct {
	Organizations []OrganizationResponse `json:"organizations"`
	Total         int                    `json:"total"`
	Page          int                    `json:"page"`
	PageSize      int                    `json:"page_size"`
	TotalPages    int                    `json:"total_pages"`
}

// API Key Response models
type CreateAPIKeyResponse struct {
	APIKey *APIKey `json:"api_key"`
	Key    string  `json:"key"` // Only returned once during creation
}
