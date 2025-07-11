import { apiSlice } from "../api/apiSlice";
import { IProfile } from "./profile";

export const profileApi = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getProfile: builder.query<IProfile, void>({
      query: () => "/api/v1/users/me",
      keepUnusedDataFor: 20,
    }),
  }),
});

export const { useGetProfileQuery } = profileApi;
