import React, { useState, useEffect } from 'react';
import { FiUsers, FiKey, FiActivity, FiDollarSign, FiTrendingUp, FiAlertCircle } from 'react-icons/fi';
import { Link } from 'react-router-dom';

interface DashboardStats {
  totalUsers: number;
  activeAPIKeys: number;
  monthlyRequests: number;
  totalTransactions: number;
  revenue: number;
  growthRate: number;
}

interface RecentActivity {
  id: string;
  type: 'api_call' | 'transaction' | 'user_signup' | 'wallet_created';
  description: string;
  timestamp: string;
  status: 'success' | 'error' | 'pending';
}

const B2BDashboard: React.FC = () => {
  const [stats, setStats] = useState<DashboardStats>({
    totalUsers: 0,
    activeAPIKeys: 0,
    monthlyRequests: 0,
    totalTransactions: 0,
    revenue: 0,
    growthRate: 0,
  });

  const [recentActivity, setRecentActivity] = useState<RecentActivity[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    // Mock data - replace with actual API calls
    setTimeout(() => {
      setStats({
        totalUsers: 1250,
        activeAPIKeys: 8,
        monthlyRequests: 45678,
        totalTransactions: 8923,
        revenue: 45600,
        growthRate: 23.5,
      });

      setRecentActivity([
        {
          id: '1',
          type: 'api_call',
          description: 'API call to /v1/wallets endpoint',
          timestamp: '2024-01-15T10:30:00Z',
          status: 'success',
        },
        {
          id: '2',
          type: 'transaction',
          description: 'Transaction signed: 0.5 ETH transfer',
          timestamp: '2024-01-15T10:25:00Z',
          status: 'success',
        },
        {
          id: '3',
          type: 'user_signup',
          description: 'New user registered: john.doe@company.com',
          timestamp: '2024-01-15T10:20:00Z',
          status: 'success',
        },
        {
          id: '4',
          type: 'wallet_created',
          description: 'New wallet created for organization',
          timestamp: '2024-01-15T10:15:00Z',
          status: 'success',
        },
      ]);

      setIsLoading(false);
    }, 1000);
  }, []);

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'success':
        return 'text-green-600 bg-green-100';
      case 'error':
        return 'text-red-600 bg-red-100';
      case 'pending':
        return 'text-yellow-600 bg-yellow-100';
      default:
        return 'text-gray-600 bg-gray-100';
    }
  };

  const getActivityIcon = (type: string) => {
    switch (type) {
      case 'api_call':
        return <FiKey className="w-4 h-4" />;
      case 'transaction':
        return <FiActivity className="w-4 h-4" />;
      case 'user_signup':
        return <FiUsers className="w-4 h-4" />;
      case 'wallet_created':
        return <FiDollarSign className="w-4 h-4" />;
      default:
        return <FiActivity className="w-4 h-4" />;
    }
  };

  if (isLoading) {
    return (
      <div className="p-6">
        <div className="animate-pulse">
          <div className="h-8 bg-gray-200 rounded w-1/4 mb-6"></div>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
            {[...Array(4)].map((_, i) => (
              <div key={i} className="bg-gray-200 h-32 rounded-lg"></div>
            ))}
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="p-6">
      {/* Header */}
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-900 mb-2">Business Dashboard</h1>
        <p className="text-gray-600">Monitor your organization's MPC wallet usage and performance</p>
      </div>

      {/* Stats Cards */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <div className="bg-white rounded-lg shadow p-6 border">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm font-medium text-gray-600">Total Users</p>
              <p className="text-2xl font-bold text-gray-900">{stats.totalUsers.toLocaleString()}</p>
            </div>
            <div className="p-3 bg-blue-100 rounded-full">
              <FiUsers className="w-6 h-6 text-blue-600" />
            </div>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow p-6 border">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm font-medium text-gray-600">Active API Keys</p>
              <p className="text-2xl font-bold text-gray-900">{stats.activeAPIKeys}</p>
            </div>
            <div className="p-3 bg-green-100 rounded-full">
              <FiKey className="w-6 h-6 text-green-600" />
            </div>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow p-6 border">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm font-medium text-gray-600">Monthly Requests</p>
              <p className="text-2xl font-bold text-gray-900">{stats.monthlyRequests.toLocaleString()}</p>
            </div>
            <div className="p-3 bg-purple-100 rounded-full">
              <FiActivity className="w-6 h-6 text-purple-600" />
            </div>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow p-6 border">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm font-medium text-gray-600">Total Transactions</p>
              <p className="text-2xl font-bold text-gray-900">{stats.totalTransactions.toLocaleString()}</p>
            </div>
            <div className="p-3 bg-yellow-100 rounded-full">
              <FiDollarSign className="w-6 h-6 text-yellow-600" />
            </div>
          </div>
        </div>
      </div>

      {/* Revenue and Growth */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
        <div className="bg-white rounded-lg shadow p-6 border">
          <h3 className="text-lg font-semibold text-gray-900 mb-4">Revenue Overview</h3>
          <div className="flex items-center justify-between">
            <div>
              <p className="text-3xl font-bold text-gray-900">${stats.revenue.toLocaleString()}</p>
              <p className="text-sm text-gray-600">Total revenue this month</p>
            </div>
            <div className="flex items-center text-green-600">
              <FiTrendingUp className="w-5 h-5 mr-1" />
              <span className="font-medium">+{stats.growthRate}%</span>
            </div>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow p-6 border">
          <h3 className="text-lg font-semibold text-gray-900 mb-4">Quick Actions</h3>
          <div className="space-y-3">
            <Link
              to="/b2b/api-keys"
              className="flex items-center p-3 bg-blue-50 rounded-lg hover:bg-blue-100 transition-colors"
            >
              <FiKey className="w-5 h-5 text-blue-600 mr-3" />
              <span className="font-medium text-blue-900">Manage API Keys</span>
            </Link>
            <Link
              to="/b2b/users"
              className="flex items-center p-3 bg-green-50 rounded-lg hover:bg-green-100 transition-colors"
            >
              <FiUsers className="w-5 h-5 text-green-600 mr-3" />
              <span className="font-medium text-green-900">Manage Users</span>
            </Link>
            <Link
              to="/b2b/analytics"
              className="flex items-center p-3 bg-purple-50 rounded-lg hover:bg-purple-100 transition-colors"
            >
              <FiActivity className="w-5 h-5 text-purple-600 mr-3" />
              <span className="font-medium text-purple-900">View Analytics</span>
            </Link>
          </div>
        </div>
      </div>

      {/* Recent Activity */}
      <div className="bg-white rounded-lg shadow border">
        <div className="p-6 border-b">
          <h3 className="text-lg font-semibold text-gray-900">Recent Activity</h3>
        </div>
        <div className="p-6">
          <div className="space-y-4">
            {recentActivity.map((activity) => (
              <div key={activity.id} className="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
                <div className="flex items-center">
                  <div className="p-2 bg-white rounded-lg mr-4">
                    {getActivityIcon(activity.type)}
                  </div>
                  <div>
                    <p className="font-medium text-gray-900">{activity.description}</p>
                    <p className="text-sm text-gray-600">
                      {new Date(activity.timestamp).toLocaleString()}
                    </p>
                  </div>
                </div>
                <span className={`px-2 py-1 rounded-full text-xs font-medium ${getStatusColor(activity.status)}`}>
                  {activity.status}
                </span>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

export default B2BDashboard; 