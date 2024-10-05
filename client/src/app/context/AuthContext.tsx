"use client";

import { createContext, useState, useEffect, ReactNode } from 'react';
import { useRouter } from 'next/navigation';
import api from '../api/apiClient';
import { AxiosError } from 'axios';

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
  const [videosWatched, setVideosWatched] = useState(1);
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

  const getVideos = async (append = false) => {
    try {
      const res = await api.get('/movies');
      setVideos((prevVideos) => append ? [...prevVideos, ...res.data.movies] : res.data.movies);
    } catch (error) {
      console.error('Error fetching videos:', (error as AxiosError).response!);
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
      next: videosWatched >= 4
    };

    try {
  
      const res = await api.post('/stream/sendmoviedata', data, {
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('token')}`,
          'Content-Type': 'application/json',
        },
      });

      setVideosWatched((prev) => prev + 1);
      console.log(`Datos enviados con éxito ${res.data.message}`);

      if (videosWatched >= 4) {
        await getVideos(true);
        setVideosWatched(1);
        console.log('fetching for more movies')
      }
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
