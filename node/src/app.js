const express = require('express');
const cargarRutas = require('./routes/Loader');
const dotenv = require('dotenv');
dotenv.config();

const { iniciarProductor } = require('./config/kafka');

const app = express();
const port = process.env.PORT || 3000;

app.use(express.json());

cargarRutas(app);

app.get('/', (req, res) => {
  res.send('¡Hola, mundo!');
});

app.listen(port, async () => {
  console.log(`Servidor Express corriendo en http://localhost:${port}`);
  await iniciarProductor();
});