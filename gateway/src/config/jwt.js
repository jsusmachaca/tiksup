import jwt from 'jsonwebtoken'
import 'dotenv/config'

export const validateToken = (token) => {
  try {
    const decodedToken = jwt.verify(token, process.env.SECRET_KEY)
    return decodedToken
  } catch (err) {
    return null
  }
}
