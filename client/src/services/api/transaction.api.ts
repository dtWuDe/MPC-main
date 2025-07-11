import { createAxios } from "../../config/axios.config";
export const getTransactionPaginate = async (wallet_address: string, page: number, limit: number) => {
  const axiosInstance = createAxios();
  return await axiosInstance.get(
    `/api/v1/transactions?wallet_address=${wallet_address}&page=${page}&page_size=${limit}`,//`/api/v1/transaction/get/transactions?page=${page}&page_limit=${limit}`,
    {
      withCredentials: true,
    }
  );
};

export const getTransactionDetails = async (id: string) => {
  const axiosInstance = createAxios();
  return await axiosInstance.get(
    `/api/v1/transaction/get/transaction/details/${id}`,
    {
      withCredentials: true,
    }
  );
};
