"use client";

import { useState, useContext, ChangeEvent, FormEvent } from 'react';
import AuthContext from '../../context/AuthContext';
import { useRouter } from 'next/navigation';

const AuthForm = () => {
  const [form, setForm] = useState({ first_name: '', username: '', password: '' });
  const [isLogin, setIsLogin] = useState(true);
  const [error, setError] = useState('');
  const authContext = useContext(AuthContext);
  const router = useRouter();

  if (!authContext) {
    return <div>Error: AuthContext not found</div>;
  }

  const { login, register } = authContext;

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    try {
      if (isLogin) {
        await login(form.username, form.password);
        router.push('/videos');
      } else {
        await register(form.first_name, form.username, form.password);
        router.push('/videos');
      }
    } catch (error) {
      setError('Could not log in or register. Please try again.');
      console.error("Error during login/register", error); // Manejo del error
    }
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen bg-gray-100 text-black">
      <form onSubmit={handleSubmit} className="flex flex-col w-80 p-6 bg-white rounded-lg shadow-md">
        {error && <div className="mb-4 text-red-500">{error}</div>}
        {!isLogin && (
          <input
            type="text"
            name="first_name"
            placeholder="First Name"
            value={form.first_name}
            onChange={handleChange}
            className="mb-4 p-2 border border-gray-300 rounded"
            required
          />
        )}
        <input
          type="text"
          name="username"
          placeholder="Username"
          value={form.username}
          onChange={handleChange}
          className="mb-4 p-2 border border-gray-300 rounded"
          required
        />
        <input
          type="password"
          name="password"
          placeholder="Password"
          value={form.password}
          onChange={handleChange}
          className="mb-4 p-2 border border-gray-300 rounded"
          required
        />
        <button type="submit" className="p-2 bg-blue-500 text-white rounded hover:bg-blue-600">
          {isLogin ? 'Login' : 'Register'}
        </button>
      </form>
      <p className="mt-4">
        {isLogin ? (
          <>
            ¿No tienes una cuenta?{' '}
            <button onClick={() => setIsLogin(false)} className="text-blue-500 hover:underline">
              Regístrate aquí
            </button>
          </>
        ) : (
          <>
            ¿Ya tienes una cuenta?{' '}
            <button onClick={() => setIsLogin(true)} className="text-blue-500 hover:underline">
              Inicia sesión aquí
            </button>
          </>
        )}
      </p>
    </div>
  );
};

export default AuthForm;
