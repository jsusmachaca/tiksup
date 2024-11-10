import React from 'react';
import { Bell, MessageSquare, Search, User } from 'lucide-react';
import Image from 'next/image';

const TopBar = () => (
  <div className="fixed top-0 left-0 right-0 h-16 bg-white shadow-md z-50 flex items-center justify-between px-4 text-black">
    <Image
      src="https://gqldqjebnxaxmslrgqwf.supabase.co/storage/v1/object/sign/Login/mTPwA2Cq_400x400.png?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJMb2dpbi9tVFB3QTJDcV80MDB4NDAwLnBuZyIsImlhdCI6MTcyOTUzMTIwNywiZXhwIjoxNzYxMDY3MjA3fQ.bshfpwwvzm3nj4ih87M0NuNRcPl510R4gHVTjNU3lVU&t=2024-10-21T17%3A20%3A08.182Z"
      alt="App Logo"
      width={40}
      height={40}
      className="rounded-full"
    />
    <div className="flex items-center space-x-4">
      <Search className="w-6 h-6 text-gray-600" />
      <Bell className="w-6 h-6 text-gray-600" />
      <MessageSquare className="w-6 h-6 text-gray-600" />
      <User className="w-6 h-6 text-gray-600" />
    </div>
  </div>
)

export default TopBar;