import { useRef, useEffect, useState, useContext } from 'react';
import { useInView } from 'react-intersection-observer';
import AuthContext from '../../context/AuthContext';

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

  const authContext = useContext(AuthContext);
  const [watchingTime, setWatchingTime] = useState(0);
  const [watchingRepeat, setWatchingRepeat] = useState(0);
  const [hasSentData, setHasSentData] = useState(false);

  useEffect(() => {
    const handleTimeUpdate = () => {
      if (videoRef.current) {
        setWatchingTime(videoRef.current.currentTime);
      }
    };

    const handleEnded = () => {
      setWatchingRepeat((prev) => prev + 1);
    };

    if (videoRef.current) {
      videoRef.current.addEventListener('timeupdate', handleTimeUpdate);
      videoRef.current.addEventListener('ended', handleEnded);
    }

    return () => {
      if (videoRef.current) {
        videoRef.current.removeEventListener('timeupdate', handleTimeUpdate);
        videoRef.current.removeEventListener('ended', handleEnded);
      }
    };
  }, []);

  useEffect(() => {
    if (!inView && !hasSentData && watchingTime > 0) {
      authContext?.sendVideoData(video, watchingTime, watchingRepeat);
      setHasSentData(true);
    }
  }, [inView, hasSentData, authContext, video, watchingTime, watchingRepeat]);

  return (
    <div className="video-item snap-start" ref={ref}>
      <video ref={videoRef} controls muted className="w-full h-full object-cover">
        <source src={video.url} type="video/mp4" />
        Your browser does not support the video tag.
      </video>
      <div className="absolute bottom-0 left-0 p-4">
        <h3 className="text-lg font-bold">{video.title}</h3>
        <p>Protagonist: {video.protagonist}</p>
        <p>Director: {video.director}</p>
        <p>Genres: {video.genre.join(', ')}</p>
      </div>
    </div>
  );
};

export default VideoItem;
