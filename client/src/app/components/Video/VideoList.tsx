"use client";

import { useEffect, useRef } from 'react';
import { useVideoContext } from '../../context/VideoContext';
import VideoItem from './VideoItem';
import { ChevronUp, ChevronDown } from 'lucide-react';

const VideoList = () => {
  const { getVideos, videos } = useVideoContext();
  const containerRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    getVideos().catch(error => console.error('Error fetching videos:', error));
  }, []);

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

  if (!videos || videos.length === 0) {
    return <div className="h-screen flex items-center justify-center text-white bg-black">Loading videos...</div>;
  }

  return (
    <div className="relative h-[calc(100vh-64px)] bg-gray-100 overflow-hidden">
      <div
        className="video-list h-full overflow-y-scroll snap-y snap-mandatory"
        ref={containerRef}
      >
        {videos.map((video, index) => (
          <div key={index} className="snap-start h-full flex items-center justify-center">
            <VideoItem video={video} />
          </div>
        ))}
      </div>
      <div className="absolute top-4 right-4 z-10">
        <button
          className="bg-gray-800 text-white p-2 rounded-full mb-2"
          onClick={() => handleScroll('up')}
        >
          <ChevronUp className="w-6 h-6" />
        </button>
      </div>
      <div className="absolute bottom-4 right-4 z-10">
        <button
          className="bg-gray-800 text-white p-2 rounded-full"
          onClick={() => handleScroll('down')}
        >
          <ChevronDown className="w-6 h-6" />
        </button>
      </div>
    </div>
  )
}

export default VideoList;