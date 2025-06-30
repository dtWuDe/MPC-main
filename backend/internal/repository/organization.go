package repository

import (
	"context"
	"mpc/internal/model"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrganizationRepository struct {
	db *pgxpool.Pool
}

func NewOrganizationRepository(db *pgxpool.Pool) *OrganizationRepository {
	return &OrganizationRepository{
		db: db,
	}
}

// CreateOrganization creates a new organization
func (r *OrganizationRepository) CreateOrganization(ctx context.Context, org *model.Organization) error {
	query := `
		INSERT INTO organizations (id, name, domain, api_key_prefix, status, plan, max_users, max_api_requests_per_month, settings, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err := r.db.Exec(ctx, query,
		org.ID, org.Name, org.Domain, org.APIKeyPrefix, org.Status, org.Plan,
		org.MaxUsers, org.MaxAPIRequestsPerMonth, org.Settings, org.CreatedAt, org.UpdatedAt,
	)
	return err
}

// GetOrganizationByID retrieves an organization by ID
func (r *OrganizationRepository) GetOrganizationByID(ctx context.Context, orgID uuid.UUID) (*model.Organization, error) {
	query := `
		SELECT id, name, domain, api_key_prefix, status, plan, max_users, max_api_requests_per_month, settings, created_at, updated_at
		FROM organizations WHERE id = $1
	`

	var org model.Organization
	err := r.db.QueryRow(ctx, query, orgID).Scan(
		&org.ID, &org.Name, &org.Domain, &org.APIKeyPrefix, &org.Status, &org.Plan,
		&org.MaxUsers, &org.MaxAPIRequestsPerMonth, &org.Settings, &org.CreatedAt, &org.UpdatedAt,
	)
	if err != nil {
		return &model.Organization{}, err
	}
	return &org, nil
}

// CreateAPIKey creates a new API key
func (r *OrganizationRepository) CreateAPIKey(ctx context.Context, apiKey *model.APIKey) error {
	query := `
		INSERT INTO api_keys (id, organization_id, name, key_hash, permissions, status, expires_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.Exec(ctx, query,
		apiKey.ID, apiKey.OrganizationID, apiKey.Name, apiKey.KeyHash, apiKey.Permissions,
		apiKey.Status, apiKey.ExpiresAt, apiKey.CreatedAt, apiKey.UpdatedAt,
	)
	return err
}

// GetAPIKeyByHash retrieves an API key by its hash
func (r *OrganizationRepository) GetAPIKeyByHash(ctx context.Context, keyHash string) (*model.APIKey, error) {
	query := `
		SELECT id, organization_id, name, key_hash, permissions, status, last_used_at, expires_at, created_at, updated_at
		FROM api_keys WHERE key_hash = $1
	`

	var apiKey model.APIKey
	err := r.db.QueryRow(ctx, query, keyHash).Scan(
		&apiKey.ID, &apiKey.OrganizationID, &apiKey.Name, &apiKey.KeyHash, &apiKey.Permissions,
		&apiKey.Status, &apiKey.LastUsedAt, &apiKey.ExpiresAt, &apiKey.CreatedAt, &apiKey.UpdatedAt,
	)
	if err != nil {
		return &model.APIKey{}, err
	}
	return &apiKey, nil
}

// GetAPIKeyByID retrieves an API key by ID
func (r *OrganizationRepository) GetAPIKeyByID(ctx context.Context, apiKeyID uuid.UUID) (*model.APIKey, error) {
	query := `
		SELECT id, organization_id, name, key_hash, permissions, status, last_used_at, expires_at, created_at, updated_at
		FROM api_keys WHERE id = $1
	`

	var apiKey model.APIKey
	err := r.db.QueryRow(ctx, query, apiKeyID).Scan(
		&apiKey.ID, &apiKey.OrganizationID, &apiKey.Name, &apiKey.KeyHash, &apiKey.Permissions,
		&apiKey.Status, &apiKey.LastUsedAt, &apiKey.ExpiresAt, &apiKey.CreatedAt, &apiKey.UpdatedAt,
	)
	if err != nil {
		return &model.APIKey{}, err
	}
	return &apiKey, nil
}

// GetAPIKeysByOrganization gets all API keys for an organization
func (r *OrganizationRepository) GetAPIKeysByOrganization(ctx context.Context, orgID uuid.UUID) ([]*model.APIKey, error) {
	query := `
		SELECT id, organization_id, name, key_hash, permissions, status, last_used_at, expires_at, created_at, updated_at
		FROM api_keys WHERE organization_id = $1 ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query, orgID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apiKeys []*model.APIKey
	for rows.Next() {
		var apiKey model.APIKey
		err := rows.Scan(
			&apiKey.ID, &apiKey.OrganizationID, &apiKey.Name, &apiKey.KeyHash, &apiKey.Permissions,
			&apiKey.Status, &apiKey.LastUsedAt, &apiKey.ExpiresAt, &apiKey.CreatedAt, &apiKey.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		apiKeys = append(apiKeys, &apiKey)
	}
	return apiKeys, nil
}

// DeleteAPIKey deletes an API key
func (r *OrganizationRepository) DeleteAPIKey(ctx context.Context, apiKeyID uuid.UUID) error {
	query := `DELETE FROM api_keys WHERE id = $1`
	_, err := r.db.Exec(ctx, query, apiKeyID)
	return err
}

// UpdateAPIKeyLastUsed updates the last used timestamp for an API key
func (r *OrganizationRepository) UpdateAPIKeyLastUsed(ctx context.Context, apiKeyID uuid.UUID, lastUsed *time.Time) error {
	query := `UPDATE api_keys SET last_used_at = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.Exec(ctx, query, lastUsed, time.Now(), apiKeyID)
	return err
}

// CreateAPIUsage records API usage
func (r *OrganizationRepository) CreateAPIUsage(ctx context.Context, usage *model.APIUsage) error {
	query := `
		INSERT INTO api_usage (id, organization_id, api_key_id, endpoint, method, status_code, response_time_ms, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.Exec(ctx, query,
		usage.ID, usage.OrganizationID, usage.APIKeyID, usage.Endpoint, usage.Method,
		usage.StatusCode, usage.ResponseTimeMs, usage.CreatedAt,
	)
	return err
}

// GetAPIUsageByOrganization gets API usage count for an organization within a time range
func (r *OrganizationRepository) GetAPIUsageByOrganization(ctx context.Context, orgID uuid.UUID, startTime, endTime time.Time) (int64, error) {
	query := `
		SELECT COUNT(*) FROM api_usage 
		WHERE organization_id = $1 AND created_at >= $2 AND created_at <= $3
	`

	var count int64
	err := r.db.QueryRow(ctx, query, orgID, startTime, endTime).Scan(&count)
	return count, err
}

// GetAverageResponseTime gets average response time for an organization
func (r *OrganizationRepository) GetAverageResponseTime(ctx context.Context, orgID uuid.UUID) (float64, error) {
	query := `
		SELECT AVG(response_time_ms) FROM api_usage 
		WHERE organization_id = $1 AND response_time_ms IS NOT NULL
	`

	var avgTime float64
	err := r.db.QueryRow(ctx, query, orgID).Scan(&avgTime)
	if err != nil {
		return 0, err
	}
	return avgTime, nil
}

// CheckPrefixExists checks if an API key prefix already exists
func (r *OrganizationRepository) CheckPrefixExists(ctx context.Context, prefix string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM organizations WHERE api_key_prefix = $1)`

	var exists bool
	err := r.db.QueryRow(ctx, query, prefix).Scan(&exists)
	return exists, err
}

// GetOrganizations gets all organizations with pagination
func (r *OrganizationRepository) GetOrganizations(ctx context.Context, limit, offset int) ([]*model.Organization, error) {
	query := `
		SELECT id, name, domain, api_key_prefix, status, plan, max_users, max_api_requests_per_month, settings, created_at, updated_at
		FROM organizations ORDER BY created_at DESC LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var organizations []*model.Organization
	for rows.Next() {
		var org model.Organization
		err := rows.Scan(
			&org.ID, &org.Name, &org.Domain, &org.APIKeyPrefix, &org.Status, &org.Plan,
			&org.MaxUsers, &org.MaxAPIRequestsPerMonth, &org.Settings, &org.CreatedAt, &org.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		organizations = append(organizations, &org)
	}
	return organizations, nil
}

// UpdateOrganization updates an organization
func (r *OrganizationRepository) UpdateOrganization(ctx context.Context, org *model.Organization) error {
	query := `
		UPDATE organizations 
		SET name = $2, domain = $3, status = $4, plan = $5, max_users = $6, 
		    max_api_requests_per_month = $7, settings = $8, updated_at = $9
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		org.ID, org.Name, org.Domain, org.Status, org.Plan,
		org.MaxUsers, org.MaxAPIRequestsPerMonth, org.Settings, time.Now(),
	)
	return err
}
