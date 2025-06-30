import React, { useState } from 'react';
import { 
  FiBell, 
  FiSearch, 
  FiUser, 
  FiLogOut, 
  FiSettings,
  FiChevronDown,
  FiShield
} from 'react-icons/fi';

const B2BHeader: React.FC = () => {
  const [showUserMenu, setShowUserMenu] = useState(false);
  const [showNotifications, setShowNotifications] = useState(false);

  const notifications = [
    {
      id: '1',
      title: 'New API Key Created',
      message: 'A new API key was created for your organization',
      time: '2 minutes ago',
      type: 'info'
    },
    {
      id: '2',
      title: 'High API Usage Alert',
      message: 'You\'ve used 80% of your monthly API quota',
      time: '1 hour ago',
      type: 'warning'
    },
    {
      id: '3',
      title: 'Transaction Completed',
      message: 'Successfully signed transaction: 0.5 ETH transfer',
      time: '3 hours ago',
      type: 'success'
    }
  ];

  const getNotificationIcon = (type: string) => {
    switch (type) {
      case 'success':
        return 'text-green-600 bg-green-100';
      case 'warning':
        return 'text-yellow-600 bg-yellow-100';
      case 'error':
        return 'text-red-600 bg-red-100';
      default:
        return 'text-blue-600 bg-blue-100';
    }
  };

  return (
    <header className="fixed top-0 left-64 right-0 bg-white shadow-sm border-b border-gray-200 z-30">
      <div className="flex items-center justify-between px-6 py-4">
        {/* Left side - Search */}
        <div className="flex items-center flex-1 max-w-md">
          <div className="relative w-full">
            <FiSearch className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-4 h-4" />
            <input
              type="text"
              placeholder="Search..."
              className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
        </div>

        {/* Right side - Notifications and User */}
        <div className="flex items-center space-x-4">
          {/* Notifications */}
          <div className="relative">
            <button
              onClick={() => setShowNotifications(!showNotifications)}
              className="relative p-2 text-gray-600 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition-colors"
            >
              <FiBell className="w-5 h-5" />
              <span className="absolute top-1 right-1 w-2 h-2 bg-red-500 rounded-full"></span>
            </button>

            {/* Notifications Dropdown */}
            {showNotifications && (
              <div className="absolute right-0 mt-2 w-80 bg-white rounded-lg shadow-lg border border-gray-200 z-50">
                <div className="p-4 border-b border-gray-200">
                  <h3 className="text-lg font-semibold text-gray-900">Notifications</h3>
                </div>
                <div className="max-h-96 overflow-y-auto">
                  {notifications.map((notification) => (
                    <div key={notification.id} className="p-4 border-b border-gray-100 hover:bg-gray-50">
                      <div className="flex items-start">
                        <div className={`p-2 rounded-full mr-3 ${getNotificationIcon(notification.type)}`}>
                          <FiShield className="w-4 h-4" />
                        </div>
                        <div className="flex-1">
                          <h4 className="text-sm font-medium text-gray-900">{notification.title}</h4>
                          <p className="text-sm text-gray-600 mt-1">{notification.message}</p>
                          <p className="text-xs text-gray-500 mt-2">{notification.time}</p>
                        </div>
                      </div>
                    </div>
                  ))}
                </div>
                <div className="p-4 border-t border-gray-200">
                  <button className="w-full text-sm text-blue-600 hover:text-blue-700 font-medium">
                    View All Notifications
                  </button>
                </div>
              </div>
            )}
          </div>

          {/* User Menu */}
          <div className="relative">
            <button
              onClick={() => setShowUserMenu(!showUserMenu)}
              className="flex items-center space-x-2 p-2 text-gray-600 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition-colors"
            >
              <div className="w-8 h-8 bg-blue-600 rounded-full flex items-center justify-center">
                <FiUser className="w-4 h-4 text-white" />
              </div>
              <div className="hidden md:block text-left">
                <p className="text-sm font-medium text-gray-900">John Doe</p>
                <p className="text-xs text-gray-500">Admin</p>
              </div>
              <FiChevronDown className="w-4 h-4" />
            </button>

            {/* User Dropdown */}
            {showUserMenu && (
              <div className="absolute right-0 mt-2 w-48 bg-white rounded-lg shadow-lg border border-gray-200 z-50">
                <div className="py-2">
                  <div className="px-4 py-2 border-b border-gray-100">
                    <p className="text-sm font-medium text-gray-900">John Doe</p>
                    <p className="text-xs text-gray-500">john.doe@company.com</p>
                  </div>
                  <button className="w-full flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">
                    <FiUser className="w-4 h-4 mr-3" />
                    Profile
                  </button>
                  <button className="w-full flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">
                    <FiSettings className="w-4 h-4 mr-3" />
                    Settings
                  </button>
                  <div className="border-t border-gray-100 mt-2 pt-2">
                    <button className="w-full flex items-center px-4 py-2 text-sm text-red-600 hover:bg-red-50">
                      <FiLogOut className="w-4 h-4 mr-3" />
                      Sign Out
                    </button>
                  </div>
                </div>
              </div>
            )}
          </div>
        </div>
      </div>

      {/* Organization Info */}
      <div className="px-6 py-2 bg-gray-50 border-t border-gray-200">
        <div className="flex items-center justify-between">
          <div className="flex items-center space-x-4">
            <span className="text-sm text-gray-600">Organization:</span>
            <span className="text-sm font-medium text-gray-900">Acme Corporation</span>
            <span className="px-2 py-1 text-xs bg-green-100 text-green-800 rounded-full">Professional Plan</span>
          </div>
          <div className="flex items-center space-x-4 text-sm text-gray-600">
            <span>API Usage: 45,678 / 100,000</span>
            <div className="w-24 bg-gray-200 rounded-full h-2">
              <div className="bg-blue-600 h-2 rounded-full" style={{ width: '45%' }}></div>
            </div>
          </div>
        </div>
      </div>
    </header>
  );
};

export default B2BHeader; 