import React from "react";
import { formatCurrency } from "../../utils/format_currency";
import { ItemCurrencyLoading } from "../../components/home/item_currency";
import HeaderDefault from "../../components/header/header_default";
import VNDIcon from "../../assets/png/vnd_icon.jpg";
import USDIcon from "../../assets/png/usd_icon.png";
import ETHIcon from "../../assets/png/eth_icon.png";



type Props = {
  currency: any;
  handleStepTransfer: (key: string) => void;
  handleCurrencyData: (data: any) => void;
  isLoading: boolean;
};
const SelectCurrency: React.FC<Props> = ({ ...props }) => {
  const handleSelectCurrency = () => {
    props.handleStepTransfer("search_user");
  };
  const handleCurrencyData = (data: any) => {
    props.handleCurrencyData(data);
  };


  return (
    <div class={`container-center`}>
      <HeaderDefault title="Chọn loại tiền"></HeaderDefault>
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
            wallet={props.currency[0]}
            currency={`VND`}
            onClick={handleSelectCurrency}
            handleCurrencyData={handleCurrencyData}
          ></ItemCurrency>
          <ItemCurrency
            image={USDIcon}
            wallet={props.currency[1]}
            currency={`USD`}
            onClick={handleSelectCurrency}
            handleCurrencyData={handleCurrencyData}
          ></ItemCurrency>
          <ItemCurrency
            image={ETHIcon}
            wallet={props.currency[2]}
            currency={`ETH`}
            onClick={handleSelectCurrency}
            handleCurrencyData={handleCurrencyData}
          ></ItemCurrency>
        </>
      )}
    </div>
  );
};
const ItemCurrency = ({ ...props }) => {
  return (
    <div
      onClick={() => {
        props.onClick();
        props.handleCurrencyData({
          ...props.wallet,
          currency: props.currency,
          image: props.image,
        });
      }}
      className={`flex justify-between items-center my-4 hover:bg-gray-100 p-2 rounded-lg cursor-pointer`}
    >
      <div className={`flex items-center`}>
        <img src={props.image} className={`w-12 h-12`}></img>
        <div className={`ml-2`}>
          <h1 className={`font-semibold text-xl`}>{props.currency}</h1>
        </div>
      </div>
      <div>
        <h1 className={`w-full font-semibold text-xl text-end`}>
          {/* {formatCurrency(props.wallet.balance, "VND")} */}
        </h1>
      </div>
    </div>
  );
};
export default SelectCurrency;
