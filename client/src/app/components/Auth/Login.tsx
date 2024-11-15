'use client'

import { useState, useContext, ChangeEvent, FormEvent } from 'react'
import { motion } from 'framer-motion'
import Image from 'next/image'
import AuthContext from '../../context/AuthContext'
import { useRouter } from 'next/navigation'

export default function AuthForm() {
  const [form, setForm] = useState({ first_name: '', username: '', password: '' })
  const [isLogin, setIsLogin] = useState(true)
  const [error, setError] = useState('')
  const authContext = useContext(AuthContext)
  const router = useRouter()

  if (!authContext) {
    return <div>Error: AuthContext not found</div>
  }

  const { login, register } = authContext

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    try {
      if (isLogin) {
        await login(form.username, form.password)
        router.push('/videos')
      } else {
        await register(form.first_name, form.username, form.password)
        router.push('/videos')
      }
    } catch (error) {
      setError('Could not log in or register. Please try again.')
      console.error("Error during login/register", error)
    }
  }

  return (
    <div className="flex items-center justify-center min-h-screen bg-gradient-to-br from-blue-200 to-white">
      <motion.div
        initial={{ opacity: 0, y: -50 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
        className="w-full max-w-md"
      >
        <div className="text-center mb-8">
          <Image
            src="https://gqldqjebnxaxmslrgqwf.supabase.co/storage/v1/object/sign/Login/mTPwA2Cq_400x400.png?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJMb2dpbi9tVFB3QTJDcV80MDB4NDAwLnBuZyIsImlhdCI6MTcyOTUzMTIwNywiZXhwIjoxNzYxMDY3MjA3fQ.bshfpwwvzm3nj4ih87M0NuNRcPl510R4gHVTjNU3lVU&t=2024-10-21T17%3A20%3A08.182Z"
            alt="App Logo"
            width={100}
            height={100}
            className="mx-auto rounded-full shadow-lg"
          />
        </div>
        <motion.form
          onSubmit={handleSubmit}
          className="bg-white shadow-2xl rounded-lg px-8 pt-6 pb-8 mb-4"
          initial={{ scale: 0.9 }}
          animate={{ scale: 1 }}
          transition={{ duration: 0.3 }}
        >
          <h2 className="text-3xl font-bold text-center mb-6 text-blue-600">
            {isLogin ? 'Welcome To Tiksup' : 'Sign Up'}
          </h2>
          {error && (
            <motion.div
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              className="mb-4 text-red-500 text-center"
            >
              {error}
            </motion.div>
          )}
          {!isLogin && (
            <motion.div
              initial={{ height: 0, opacity: 0 }}
              animate={{ height: 'auto', opacity: 1 }}
              exit={{ height: 0, opacity: 0 }}
              transition={{ duration: 0.3 }}
            >
              <input
                type="text"
                name="first_name"
                placeholder="First Name"
                value={form.first_name}
                onChange={handleChange}
                className="w-full px-3 py-2 mb-4 text-gray-700 border border-blue-300 rounded-lg focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition duration-300"
                required
              />
            </motion.div>
          )}
          <input
            type="text"
            name="username"
            placeholder="Username"
            value={form.username}
            onChange={handleChange}
            className="w-full px-3 py-2 mb-4 text-gray-700 border border-blue-300 rounded-lg focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition duration-300"
            required
          />
          <input
            type="password"
            name="password"
            placeholder="Password"
            value={form.password}
            onChange={handleChange}
            className="w-full px-3 py-2 mb-6 text-gray-700 border border-blue-300 rounded-lg focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition duration-300"
            required
          />
          <motion.button
            type="submit"
            className="w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg focus:outline-none focus:shadow-outline transition duration-300"
            whileHover={{ scale: 1.05 }}
            whileTap={{ scale: 0.95 }}
          >
            {isLogin ? 'Login' : 'Register'}
          </motion.button>
        </motion.form>
        <motion.p
          className="text-center text-blue-800"
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ delay: 0.5 }}
        >
          {isLogin ? (
            <>
              ¿No tienes una cuenta?{' '}
              <button onClick={() => setIsLogin(false)} className="font-bold hover:underline">
                Regístrate aquí
              </button>
            </>
          ) : (
            <>
              ¿Ya tienes una cuenta?{' '}
              <button onClick={() => setIsLogin(true)} className="font-bold hover:underline">
                Inicia sesión aquí
              </button>
            </>
          )}
        </motion.p>
      </motion.div>
    </div>
  )
}