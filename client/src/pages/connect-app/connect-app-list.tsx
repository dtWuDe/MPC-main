import { useEffect, useState } from "react";
import { FiKey, FiTrash2, FiCopy, FiEye, FiEyeOff, FiShield, FiServer, FiSettings } from "react-icons/fi";
import toast, { Toaster } from "react-hot-toast";

interface APIKey {
  _id: string;
  name: string;
  key: string;
  permissions: string[];
  status: 'active' | 'inactive' | 'suspended';
  createdAt: string;
  lastUsed: string;
  requestsCount: number;
}

interface MPCService {
  _id: string;
  name: string;
  type: 'wallet' | 'signing' | 'custody' | 'settlement';
  status: 'running' | 'stopped' | 'maintenance';
  endpoint: string;
  description: string;
}

export default function B2BAPIManagement() {
  // Mock data - trong thực tế sẽ từ API
  const [apiKeys, setApiKeys] = useState<APIKey[]>([
    {
      _id: '1',
      name: 'Production API Key',
      key: 'sk_live_4eC39HqLyjWDarjtT1zdp7dc',
      permissions: ['wallet:read', 'wallet:write', 'transaction:sign'],
      status: 'active',
      createdAt: '2024-01-15',
      lastUsed: '2024-06-10',
      requestsCount: 15420
    },
    {
      _id: '2', 
      name: 'Development API Key',
      key: 'sk_test_BQokikJOvBiI2HlWgH4olfQ2',
      permissions: ['wallet:read', 'transaction:sign'],
      status: 'active',
      createdAt: '2024-02-01',
      lastUsed: '2024-06-09',
      requestsCount: 3241
    }
  ]);

  const [mpcServices, setMpcServices] = useState<MPCService[]>([
    {
      _id: '1',
      name: 'Multi-Signature Wallet Service',
      type: 'wallet',
      status: 'running',
      endpoint: 'https://api.yourcompany.com/v1/mpc/wallet',
      description: 'Secure multi-party computation for wallet management'
    },
    {
      _id: '2',
      name: 'Transaction Signing Service', 
      type: 'signing',
      status: 'running',
      endpoint: 'https://api.yourcompany.com/v1/mpc/signing',
      description: 'Distributed transaction signing with threshold signatures'
    },
    {
      _id: '3',
      name: 'Digital Asset Custody',
      type: 'custody', 
      status: 'maintenance',
      endpoint: 'https://api.yourcompany.com/v1/mpc/custody',
      description: 'Enterprise-grade digital asset custody solution'
    }
  ]);

  const [isLoading, setIsLoading] = useState(false);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [activeTab, setActiveTab] = useState<'api-keys' | 'services'>('api-keys');
  const [hoveredKeyId, setHoveredKeyId] = useState<string | null>(null);
  const [visibleKeys, setVisibleKeys] = useState<Set<string>>(new Set());
  const [deleteKeyId, setDeleteKeyId] = useState<string | null>(null);

  const toggleKeyVisibility = (keyId: string) => {
    const newVisibleKeys = new Set(visibleKeys);
    if (newVisibleKeys.has(keyId)) {
      newVisibleKeys.delete(keyId);
    } else {
      newVisibleKeys.add(keyId);
    }
    setVisibleKeys(newVisibleKeys);
  };

  const copyToClipboard = (text: string) => {
    navigator.clipboard.writeText(text);
    toast.success('API Key đã được sao chép!');
  };

  const handleDeleteKey = (id: string) => {
    setDeleteKeyId(id);
  };

  const confirmDeleteKey = () => {
    if (deleteKeyId) {
      setApiKeys(apiKeys.filter(key => key._id !== deleteKeyId));
      toast.success('Đã xóa API Key thành công!');
      setDeleteKeyId(null);
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'active':
      case 'running':
        return 'bg-green-100 text-green-800';
      case 'inactive':
      case 'stopped':
        return 'bg-red-100 text-red-800';
      case 'suspended':
      case 'maintenance':
        return 'bg-yellow-100 text-yellow-800';
      default:
        return 'bg-gray-100 text-gray-800';
    }
  };

  const getServiceIcon = (type: string) => {
    switch (type) {
      case 'wallet':
        return <FiShield size={20} />;
      case 'signing':
        return <FiKey size={20} />;
      case 'custody':
        return <span style={{ width: 20, height: 20, display: 'inline-flex', alignItems: 'center', justifyContent: 'center' }}><FiServer size={20} /></span>;
      default:
        // return <FiSettings width={20} height={20} />;
        return null;
    }
  };

  return (
    <>
      <div className="p-6 border bg-white sm:m-4 sm:rounded-xl sm:shadow-lg min-h-screen">
        {/* Header */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-2">API Management</h1>
          <p className="text-gray-600">Quản lý API keys và MPC services cho doanh nghiệp của bạn</p>
        </div>

        {/* Navigation Tabs */}
        <div className="flex space-x-1 bg-gray-100 p-1 rounded-lg mb-6 w-fit">
          <button
            className={`px-4 py-2 rounded-md font-medium transition-colors ${
              activeTab === 'api-keys'
                ? 'bg-white text-blue-600 shadow-sm'
                : 'text-gray-600 hover:text-gray-900'
            }`}
            onClick={() => setActiveTab('api-keys')}
          >
            <FiKey  />
            API Keys
          </button>
          <button
            className={`px-4 py-2 rounded-md font-medium transition-colors ${
              activeTab === 'services'
                ? 'bg-white text-blue-600 shadow-sm'
                : 'text-gray-600 hover:text-gray-900'
            }`}
            onClick={() => setActiveTab('services')}
          >
            <FiServer  />
            MPC Services
          </button>
        </div>

        {/* API Keys Tab */}
        {activeTab === 'api-keys' && (
          <div>
            {apiKeys.length === 0 ? (
              <div className="text-center py-12">
                <span className="flex justify-center mb-4">
                  <FiKey size={48} />
                </span>
                <p className="text-gray-500 text-lg">Chưa có API key nào được tạo!</p>
              </div>
            ) : (
              <div className="space-y-4 mb-6">
                {apiKeys.map((apiKey) => (
                  <div
                    key={apiKey._id}
                    className="bg-gray-50 border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
                    onMouseEnter={() => setHoveredKeyId(apiKey._id)}
                    onMouseLeave={() => setHoveredKeyId(null)}
                  >
                    <div className="flex items-start justify-between">
                      <div className="flex-1">
                        <div className="flex items-center mb-2">
                          <h3 className="text-lg font-semibold text-gray-900 mr-3">
                            {apiKey.name}
                          </h3>
                          <span
                            className={`px-2 py-1 rounded-full text-xs font-medium ${getStatusColor(
                              apiKey.status
                            )}`}
                          >
                            {apiKey.status.toUpperCase()}
                          </span>
                        </div>
                        
                        <div className="flex items-center mb-3 bg-white border rounded-md p-3">
                          <code className="flex-1 text-sm font-mono text-gray-800">
                            {visibleKeys.has(apiKey._id)
                              ? apiKey.key
                              : '•'.repeat(apiKey.key.length)}
                          </code>
                          <div className="flex space-x-2 ml-3">
                            <button
                              onClick={() => toggleKeyVisibility(apiKey._id)}
                              className="text-gray-500 hover:text-gray-700"
                            >
                              {visibleKeys.has(apiKey._id) ? (
                                <FiEyeOff size={16} />
                              ) : (
                                <FiEye size={16}/>
                              )}
                            </button>
                            <button
                              onClick={() => copyToClipboard(apiKey.key)}
                              className="text-gray-500 hover:text-gray-700"
                            >
                              <span style={{ width: 16, height: 16, display: 'inline-flex', alignItems: 'center', justifyContent: 'center' }}>
                                <FiCopy size={16} />
                              </span>
                            </button>
                          </div>
                        </div>

                        <div className="grid grid-cols-1 md:grid-cols-3 gap-4 text-sm text-gray-600 mb-3">
                          <div>
                            <span className="font-medium">Tạo ngày:</span> {apiKey.createdAt}
                          </div>
                          <div>
                            <span className="font-medium">Sử dụng lần cuối:</span> {apiKey.lastUsed}
                          </div>
                          <div>
                            <span className="font-medium">Số requests:</span> {apiKey.requestsCount.toLocaleString()}
                          </div>
                        </div>

                        <div className="flex flex-wrap gap-2">
                          {apiKey.permissions.map((permission, index) => (
                            <span
                              key={index}
                              className="px-2 py-1 bg-blue-100 text-blue-800 text-xs rounded-md"
                            >
                              {permission}
                            </span>
                          ))}
                        </div>
                      </div>

                      <div
                        className={`transition-opacity duration-200 ${
                          hoveredKeyId === apiKey._id ? 'opacity-100' : 'opacity-0'
                        }`}
                      >
                        <button
                          className="text-red-600 hover:text-red-800 p-2"
                          onClick={() => handleDeleteKey(apiKey._id)}
                        >
                          <FiTrash2 size={20} />
                        </button>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            )}

            <div className="flex justify-center">
              <button
                className="bg-blue-600 hover:bg-blue-700 text-white font-semibold py-3 px-6 rounded-lg shadow-md transition-colors duration-300 flex items-center"
                onClick={() => setIsModalOpen(true)}
              >
                <FiKey size={16}  />
                Tạo API Key mới
              </button>
            </div>
          </div>
        )}

        {/* MPC Services Tab */}
        {activeTab === 'services' && (
          <div>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {mpcServices.map((service) => (
                <div
                  key={service._id}
                  className="bg-gradient-to-br from-gray-50 to-gray-100 border border-gray-200 rounded-xl p-6 hover:shadow-lg transition-all duration-300"
                >
                  <div className="flex items-center justify-between mb-4">
                    <div className="flex items-center">
                      <div className="p-2 bg-blue-100 rounded-lg mr-3">
                        {getServiceIcon(service.type)}
                      </div>
                      <div>
                        <h3 className="font-semibold text-gray-900">{service.name}</h3>
                        <p className="text-sm text-gray-500 capitalize">{service.type}</p>
                      </div>
                    </div>
                    <span
                      className={`px-2 py-1 rounded-full text-xs font-medium ${getStatusColor(
                        service.status
                      )}`}
                    >
                      {service.status.toUpperCase()}
                    </span>
                  </div>

                  <p className="text-gray-600 text-sm mb-4">{service.description}</p>

                  <div className="bg-white border rounded-md p-3 mb-4">
                    <div className="flex items-center justify-between">
                      <code className="text-xs text-gray-800 break-all">
                        {service.endpoint}
                      </code>
                      <button
                        onClick={() => copyToClipboard(service.endpoint)}
                        className="text-gray-500 hover:text-gray-700 ml-2"
                      >
                        <FiCopy size={16} />
                      </button>
                    </div>
                  </div>

                  <div className="flex space-x-2">
                    <button className="flex-1 bg-blue-600 hover:bg-blue-700 text-white text-sm py-2 px-3 rounded-md transition-colors">
                      Cấu hình
                    </button>
                    <button className="flex-1 bg-gray-200 hover:bg-gray-300 text-gray-800 text-sm py-2 px-3 rounded-md transition-colors">
                      Logs
                    </button>
                  </div>
                </div>
              ))}
            </div>
          </div>
        )}
      </div>

      {/* Delete Confirmation Modal */}
      {deleteKeyId && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 max-w-md w-full mx-4">
            <h3 className="text-lg font-semibold text-gray-900 mb-2">
              Xác nhận xóa API Key
            </h3>
            <p className="text-gray-600 mb-6">
              Bạn có chắc chắn muốn xóa API Key này? Hành động này không thể hoàn tác.
            </p>
            <div className="flex space-x-3">
              <button
                className="flex-1 bg-red-600 hover:bg-red-700 text-white py-2 px-4 rounded-md transition-colors"
                onClick={confirmDeleteKey}
              >
                Xóa
              </button>
              <button
                className="flex-1 bg-gray-200 hover:bg-gray-300 text-gray-800 py-2 px-4 rounded-md transition-colors"
                onClick={() => setDeleteKeyId(null)}
              >
                Hủy
              </button>
            </div>
          </div>
        </div>
      )}

      <Toaster position="top-center" />
    </>
  );
}