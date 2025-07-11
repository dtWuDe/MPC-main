import React, { useState, useEffect } from "react";
import DrawerBottom from "../../components/transfer/drawer_security";
import { DataSendETH } from "../../types/transfer";
import { formatCurrency } from "../../utils/format_currency";
import { convertCurrency } from "../../utils/convert_currency";
import CurrencyInput from "react-currency-input-field";
import { symbolCurrency } from "../../utils/symbol_currency";
import avatar from "../../assets/png/default_avatar.jpg";
import HeaderTransfer from "../../components/header/header_transfer";
import axios from "axios";
import { createAxios } from "../../config/axios.config";

const InputAmount = ({ ...props }) => {
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
  console.log("props", props);
  const [dataSend, setDataSend] = useState<DataSendETH>({
    from_address: props.wallet.address,
    to_address: props.userData.address,
    amount: "",
    symbol: props.currencyData.currency,
    share_data: localStorage.getItem(`share_data_${props.wallet.user_id}`) || "",
    chain_id: 11155111,
    security_code: "",
  });
  const [error, setError] = useState<string | null>();
  const [isOpen, setIsOpen] = useState<boolean>(false);

  const handleChange = (event: React.ChangeEvent) => {
    setError(null);
    const { name, value } = event.target as HTMLInputElement;
    setDataSend({
      ...dataSend,
      [name]: value,
    });
  };

  const handleSubmit = () => {
    const amount = dataSend.amount;

    if (!amount) {
      return setError(`Vui lòng nhập số tiền`);
    }
    if (Number(amount) > Number(balance)) {
      return setError(`Số dư không đủ`);
    }
    // if (amount < 100) {
    //   return setError(
    //     `Số tiền chuyển tối thiểu là ${formatCurrency(
    //       convertCurrency(100, props.currencyData.currency),
    //       props.currencyData.currency
    //     )}`
    //   );
    // }
    // if (amount > 100000000) {
    //   return setError(
    //     `Số tiền chuyển tối đa là ${formatCurrency(100000000, "VND")}`
    //   );
    // }
    setIsOpen(!isOpen);
  };

  return (
    <div className="container-center">
      <HeaderTransfer
        onClick={() => {
          props.handleStepTransfer("search_user");
        }}
        title="Nhập số tiền"
      />
      <div className="mt-14">
        <CurrencyInput
          allowDecimals={true}
          decimalsLimit={8}
          id="input-example"
          name="input-name"
          placeholder={`0 ${props.currencyData.currency}`}
          defaultValue={0}
          maxLength={12}
          className={`border-0 mt-5 focus:outline-none text-center text-6xl w-full font-semibold bg-white ${
            error && "text-red-500"
          }`}
          onValueChange={(value) => {
            setDataSend({ ...dataSend, amount: value || "0" });
            setError(null);
          }}
        />
        <textarea
          onChange={handleChange}
          name="message"
          maxLength={100}
          className="border mt-10 rounded-xl w-full h-24 p-4"
          placeholder="Nhập nội dung đính kèm"
        ></textarea>
      </div>
      <div className="rounded-xl bg-gray-100 p-4 flex">
        <img
          className="w-10 h-10 rounded-full object-cover"
          src={props.currencyData.image}
          alt="Currency"
        />
        <div className="ml-4">
          <h1 className="text-sm text-gray-500">Số dư:</h1>
          <h1 className="text-md">
            {balance 
            ? formatCurrency( balance, props.currencyData.currency)
            : "0"
            }
          </h1>
        </div>
      </div>
      <div className="rounded-xl bg-gray-100 mt-2 p-4 flex">
        <img
          className="w-10 h-10 rounded-full object-cover"
          src={props.userData.avatar ? props.userData.avatar : avatar}
          alt="User Avatar"
        />
        <div className="ml-4">
          <h1 className="text-sm text-gray-500">Người nhận:</h1>
          <h1 className="text-md">{props.userData.address}</h1>
        </div>
      </div>
      <button
        onClick={handleSubmit}
        className="w-full mt-10 p-3 bg-blue-600 transition-colors duration-500 ease-in-out hover:bg-blue-500 font-semibold text-white rounded-full"
      >
        Gửi
      </button>
      <DrawerBottom data={dataSend} onClose={handleSubmit} state={isOpen} />
    </div>
  );
};

export default InputAmount;
