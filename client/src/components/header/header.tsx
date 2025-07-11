
// import { Link } from "react-router-dom";
// import { useGetProfileQuery } from "../../redux/features/profile/profileApi";
// import AuthImg from "../../assets/png/auth_img.jpg";
// export default function Header() {
//   const { data: user } = useGetProfileQuery();
//   return (
//     <div>
//       <div
//         class={`h-20 w-full flex fixed top-0 left-0 right-0 z-50 bg-white p-4 items-center shadow-sm`}
//       >
//         <Link to="/">
//           <img class={`w-40`} />
//         </Link>
//         <div class={`ml-auto`}>
//           <img
//             class={`w-14 h-14 object-cover rounded-full`}
//             src={user?.data.userData.avatar}
//           ></img>
//         </div>
//       </div>
//       <div class={`h-20`}></div>
//     </div>
//   );
// }

import { Link } from "react-router-dom";
import { useGetProfileQuery } from "../../redux/features/profile/profileApi";
import AuthImg from "../../assets/png/auth_img1.jpg";
import { NonceManager } from "ethers";

export default function Header() {
  const { data: user } = useGetProfileQuery();

  return (
    <div>
      <div
        class={`h-20 w-full flex fixed top-0 left-0 right-0 z-50 bg-white p-4 items-center shadow-sm`}
      >
        <Link to="/" class="flex items-center gap-3">
          <img src={AuthImg} alt="Logo" class="h-16 w-16 object-cover rounded-full" />
          <span class="text-4xl font-bold text-gray-800">Vaulta</span>
        </Link>
        <div class={`ml-auto`}>
          <img
            class={`w-14 h-14 object-cover rounded-full`}
            //src={user?.data.userData.avatar} // -change
            alt="User avatar"
          />
        </div>
      </div>
      <div class={`h-20`}></div>
    </div>
  );
}
