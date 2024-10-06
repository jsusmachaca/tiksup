import { client } from '../config/redis.js'
import { validateToken } from '../config/jwt.js'
import axios from 'axios'
import 'dotenv/config'

export const getMovies = async (req, res) => {
  try {
    const authHeader = req.headers.authorization
    if (!authHeader || !authHeader.startsWith('Bearer'))
      return res.status(401).send('Token not provided')

    const token = authHeader.substring(7)
    const decodedToken = validateToken(token)
    if (decodedToken === null) return res.status(401).json({ error: 'Invalid token' })

    const endpointURL = `${process.env.WORKER_URL}/user-info`

    const response = await axios.get(endpointURL, {
      headers: {
        Authorization: authHeader
      }
    })

    const recommendationsString = await client.get(`user:${decodedToken.user_id}:recommendations`)
    const recommendations = JSON.parse(recommendationsString)
    if (
      response.data.preferences.genre_score.length === 0 ||
      response.data.preferences.protagonist_score.length === 0 ||
      response.data.preferences.director_score.length === 0 ||
      !recommendations || recommendations.movies.length === 0
    ) {
      console.log('Using random data')

      const endpointURL = `${process.env.WORKER_URL}/random-movies`

      const randomMovies = await axios.get(endpointURL, {
        headers: {
          Authorization: authHeader
        }
      })

      return res.json(randomMovies.data)
    }

    res.json(recommendations)
  } catch (err) {
    res.status(500).send('Error: ' + err.message)
  }
}
