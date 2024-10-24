"use client";

import { useRef, useEffect, useState, useCallback } from 'react';
import { useInView } from 'react-intersection-observer';
import { useVideoContext } from '../../context/VideoContext';
import { Heart, MessageCircle, Share2, Play } from 'lucide-react';

interface VideoItemProps {
  video: {
    id: string;
    url: string;
    title: string;
    protagonist: string;
    director: string;
    genre: string[];
  };
}

const VideoItem = ({ video }: VideoItemProps) => {
  const videoRef = useRef<HTMLVideoElement>(null);
  const { ref, inView } = useInView({
    threshold: 0.5,
  });

  const { sendVideoData } = useVideoContext();
  const [watchingTime, setWatchingTime] = useState(0);
  const [watchingRepeat, setWatchingRepeat] = useState(0);
  const [hasDataBeenSent, setHasDataBeenSent] = useState(false);
  const [isPlaying, setIsPlaying] = useState(false);
  const playPromiseRef = useRef<Promise<void> | null>(null);

  const handlePlay = useCallback(() => {
    const videoElement = videoRef.current;
    if (videoElement) {
      playPromiseRef.current = videoElement.play();
      playPromiseRef.current
        .then(() => {
          setIsPlaying(true);
        })
        .catch((error) => {
          if (error.name !== 'AbortError') {
            console.error('Error playing video:', error);
          }
        });
    }
  }, []);

  const handlePause = useCallback(() => {
    const videoElement = videoRef.current;
    if (videoElement) {
      if (playPromiseRef.current) {
        playPromiseRef.current
          .then(() => {
            videoElement.pause();
            setIsPlaying(false);
          })
          .catch((error) => {
            if (error.name !== 'AbortError') {
              console.error('Error pausing video:', error);
            }
          });
      } else {
        videoElement.pause();
        setIsPlaying(false);
      }
    }
  }, []);


  useEffect(() => {
    const videoElement = videoRef.current;

    const handleTimeUpdate = () => {
      if (videoElement) {
        setWatchingTime(videoElement.currentTime);
      }
    };

    const handleEnded = () => {
      setWatchingRepeat((prev) => prev + 1);
    };

    if (videoElement) {
      videoElement.addEventListener('timeupdate', handleTimeUpdate);
      videoElement.addEventListener('ended', handleEnded);
    }

    return () => {
      if (videoElement) {
        videoElement.removeEventListener('timeupdate', handleTimeUpdate);
        videoElement.removeEventListener('ended', handleEnded);
      }
    };
  }, []);

  useEffect(() => {
    if (!inView && !hasDataBeenSent && watchingTime > 0) {
      console.log("Sending video data...", { video, watchingTime, watchingRepeat });
      sendVideoData(video, watchingTime, watchingRepeat);
      setHasDataBeenSent(true);
    }
  }, [inView, hasDataBeenSent, sendVideoData, video, watchingTime, watchingRepeat]);

  useEffect(() => {
    const videoElement = videoRef.current;
    if (videoElement) {
      if (inView) {
        videoElement.currentTime = 0; 
        handlePlay();
      } else {
        handlePause();
      }
    }
  }, [inView, handlePlay, handlePause]);

  const togglePlay = () => {
    if (isPlaying) {
      handlePause();
    } else {
      handlePlay();
    }
  };
  return (
    <div className="video-item snap-start h-full w-full flex items-center justify-center bg-gray-100" ref={ref}>
      <div className="relative w-full max-w-[400px] h-[calc(100vh-120px)] bg-white rounded-lg shadow-lg overflow-hidden">
        <video 
          ref={videoRef}
          className="w-full h-full object-cover"
          loop
          playsInline
          onClick={togglePlay}
          aria-label={`Video: ${video.title} by ${video.protagonist}`}
        >
          <source src={video.url} type="video/mp4" />
          Your browser does not support the video tag.
        </video>
        <div className="absolute inset-0 bg-black bg-opacity-30 flex items-center justify-center">
          {!isPlaying && (
            <button 
              onClick={togglePlay}
              className="text-white bg-opacity-50 bg-black rounded-full p-4 transition-transform transform hover:scale-110"
            >
              <Play className="w-12 h-12" />
            </button>
          )}
        </div>
        <div className="absolute right-4 bottom-20 flex flex-col items-center space-y-4">
          <button className="p-3 bg-white rounded-full shadow-md transition-transform transform hover:scale-110" aria-label="Like">
            <Heart className="w-6 h-6 text-red-500" />
          </button>
          <button className="p-3 bg-white rounded-full shadow-md transition-transform transform hover:scale-110" aria-label="Comment">
            <MessageCircle className="w-6 h-6 text-blue-500" />
          </button>
          <button className="p-3 bg-white rounded-full shadow-md transition-transform transform hover:scale-110" aria-label="Share">
            <Share2 className="w-6 h-6 text-green-500" />
          </button>
        </div>
        <div className="absolute left-4 right-4 bottom-4 bg-white bg-opacity-75 p-4 rounded-lg">
          <h3 className="text-lg font-bold text-gray-800">{video.title}</h3>
          <p className="text-sm text-gray-600">{video.protagonist}</p>
          <p className="text-xs text-gray-500 mt-1">Director: {video.director}</p>
          <div className="flex flex-wrap mt-2">
            {video.genre.map((genre, index) => (
              <span key={index} className="text-xs bg-blue-100 text-blue-800 rounded-full px-2 py-1 mr-1 mb-1">
                {genre}
              </span>
            ))}
          </div>
        </div>
      </div>
    </div>
  )
}
export default VideoItem;