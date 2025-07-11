import { ItemCurrencyLoading, ItemCurrency } from "./item_currency";
import VNDIcon from "../../assets/png/vnd_icon.jpg";
import USDIcon from "../../assets/png//usd_icon.png";
import ETHIcon from "../../assets/png//eth_icon.png";
import axios from "axios";
import { createAxios } from "../../config/axios.config";
import { useEffect, useState } from "react";

const Assets = ({ ...props }) => {
  const walletData = props.walletData;
  const [balance, setBalance] = useState<string>("0");

  useEffect(() => {
    const fetchBalance = async () => {
      try {
        const axiosInstance = createAxios();
        const res = await axiosInstance.get("/api/v1/wallets/balance");
        setBalance(res.data?.payload?.balance ?? 0);
      } catch (error) {
        console.error("Failed to fetch balance:", error);
      }
    };
    fetchBalance();
  }, []);

  return (
    <div class={``}>
      <h1 className="mt-2 font-semibold text-sm">Tài sản</h1>
      {props.isLoading ? (
        <>
          <ItemCurrencyLoading />
          <ItemCurrencyLoading />
          <ItemCurrencyLoading />
        </>
      ) : (
        <>
          <ItemCurrency
            image={VNDIcon}
            // item={walletData?.currencies[0]}
            symbol="VND"
            name="Vietnamese Dong"
          />
          <ItemCurrency
            image={USDIcon}
            // item={walletData?.currencies[1]}
            symbol="USD"
            name="US Dollar"
          />
          <ItemCurrency
            image={ETHIcon}
            // item={walletData?.currencies[2]}
            item={{ balance: balance }}
            symbol="ETH"
            name="Ethereum"
          />
        </>
      )}
    </div>
  );
};

export default Assets;
