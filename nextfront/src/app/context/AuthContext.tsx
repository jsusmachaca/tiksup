"use client";

import { createContext, useState, useEffect, ReactNode } from 'react';
import { useRouter } from 'next/navigation';
import api from '../api/api';

interface AuthContextProps {
  user: { username: string } | null;
  login: (username: string, password: string) => Promise<void>;
  register: (first_name: string, username: string, password: string) => Promise<void>;
  logout: () => void;
  getVideos: () => Promise<void>;
  sendVideoData: (video: any, watchingTime: number, watchingRepeat: number) => Promise<void>;
  videos: any[];
}

const AuthContext = createContext<AuthContextProps | undefined>(undefined);

interface AuthProviderProps {
  children: ReactNode;
}

export const AuthProvider = ({ children }: AuthProviderProps) => {
  const [user, setUser] = useState<{ username: string } | null>(null);
  const [videos, setVideos] = useState<any[]>([]);
  const router = useRouter();

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      setUser({ username: 'user1245' });
    }
  }, []);

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
    } catch (error) {
      console.error('Error during login:', error);
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
    } catch (error) {
      console.error('Error during registration:', error);
    }
  };

  const logout = () => {
    localStorage.removeItem('token');
    setUser(null);
    router.push('/login');
  };

  const getVideos = async () => {
    try {
      const res = await api.get('/movies');
      setVideos(res.data.movies);
    } catch (error) {
      console.error('Error fetching videos:', error);
    }
  };

  const sendVideoData = async (video:   any, watchingTime: number, watchingRepeat: number) => {
    console.log(video.watching_repeat);
    const roundedWatchingTime = parseFloat(watchingTime.toFixed(2));
    const data = {
      movie_id: video.id,
      watching_time: roundedWatchingTime,
      watching_repeat: watchingRepeat,
      data: {
        genre: video.genre,
        protagonist: video.protagonist,
        director: video.director,
      },
      next: false
    };
  
    try {
  
      const res = await api.post('/stream/sendMovieData', data, {
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('token')}`,
          'Content-Type': 'application/json',
        },
      });
  
      console.log('Datos enviados con éxito');
    } catch (error) {
      console.error('Error sending video data:');
    }
  };
  


  return (
    <AuthContext.Provider value={{ user, login, register, logout, getVideos, sendVideoData, videos }}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthContext;
