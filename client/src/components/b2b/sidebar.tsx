import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { 
  FiHome, 
  FiUsers, 
  FiKey, 
  FiBarChart, 
  FiSettings, 
  FiShield,
  FiActivity,
  FiDollarSign
} from 'react-icons/fi';

const B2BSidebar: React.FC = () => {
  const location = useLocation();

  const menuItems = [
    {
      path: '/b2b/dashboard',
      name: 'Dashboard',
      icon: FiHome,
    },
    {
      path: '/b2b/users',
      name: 'User Management',
      icon: FiUsers,
    },
    {
      path: '/b2b/api-keys',
      name: 'API Keys',
      icon: FiKey,
    },
    {
      path: '/b2b/analytics',
      name: 'Analytics',
      icon: FiBarChart,
    },
    {
      path: '/b2b/settings',
      name: 'Settings',
      icon: FiSettings,
    },
  ];

  const isActive = (path: string) => {
    return location.pathname === path || location.pathname.startsWith(path);
  };

  return (
    <div className="fixed left-0 top-0 h-full w-64 bg-white shadow-lg border-r border-gray-200 z-40">
      {/* Logo */}
      <div className="p-6 border-b border-gray-200">
        <div className="flex items-center">
          <div className="w-8 h-8 bg-blue-600 rounded-lg flex items-center justify-center mr-3">
            <FiShield className="w-5 h-5 text-white" />
          </div>
          <div>
            <h1 className="text-lg font-bold text-gray-900">MPC Business</h1>
            <p className="text-xs text-gray-500">Enterprise Portal</p>
          </div>
        </div>
      </div>

      {/* Navigation */}
      <nav className="p-4">
        <ul className="space-y-2">
          {menuItems.map((item) => {
            const Icon = item.icon;
            return (
              <li key={item.path}>
                <Link
                  to={item.path}
                  className={`flex items-center px-4 py-3 rounded-lg transition-colors ${
                    isActive(item.path)
                      ? 'bg-blue-50 text-blue-700 border-r-2 border-blue-700'
                      : 'text-gray-700 hover:bg-gray-50 hover:text-gray-900'
                  }`}
                >
                  <Icon className="w-5 h-5 mr-3" />
                  <span className="font-medium">{item.name}</span>
                </Link>
              </li>
            );
          })}
        </ul>
      </nav>

      {/* Quick Stats */}
      <div className="p-4 border-t border-gray-200 mt-auto">
        <div className="bg-gray-50 rounded-lg p-4">
          <h3 className="text-sm font-medium text-gray-900 mb-3">Quick Stats</h3>
          <div className="space-y-2">
            <div className="flex items-center justify-between text-sm">
              <span className="text-gray-600">Active Users</span>
              <span className="font-medium text-gray-900">1,250</span>
            </div>
            <div className="flex items-center justify-between text-sm">
              <span className="text-gray-600">API Requests</span>
              <span className="font-medium text-gray-900">45.6K</span>
            </div>
            <div className="flex items-center justify-between text-sm">
              <span className="text-gray-600">Transactions</span>
              <span className="font-medium text-gray-900">8.9K</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default B2BSidebar; 