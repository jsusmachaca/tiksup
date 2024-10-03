"use client";

import { useContext, useEffect, useRef, useState } from 'react';
import { useInView } from 'react-intersection-observer';
import AuthContext from '../../context/AuthContext';

interface VideoItemProps {
  video: any;
}

const VideoItem = ({ video }: VideoItemProps) => {
  const videoRef = useRef<HTMLVideoElement>(null);
  const { ref, inView } = useInView({
    threshold: 0.5,
  });

  const authContext = useContext(AuthContext);
  const [watchingTime, setWatchingTime] = useState(0);
  const [watchingRepeat, setWatchingRepeat] = useState(0);
  const [intervalId, setIntervalId] = useState<NodeJS.Timeout | null>(null);

  useEffect(() => {
    const videoElement = videoRef.current;
    if (videoElement) {
      if (inView) {
        videoElement.play();
        setWatchingRepeat(prev => prev + 1);

        const id = setInterval(() => {
          setWatchingTime(prev => prev + 1);
        }, 1000);
        setIntervalId(id);
      } else {
        videoElement.pause();
        if (intervalId) {
          clearInterval(intervalId);
          setIntervalId(null);
        }

        if (authContext) {
          authContext.sendVideoData(video, watchingTime, watchingRepeat);
        }
      }
    }
  }, [inView, authContext, video, watchingTime, watchingRepeat, intervalId]);

  if (!video) {
    return null; // O muestra un mensaje de error
  }

  return (
    <div className="video-item snap-start" ref={ref}>
      <h3 className="text-lg font-bold">{video.title}</h3>
      <video ref={videoRef} controls muted className="w-full h-auto">
        <source src={video.url} type="video/mp4" />
        Your browser does not support the video tag.
      </video>
      <p>Protagonist: {video.protagonist}</p>
      <p>Director: {video.director}</p>
      <p>Genres: {video.genre.join(', ')}</p>
    </div>
  );
};

export default VideoItem;
