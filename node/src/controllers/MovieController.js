const { movieClient } = require('../services/GrpcService');
const axios = require('axios');

const getMovies= async(req, res) => {
  const authHeader = req.headers['authorization'];
  const token = authHeader && authHeader.startsWith('Bearer ') 
                ? authHeader.split(' ')[1] 
                : null;

  if (!token) {
    return res.status(401).send('Token no proporcionado');
  }

  try{
    const endpointURL = `${process.env.WORKER_URL}/movies`;

    const request = {
        token : token
    };

    const response = await axios.post(endpointURL, request);
    
    res.send(response.data);

  }catch(err){
    res.status(500).send('Error: ' + err.message);
  }
};

module.exports = { getMovies };