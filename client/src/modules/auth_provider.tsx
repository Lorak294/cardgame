"use client";

import React, { createContext, useState, useEffect } from "react";
import { useRouter } from "next/navigation";

export type UserInfo = {
  username: string;
  id: string;
};

export const AuthContext = createContext<{
  authenticated: boolean;
  setAuthenticated: (auth: boolean) => void;
  user: UserInfo;
  setUser: (user: UserInfo) => void;
}>({
  authenticated: false,
  setAuthenticated: () => {},
  user: { username: "", id: "" },
  setUser: () => {},
});

const AuthContextProvider = ({ children }: { children: React.ReactNode }) => {
  const [authenticated, setAuthenticated] = useState(false);
  const [user, setUser] = useState<UserInfo>({ username: "", id: "" });

  const router = useRouter();

  useEffect(() => {
    const userInfo = localStorage.getItem("user_info");

    if (!userInfo) {
      if (window.location.pathname != "/signup") {
        router.push("/login");
        return;
      }
    } else {
      const stored_user: UserInfo = JSON.parse(userInfo);
      if (stored_user) {
        setUser({
          username: stored_user.username,
          id: stored_user.id,
        });
      }
      setAuthenticated(true);
    }
  }, [authenticated]);

  return (
    <AuthContext.Provider
      value={{
        authenticated: authenticated,
        setAuthenticated: setAuthenticated,
        user: user,
        setUser: setUser,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export default AuthContextProvider;
