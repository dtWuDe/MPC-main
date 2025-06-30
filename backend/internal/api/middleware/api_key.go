package middleware

import (
	"context"
	"fmt"
	"mpc/internal/model"
	"mpc/internal/service"
	"mpc/pkg/errors"
	"mpc/pkg/logger"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	APIKeyHeader = "X-API-Key"
	OrgIDHeader  = "X-Organization-ID"
)

type APIKeyMiddleware struct {
	organizationService *service.OrganizationService
}

func NewAPIKeyMiddleware(orgService *service.OrganizationService) *APIKeyMiddleware {
	return &APIKeyMiddleware{
		organizationService: orgService,
	}
}

// APIKeyAuth middleware for API key authentication
func (m *APIKeyMiddleware) APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip API key auth for certain endpoints
		if shouldSkipAPIKeyAuth(c.Request.URL.Path) {
			c.Next()
			return
		}

		apiKey := c.GetHeader(APIKeyHeader)
		if apiKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "API key is required",
				"code":  "MISSING_API_KEY",
			})
			return
		}

		// Validate API key and get organization
		org, apiKeyData, err := m.validateAPIKey(c.Request.Context(), apiKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid API key",
				"code":  "INVALID_API_KEY",
			})
			return
		}

		// Check organization status
		if org.Status != "active" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Organization is not active",
				"code":  "ORGANIZATION_INACTIVE",
			})
			return
		}

		// Check API key expiration
		if apiKeyData.ExpiresAt != nil && time.Now().After(*apiKeyData.ExpiresAt) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "API key has expired",
				"code":  "API_KEY_EXPIRED",
			})
			return
		}

		// Check rate limits
		if err := m.checkRateLimit(c.Request.Context(), org.ID); err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
				"code":  "RATE_LIMIT_EXCEEDED",
			})
			return
		}

		// Store organization and API key in context
		c.Set("organization_id", org.ID)
		c.Set("api_key_id", apiKeyData.ID)
		c.Set("api_permissions", apiKeyData.Permissions)

		// Update last used timestamp
		go m.updateLastUsed(c.Request.Context(), apiKeyData.ID)

		// Record API usage
		go m.recordAPIUsage(c.Request.Context(), org.ID, &apiKeyData.ID, c.Request.URL.Path, c.Request.Method, 200, nil)

		c.Next()
	}
}

// RequirePermission middleware to check specific permissions
func (m *APIKeyMiddleware) RequirePermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		permissions, exists := c.Get("api_permissions")
		if !exists {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "API permissions not found in context",
				"code":  "INTERNAL_ERROR",
			})
			return
		}

		perms, ok := permissions.([]string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid permissions format",
				"code":  "INTERNAL_ERROR",
			})
			return
		}

		hasPermission := false
		for _, p := range perms {
			if p == permission || p == "*" {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": fmt.Sprintf("Permission '%s' is required", permission),
				"code":  "INSUFFICIENT_PERMISSIONS",
			})
			return
		}

		c.Next()
	}
}

// GetOrganizationID helper to get organization ID from context
func GetOrganizationID(c *gin.Context) (uuid.UUID, error) {
	orgID, exists := c.Get("organization_id")
	if !exists {
		return uuid.Nil, errors.ErrOrganizationNotFound
	}

	if id, ok := orgID.(uuid.UUID); ok {
		return id, nil
	}

	return uuid.Nil, errors.ErrOrganizationNotFound
}

// GetAPIKeyID helper to get API key ID from context
func GetAPIKeyID(c *gin.Context) (uuid.UUID, error) {
	apiKeyID, exists := c.Get("api_key_id")
	if !exists {
		return uuid.Nil, errors.ErrAPIKeyNotFound
	}

	if id, ok := apiKeyID.(uuid.UUID); ok {
		return id, nil
	}

	return uuid.Nil, errors.ErrAPIKeyNotFound
}

func (m *APIKeyMiddleware) validateAPIKey(ctx context.Context, apiKey string) (*model.Organization, *model.APIKey, error) {
	// Use the real organization service to validate API key
	return m.organizationService.ValidateAPIKey(ctx, apiKey)
}

func (m *APIKeyMiddleware) checkRateLimit(ctx context.Context, orgID uuid.UUID) error {
	// Use the real organization service to check rate limits
	return m.organizationService.CheckRateLimit(ctx, orgID)
}

func (m *APIKeyMiddleware) updateLastUsed(ctx context.Context, apiKeyID uuid.UUID) {
	// Update the last_used_at timestamp for the API key
	// This should be done asynchronously to not block the request
}

func (m *APIKeyMiddleware) recordAPIUsage(ctx context.Context, orgID uuid.UUID, apiKeyID *uuid.UUID, endpoint, method string, statusCode int, responseTimeMs *int) {
	err := m.organizationService.RecordAPIUsage(ctx, orgID, apiKeyID, endpoint, method, statusCode, responseTimeMs)
	if err != nil {
		logger.Error("Failed to record API usage", err)
	}
}

func shouldSkipAPIKeyAuth(path string) bool {
	// Skip API key auth for these endpoints
	skipPaths := []string{
		"/api/v1/health",
		"/api/v1/auth/login",
		"/api/v1/auth/register",
		"/swagger",
		"/docs",
	}

	for _, skipPath := range skipPaths {
		if strings.HasPrefix(path, skipPath) {
			return true
		}
	}

	return false
}
