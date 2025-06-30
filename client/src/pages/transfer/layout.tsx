import { Outlet, useNavigate, useLocation } from "react-router-dom";
import { useEffect } from "react";

const TransferLayout = () => {
  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    if (location.pathname === "/transfer") {
      navigate("/transfer/select-currency");
    }
  }, [location.pathname, navigate]);

  return (
    <div className="w-full h-full p-4">
      <Outlet />
    </div>
  );
};

export default TransferLayout;