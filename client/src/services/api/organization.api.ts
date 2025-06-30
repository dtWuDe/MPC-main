import axios from 'axios';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:5001';

export interface Organization {
  id: string;
  name: string;
  domain?: string;
  api_key_prefix: string;
  status: string;
  plan: string;
  max_users: number;
  max_api_requests_per_month: number;
  settings?: Record<string, any>;
  created_at: string;
  updated_at: string;
}

export interface APIKey {
  id: string;
  organization_id: string;
  name: string;
  key_hash: string;
  permissions: string[];
  status: string;
  last_used_at?: string;
  expires_at?: string;
  created_at: string;
  updated_at: string;
}

export interface CreateAPIKeyRequest {
  name: string;
  permissions: string[];
  expires_at?: string;
}

export interface CreateAPIKeyResponse {
  api_key: APIKey;
  key: string; // Only returned once during creation
}

export interface UsageStats {
  total_requests: number;
  requests_this_month: number;
  average_response_time: number;
}

export interface CreateOrganizationRequest {
  name: string;
  domain?: string;
  plan: string;
  max_users: number;
  max_api_requests_per_month: number;
  settings?: Record<string, any>;
}

// Organization API service
export const organizationAPI = {
  // Get all organizations
  getOrganizations: async (page = 1, limit = 10): Promise<{ organizations: Organization[]; page: number; limit: number }> => {
    const response = await axios.get(`${API_URL}/organizations?page=${page}&limit=${limit}`);
    return response.data;
  },

  // Create organization
  createOrganization: async (data: CreateOrganizationRequest): Promise<Organization> => {
    const response = await axios.post(`${API_URL}/organizations`, data);
    return response.data;
  },

  // Get organization by ID
  getOrganization: async (id: string): Promise<Organization> => {
    const response = await axios.get(`${API_URL}/organizations/${id}`);
    return response.data;
  },

  // Get API keys for organization
  getAPIKeys: async (organizationId: string): Promise<APIKey[]> => {
    const response = await axios.get(`${API_URL}/organizations/${organizationId}/api-keys`);
    return response.data;
  },

  // Create API key
  createAPIKey: async (organizationId: string, data: CreateAPIKeyRequest): Promise<CreateAPIKeyResponse> => {
    const response = await axios.post(`${API_URL}/organizations/${organizationId}/api-keys`, data);
    return response.data;
  },

  // Delete API key
  deleteAPIKey: async (organizationId: string, apiKeyId: string): Promise<void> => {
    await axios.delete(`${API_URL}/organizations/${organizationId}/api-keys/${apiKeyId}`);
  },

  // Get usage stats
  getUsageStats: async (organizationId: string): Promise<UsageStats> => {
    const response = await axios.get(`${API_URL}/organizations/${organizationId}/usage`);
    return response.data;
  },
};

// Mock data for development
export const mockOrganizations: Organization[] = [
  {
    id: '1',
    name: 'Acme Corp',
    domain: 'acme.com',
    api_key_prefix: 'ACME1234',
    status: 'active',
    plan: 'professional',
    max_users: 100,
    max_api_requests_per_month: 100000,
    created_at: '2024-01-15T10:00:00Z',
    updated_at: '2024-01-15T10:00:00Z',
  },
  {
    id: '2',
    name: 'TechStart Inc',
    domain: 'techstart.io',
    api_key_prefix: 'TECH5678',
    status: 'active',
    plan: 'basic',
    max_users: 10,
    max_api_requests_per_month: 10000,
    created_at: '2024-02-01T10:00:00Z',
    updated_at: '2024-02-01T10:00:00Z',
  },
];

export const mockAPIKeys: APIKey[] = [
  {
    id: '1',
    organization_id: '1',
    name: 'Production API Key',
    key_hash: 'hash123',
    permissions: ['wallet:read', 'wallet:write', 'transaction:sign'],
    status: 'active',
    last_used_at: '2024-06-10T10:00:00Z',
    created_at: '2024-01-15T10:00:00Z',
    updated_at: '2024-01-15T10:00:00Z',
  },
  {
    id: '2',
    organization_id: '1',
    name: 'Development API Key',
    key_hash: 'hash456',
    permissions: ['wallet:read', 'transaction:sign'],
    status: 'active',
    last_used_at: '2024-06-09T10:00:00Z',
    created_at: '2024-02-01T10:00:00Z',
    updated_at: '2024-02-01T10:00:00Z',
  },
];

export const mockUsageStats: UsageStats = {
  total_requests: 15420,
  requests_this_month: 3241,
  average_response_time: 245.5,
}; 