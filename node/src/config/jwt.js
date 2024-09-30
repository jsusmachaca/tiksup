const jwt = require("jsonwebtoken")

const validateToken = (token) => {
  try {
    const decodedToken = jwt.verify(token, process.env.SECRET_KEY)
    return decodedToken
  } catch (err) {
    return null
  }
}

module.exports = {
  validateToken
}