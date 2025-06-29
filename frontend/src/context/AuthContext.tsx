// src/context/AuthContext.tsx

import { createContext, useState, useEffect, useContext } from "react";
import type { ReactNode } from "react";
import type { User } from "../types/auth";
import { getToken, setToken, clearToken } from "../utils/tokenStorage";

interface AuthContextType {
  user: User | null;
  token: string | null;
  login: (user: User, token: string) => void;
  logout: () => void;
  isAuthenticated: boolean;
  isLoading: boolean;
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [user, setUser] = useState<User | null>(null);
  const [token, setAuthToken] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const storedToken = getToken();
    const storedUser = localStorage.getItem("user");

    if (storedToken && storedUser) {
      setAuthToken(storedToken);
      setUser(JSON.parse(storedUser));
    }

    setIsLoading(false); 
  }, []);

  const login = (user: User, token: string) => {
    setAuthToken(token);
    setUser(user);
    setToken(token);
    localStorage.setItem("user", JSON.stringify(user));
  };

  const logout = () => {
    setAuthToken(null);
    setUser(null);
    clearToken();
    localStorage.removeItem("user");
  };

  return (
    <AuthContext.Provider
      value={{
        user,
        token,
        login,
        logout,
        isAuthenticated: !!token,
        isLoading,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export const useAuthContext = () => {
  const context = useContext(AuthContext);
  if (!context) throw new Error("useAuthContext must be used within an AuthProvider");
  return context;
};
