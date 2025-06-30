package handler

import (
	"mpc/internal/model"
	"mpc/internal/service"
	"mpc/pkg/errors"
	"mpc/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrganizationHandler struct {
	orgService *service.OrganizationService
}

func NewOrganizationHandler(orgService *service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{
		orgService: orgService,
	}
}

// CreateOrganization creates a new organization
// @Summary Create organization
// @Description Create a new organization for B2B wallet service
// @Tags organizations
// @Accept json
// @Produce json
// @Param organization body model.CreateOrganizationRequest true "Organization details"
// @Success 201 {object} model.Organization
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /organizations [post]
func (h *OrganizationHandler) CreateOrganization(c *gin.Context) {
	var req model.CreateOrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:     "Invalid request body",
			ErrorCode: "INVALID_REQUEST",
		})
		return
	}

	org, err := h.orgService.CreateOrganization(c.Request.Context(), &req)
	if err != nil {
		logger.Error("Failed to create organization", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:     "Failed to create organization",
			ErrorCode: "INTERNAL_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, org)
}

// GetOrganization gets organization details
// @Summary Get organization
// @Description Get organization details by ID
// @Tags organizations
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {object} model.Organization
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /organizations/{id} [get]
func (h *OrganizationHandler) GetOrganization(c *gin.Context) {
	orgID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:     "Invalid organization ID",
			ErrorCode: "INVALID_ID",
		})
		return
	}

	org, err := h.orgService.GetOrganizationByID(c.Request.Context(), orgID)
	if err != nil {
		if err == errors.ErrOrganizationNotFound {
			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Error:     "Organization not found",
				ErrorCode: "ORGANIZATION_NOT_FOUND",
			})
			return
		}
		logger.Error("Failed to get organization", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:     "Failed to get organization",
			ErrorCode: "INTERNAL_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, org)
}

// CreateAPIKey creates a new API key for an organization
// @Summary Create API key
// @Description Create a new API key for an organization
// @Tags api-keys
// @Accept json
// @Produce json
// @Param organization_id path string true "Organization ID"
// @Param api_key body model.CreateAPIKeyRequest true "API key details"
// @Success 201 {object} model.CreateAPIKeyResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /organizations/{organization_id}/api-keys [post]
func (h *OrganizationHandler) CreateAPIKey(c *gin.Context) {
	orgID, err := uuid.Parse(c.Param("organization_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:     "Invalid organization ID",
			ErrorCode: "INVALID_ID",
		})
		return
	}

	var req model.CreateAPIKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:     "Invalid request body",
			ErrorCode: "INVALID_REQUEST",
		})
		return
	}

	apiKeyData, apiKey, err := h.orgService.CreateAPIKey(c.Request.Context(), orgID, &req)
	if err != nil {
		logger.Error("Failed to create API key", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:     "Failed to create API key",
			ErrorCode: "INTERNAL_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, model.CreateAPIKeyResponse{
		APIKey: apiKeyData,
		Key:    apiKey, // Only returned once
	})
}

// GetAPIKeys gets all API keys for an organization
// @Summary Get API keys
// @Description Get all API keys for an organization
// @Tags api-keys
// @Produce json
// @Param organization_id path string true "Organization ID"
// @Success 200 {array} model.APIKey
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /organizations/{organization_id}/api-keys [get]
func (h *OrganizationHandler) GetAPIKeys(c *gin.Context) {
	orgID, err := uuid.Parse(c.Param("organization_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:     "Invalid organization ID",
			ErrorCode: "INVALID_ID",
		})
		return
	}

	apiKeys, err := h.orgService.GetAPIKeysByOrganization(c.Request.Context(), orgID)
	if err != nil {
		logger.Error("Failed to get API keys", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:     "Failed to get API keys",
			ErrorCode: "INTERNAL_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, apiKeys)
}

// DeleteAPIKey deletes an API key
// @Summary Delete API key
// @Description Delete an API key
// @Tags api-keys
// @Param organization_id path string true "Organization ID"
// @Param api_key_id path string true "API Key ID"
// @Success 204 "No Content"
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /organizations/{organization_id}/api-keys/{api_key_id} [delete]
func (h *OrganizationHandler) DeleteAPIKey(c *gin.Context) {
	orgID, err := uuid.Parse(c.Param("organization_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:     "Invalid organization ID",
			ErrorCode: "INVALID_ID",
		})
		return
	}

	apiKeyID, err := uuid.Parse(c.Param("api_key_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:     "Invalid API key ID",
			ErrorCode: "INVALID_ID",
		})
		return
	}

	err = h.orgService.DeleteAPIKey(c.Request.Context(), orgID, apiKeyID)
	if err != nil {
		if err == errors.ErrInvalidAPIKey {
			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Error:     "API key not found",
				ErrorCode: "API_KEY_NOT_FOUND",
			})
			return
		}
		logger.Error("Failed to delete API key", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:     "Failed to delete API key",
			ErrorCode: "INTERNAL_ERROR",
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetUsageStats gets usage statistics for an organization
// @Summary Get usage stats
// @Description Get usage statistics for an organization
// @Tags analytics
// @Produce json
// @Param organization_id path string true "Organization ID"
// @Success 200 {object} model.UsageStats
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /organizations/{organization_id}/usage [get]
func (h *OrganizationHandler) GetUsageStats(c *gin.Context) {
	orgID, err := uuid.Parse(c.Param("organization_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:     "Invalid organization ID",
			ErrorCode: "INVALID_ID",
		})
		return
	}

	stats, err := h.orgService.GetUsageStats(c.Request.Context(), orgID)
	if err != nil {
		logger.Error("Failed to get usage stats", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:     "Failed to get usage stats",
			ErrorCode: "INTERNAL_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, stats)
}

// GetOrganizations gets all organizations with pagination
// @Summary Get organizations
// @Description Get all organizations with pagination
// @Tags organizations
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} model.OrganizationListResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /organizations [get]
func (h *OrganizationHandler) GetOrganizations(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	organizations, err := h.orgService.GetOrganizations(c.Request.Context(), limit, offset)
	if err != nil {
		logger.Error("Failed to get organizations", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:     "Failed to get organizations",
			ErrorCode: "INTERNAL_ERROR",
		})
		return
	}

	// For simplicity, we'll return a basic response
	// In a real implementation, you'd want to get total count for pagination
	c.JSON(http.StatusOK, gin.H{
		"organizations": organizations,
		"page":          page,
		"limit":         limit,
	})
}
