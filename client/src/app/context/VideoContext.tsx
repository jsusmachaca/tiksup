"use client";

import { createContext, useState, ReactNode, useContext } from 'react';
import api from '../api/apiClient';  

type VideoType = {
  id: string;
  url: string
  title: string
  protagonist: string
  director: string
  genre: string[]
  watching_repeat?: number
}


interface VideoContextProps {
  videos: VideoType[];
  getVideos: (append?: boolean) => Promise<void>;
  sendVideoData: (video: VideoType, watchingTime: number, watchingRepeat: number) => Promise<void>;
}


const VideoContext = createContext<VideoContextProps | undefined>(undefined);

interface VideoProviderProps {
  children: ReactNode;
}

export const VideoProvider = ({ children }: VideoProviderProps) => {
  const [videos, setVideos] = useState<VideoType[]>([]);
  const [videosWatched, setVideosWatched] = useState(1);

  const getVideos = async (append = false) => {
    try {
      const res = await api.get('/movies');
      setVideos((prevVideos) => append ? [...prevVideos, ...res.data.movies] : res.data.movies);
    } catch (error) {
      console.error('Error fetching videos:', error);
    }
  };

  const sendVideoData = async (video: VideoType, watchingTime: number, watchingRepeat: number) => {
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
        console.log('fetching for more movies');
      }
    } catch (error) {
      console.error('Error sending video data:', error);
    }
  };

  return (
    <VideoContext.Provider value={{ videos, getVideos, sendVideoData }}>
      {children}
    </VideoContext.Provider>
  );
};

export const useVideoContext = () => {
  const context = useContext(VideoContext);
  if (!context) {
    throw new Error('useVideoContext must be used within a VideoProvider');
  }
  return context;
};
