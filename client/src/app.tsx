import React from "react";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Outlet,
} from "react-router-dom";
import Home from "./pages/home";
import History from "./pages/transaction/history";
import TransactionDetails from "./pages/transaction/details";
import ReceivePage from "./pages/receive/receive";
import Login from "./pages/authentication/login";
import PageNotFound from "./pages/page_not_found";
import VerifyLogin from "./pages/authentication/verify_login";
import ScanQR from "./pages/payment/scanqr";
import { PaymentGateway } from "./pages/payment/payment";
import Transfer from "./pages/transfer/transfer";
import Register from "./pages/authentication/register";
import VerifyRegister from "./pages/authentication/verify_register";
import SecurityCode from "./pages/authentication/setup_security_code";
import TransferResults from "./pages/transfer/transfer_results";
import ProtectRoutes from "./utils/protect_routes";
import DepositWithdraw from "./pages/deposit-withdraw";
import PaymentResults from "./pages/payment/payment_results";
import AddCreditCard from "./pages/credit-card/pages/add-credit-card";
import CreditCard from "./pages/credit-card";
import Setting from "./pages/setting";
import SideBar from "./components/sidebar/sidebar";
import Header from "./components/header/header";
import TransferByQrCode from "./pages/receive/transfer";
import Result from "./pages/deposit-withdraw/result";
import SelectOptions from "./pages/payment/select_options";
import RegisteredRoute from "./utils/registered-route";
import B2BAPIManagement  from "./pages/connect-app/connect-app";

// B2B Components
import B2BDashboard from "./pages/b2b/dashboard";
import B2BUsers from "./pages/b2b/users";
import B2BSidebar from "./components/b2b/sidebar";
import B2BHeader from "./components/b2b/header";

import ForgotPasswordForm from "./pages/authentication/forgot-password";
import ResetPasswordForm from "./pages/authentication/reset-password";
import TransferLayout from "./pages/transfer/layout";

const AuthenticatedLayout: React.FC = () => {
  return (
    <>
      <Header />
      <div className="font-inter bg-gray-50 h-screen flex w-full">
        <SideBar state="Trang chủ" />
        <Outlet />
      </div>
    </>
  );
};

const B2BLayout: React.FC = () => {
  return (
    <>
      <B2BHeader />
      <div className="font-inter bg-gray-50 min-h-screen flex w-full">
        <B2BSidebar />
        <div className="flex-1 ml-64">
          <Outlet />
        </div>
      </div>
    </>
  );
};

const NonAuthenticatedLayout: React.FC = () => {
  return (
    <div className="font-inter flex-1">
      <Outlet />
    </div>
  );
};

export default function App() {
  return (
    <Router>
      <Routes>
         <Route element={<ProtectRoutes />}>
          <Route element={<AuthenticatedLayout />}>
            <Route path="/" element={<Home />} />
            <Route path="/setting" element={<Setting />} />
            <Route path="/connect-app" element={<B2BAPIManagement />} />
            {/* <Route path="/connect-app-list" element={<ConnectAppList />} /> */}
            <Route path="/scan-qrcode" element={<ScanQR />} />
            <Route path="/payment" element={<PaymentGateway />} />
            <Route path="/payment/results" element={<PaymentResults />} />
            <Route path="/option-payment" element={<SelectOptions />} />
            <Route path="/transfer" element={<Transfer />} />
            <Route path="/transfer/result" element={<TransferResults />} />
            <Route path="/transfer/info" element={<TransferByQrCode />} />
            <Route path="/transfer" element={<TransferLayout />}>
              <Route index element={<Transfer />} />
              <Route path="result" element={<TransferResults />} />
              <Route path="info" element={<TransferByQrCode />} />
            </Route>
            
            <Route path="/transaction/history" element={<History />} />
            <Route
              path="/transaction/details"
              element={<TransactionDetails />}
            />
            <Route path="/deposit-withdraw" element={<DepositWithdraw />} />
            <Route path="/deposit-withdraw/result" element={<Result />} />
            <Route path="/credit-card" element={<CreditCard />}>
              <Route path="add-card" element={<AddCreditCard />} />
            </Route>
            <Route path="/receive-page" element={<ReceivePage />} />
          </Route>

          {/* B2B Routes */}
          <Route element={<B2BLayout />}>
            <Route path="/b2b" element={<B2BDashboard />} />
            <Route path="/b2b/dashboard" element={<B2BDashboard />} />
            <Route path="/b2b/users" element={<B2BUsers />} />
            <Route path="/b2b/api-keys" element={<B2BAPIManagement />} />
            <Route path="/b2b/analytics" element={<div className="p-6">Analytics Page (Coming Soon)</div>} />
            <Route path="/b2b/settings" element={<div className="p-6">B2B Settings (Coming Soon)</div>} />
          </Route>

          <Route path="*" element={<PageNotFound />} />
         </Route> 

        <Route element={<RegisteredRoute />}>
          <Route element={<NonAuthenticatedLayout />}>
            <Route path="/auth/login" element={<Login />} />
            <Route path="/auth/login/verify-login" element={<VerifyLogin />} />
            <Route path="/auth/register" element={<Register />} />
            <Route
              path="/auth/register/verify-register"
              element={<VerifyRegister />}
            />
            <Route
              path="/auth/register/security-code"
              element={<SecurityCode />}
            />
            <Route
              path="/auth/login/forgot-password"
              element={<ForgotPasswordForm />}
            />
            <Route
              path="/auth/login/reset-password"
              element={<ResetPasswordForm />}
            />
          </Route>
        </Route>

        <Route path="*" element={<PageNotFound />} />
      </Routes>
    </Router>
  );
}
