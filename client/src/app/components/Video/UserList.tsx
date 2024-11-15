import React from 'react';
import UserCard from './UserCard';

const UserList = () => {
  const users = [
    { name: 'Usuario1', followers: 1200 },
    { name: 'Usuario2', followers: 3400 },
    { name: 'Usuario3', followers: 2900 },
  ];

  return (
    <div className="p-4 bg-gradient-to-br from-purple-200 to-white min-h-screen text-black">
      <h1 className="text-2xl font-bold mb-4">Usuarios Recomendados</h1>
      <div className="grid grid-cols-1 gap-4">
        {users.map((user, index) => (
          <UserCard key={index} user={user} />
        ))}
      </div>
    </div>
  );
};

export default UserList;
