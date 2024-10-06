"use client";
import { useContext, useEffect, useRef } from 'react';
import AuthContext from '../../context/AuthContext';
import VideoItem from './VideoItem';

const VideoList = () => {
  const authContext = useContext(AuthContext);
  const containerRef = useRef<HTMLDivElement>(null);

  const { getVideos, videos = [] } = authContext || { getVideos: () => {}, videos: [] };

  useEffect(() => {
    if (getVideos) {
      getVideos();
    }
  }, [getVideos]);

  const handleScroll = (direction: 'up' | 'down') => {
    const container = containerRef.current;
    if (container) {
      const scrollAmount = direction === 'up' ? -window.innerHeight : window.innerHeight;
      container.scrollBy({
        top: scrollAmount,
        behavior: 'smooth',
      });
    }
  };
  
  if (!authContext) {
    return <div>Error: AuthContext not found</div>;
  }

  if (videos.length === 0) {
    return <div>Loading videos...</div>;
  }

  return (
    <div className="relative h-screen overflow-hidden">
      <div
        className="video-list h-full overflow-y-scroll snap-y snap-mandatory"
        ref={containerRef}
      >
        {videos.map((video, index) => (
          <VideoItem key={index} video={video} />
        ))}
      </div>
      <div className="absolute top-1/2 left-0 transform -translate-y-1/2">
        <button
          className="bg-gray-800 text-white p-2 rounded-full"
          onClick={() => handleScroll('up')}
        >
          ↑
        </button>
      </div>
      <div className="absolute bottom-1/2 left-0 transform translate-y-1/2">
        <button
          className="bg-gray-800 text-white p-2 rounded-full"
          onClick={() => handleScroll('down')}
        >
          ↓
        </button>
      </div>
    </div>
  );
};

export default VideoList;
