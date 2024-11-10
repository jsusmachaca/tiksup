import jwt from 'jsonwebtoken'
import 'dotenv/config'

export const generateToken = (data) => jwt.sign(data, process.env.SECRET_KEY, { expiresIn: '1y' })

export const validateToken = (token) => {
  try {
    const decodedToken = jwt.verify(token, process.env.SECRET_KEY)
    return decodedToken
  } catch (err) {
    return null
  }
}
