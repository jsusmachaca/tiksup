"use client";

import { useContext, useEffect } from 'react';
import AuthContext from '../../context/AuthContext';
import VideoItem from './VideoItem';

const VideoList = () => {
  const authContext = useContext(AuthContext);

  if (!authContext) {
    return <div>Error: AuthContext not found</div>;
  }

  const { getVideos, videos } = authContext;

  useEffect(() => {
    getVideos();
  }, []);

  if (videos.length === 0) {
    return <div>Loading videos...</div>;
  }

  return (
    <div className="video-list h-screen overflow-y-scroll snap-y snap-mandatory">
      {videos.map((video) => (
        <VideoItem key={video.id} video={video} />
      ))}
    </div>
  );
};

export default VideoList;
