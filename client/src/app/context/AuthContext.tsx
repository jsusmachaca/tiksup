"use client";

import { createContext, useState, ReactNode, useContext } from 'react';
import { useRouter } from 'next/navigation';
import api from '../api/apiClient';
import { AxiosError } from 'axios';

interface AuthContextProps {
  user: { username: string } | null;
  login: (username: string, password: string) => Promise<void>;
  register: (first_name: string, username: string, password: string) => Promise<void>;
  logout: () => void;
}

const AuthContext = createContext<AuthContextProps | undefined>(undefined);

interface AuthProviderProps {
  children: ReactNode;
}

export const AuthProvider = ({ children }: AuthProviderProps) => {
  const [user, setUser] = useState<{ username: string } | null>(null);
  const router = useRouter();

  const login = async (username: string, password: string) => {
    try {
      const res = await api.post('/auth/login', {
        username,
        password,
      });
      const data = res.data;
      if (data.access_token) {
        localStorage.setItem('token', data.access_token);
        setUser({ username });
        router.push('/videos');
      }
    } catch (error: unknown) {
      if (error instanceof AxiosError) {
        console.error('Error during login:', error.response?.data || error.message);
      } else {
        console.error('Unexpected error during login:', error);
      }
    }
  };

  const register = async (first_name: string, username: string, password: string) => {
    try {
      const res = await api.post('/auth/register', {
        first_name,
        username,
        password,
      });
      const data = res.data;
      if (data.success) {
        await login(username, password);
      }
    } catch (error: unknown) {
      if (error instanceof AxiosError) {
        console.error('Error during registration:', error.response?.data || error.message);
      } else {
        console.error('Unexpected error during registration:', error);
      }
    }
  };

  const logout = () => {
    localStorage.removeItem('token');
    setUser(null);
    router.push('/login');
  };

  return (
    <AuthContext.Provider value={{ user, login, register, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuthContext = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuthContext must be used within an AuthProvider');
  }
  return context;
};

export default AuthContext;
