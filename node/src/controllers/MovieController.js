const { client } = require('../config/redis');
const { validateToken } = require('../config/jwt');

const getMovies= async(req, res) => {
  try{
    const authHeader = req.headers.authorization;
    if (!authHeader || !authHeader.startsWith('Bearer'))
      return res.status(401).send('Token no proporcionado');

    const token = authHeader.substring(7)
    const decodedToken = validateToken(token)
    if (decodedToken === null) return res.status(401).json({ error: 'Token no valido' });

    const recommendations = await client.get(`user:${decodedToken.user_id}:recommendations`)

    res.json(JSON.parse(recommendations))
  }catch(err){
    res.status(500).send('Error: ' + err.message);
  }
};

module.exports = { getMovies };