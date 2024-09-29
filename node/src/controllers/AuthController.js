const { authClient } = require('../services/GrpcService');
const { registerUserSchema, loginUserSchema } = require('../schemas/UserSchema');

const registerUser = (req, res) => {
  const { first_name, username, email, password } = req.body;

   const { error } = registerUserSchema.validate({ first_name, username, email, password });

   if (error) {
     return res.status(400).send(`Error de validación: ${error.details[0].message}`);
   }

  authClient.registerUser({ first_name, username, email, password }, (err, response) => {
    if (err) {
      return res.status(500).send('Error en la petición gRPC: ' + err.message);
    }
    res.send(response.token); 
  });
};

const loginUser = (req, res) => {
  const { username, password } = req.body; 
  const { error } = loginUserSchema.validate({ username, password });

  if (error) {
    return res.status(400).send(`Error de validación: ${error.details[0].message}`);
  }
  authClient.loginUser({ username, password }, (err, response) => {
    if (err) {
      return res.status(500).send('Error en la petición gRPC: ' + err.message);
    }
    res.send(response.token); 
  });
};


module.exports = { registerUser, loginUser };