import ItemTransaction from "../transaction/item_transaction";
import USDTIcon from "../../assets/svg/usdt.svg";
import { getTransactionPaginate } from "../../services/api/transaction.api";
import { useQuery } from "@tanstack/react-query";
import React from "react";
import EmptyTrans from "../../assets/png/empty_transaction.jpg";


interface props {
  wallet_address: string;
}

const RecentTransaction: React.FC<props> = ({ wallet_address }) => {
  const { data, isLoading } = useQuery({
    queryFn: async () => {
      const response = await getTransactionPaginate(wallet_address, 1, 5);
      return response.data.payload;
    },
    queryKey: ["recent-transaction"],
  });
  console.log("data", data);
  if (isLoading) {
    return (
      <>
        {Array(5)
          .fill(0)
          .map((_, index) => (
            <div key={index} className="animate-pulse py-1">
              <div className="flex justify-between items-center my-2 py-1 px-4 bg-gray-100 rounded-lg">
                <div className="flex items-center">
                  <div className="bg-gray-300 rounded-full h-10 w-10"></div>
                  <div className="ml-2">
                    <div className="bg-gray-300 h-4 w-40 mb-2 rounded"></div>
                    <div className="bg-gray-300 h-4 w-16 rounded"></div>
                  </div>
                </div>
                <div>
                  <div className="bg-gray-300 h-4 w-20 rounded"></div>
                </div>
              </div>
            </div>
          ))}
      </>
    );
  }
  return (
    <div>
      {!isLoading && data?.transactions?.length > 0 ? (
        data?.transactions?.map((item: any, key: any) => (
          <ItemTransaction
            item={item}
            wallet_address={wallet_address}
            key={key}
            icon={USDTIcon}
          />
        ))
      ) : (
        <>
          <h1 class={`text-center font-semibold`}>Chưa có giao dịch nào</h1>
          <img class={`w-40 mx-auto`} src={EmptyTrans}></img>
        </>
      )}
    </div>
  );
};

export default RecentTransaction;
