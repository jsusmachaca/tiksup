import React from 'react';

const UserCard = ({ user: { name, followers } }: { user: { name: string; followers: number } }) => {
  return (
    <div className="bg-white shadow-md rounded-lg p-4 flex items-center text-black">
      <div className="w-12 h-12 rounded-full bg-gray-300 mr-4 flex-shrink-0"></div>
      <div>
        <h3 className="font-bold text-lg">{name}</h3>
        <p className="text-sm text-gray-600">{followers} seguidores</p>
        <button className="text-blue-500 text-sm mt-1">Seguir</button>
      </div>
    </div>
  );
};

export default UserCard;
