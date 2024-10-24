"use client";

import VideoList from "../components/Video/VideoList";
import TopBar from "../components/Video/Topbar";  
import UserList from "../components/Video/UserList";

export default function VideosPage() {
  return (
    <div className="bg-gray-100 min-h-screen">
      <TopBar />
      <div className="pt-16 flex">
        <UserList />
        <div className="flex-grow ml-64">
          <VideoList />
        </div>
      </div>
    </div>
  ) 
};