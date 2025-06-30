import { Outlet } from "react-router-dom";

export default function RegisteredRoute() {
  // const { security_code } = useAppSelector(
  //   (state) => state.user.userState?.userData
  // );

  // return security_code ? (
  //   <Outlet />
  // ) : (
  //   <Navigate to={"/auth/register/security-code"} />
  // );
  return <Outlet />;
}
