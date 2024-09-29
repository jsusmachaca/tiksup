const { movieClient } = require('../services/GrpcService');

const getMovies= (req, res) => {
  const authHeader = req.headers['authorization'];
  const token = authHeader && authHeader.startsWith('Bearer ') 
                ? authHeader.split(' ')[1] 
                : null;

  if (!token) {
    return res.status(401).send('Token no proporcionado');
  }

  movieClient.GetMoviesByToken({ token }, (err, response) => {
    if (err) {
      return res.status(500).send('Error en la petición gRPC: ' + err.message);
    }
    res.send(response.movies);
  });
};

module.exports = { getMovies };