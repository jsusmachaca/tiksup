const { producer } = require('../config/kafka');
const { movieDataSchema } = require('../schemas/MovieDataSchema');

const postUserMovieData = async (req, res) => {
  const { user_id, video_id, waching_time, waching_repeat, preferences, next } = req.body;

  const { error } = movieDataSchema.validate({ user_id, video_id, waching_time, waching_repeat, preferences, next });

  if (error) {
    return res.status(400).send(`Error de validación: ${error.details[0].message}`);
  }

  try {
    const mensajeJson = {
      user_id,
      video_id,
      waching_time,
      waching_repeat,
      preferences,
      next,
    };

    const mensajeString = JSON.stringify(mensajeJson);

    await producer.send({
      topic: 'nodetest', 
      messages: [{ value: mensajeString }],
    });

    res.status(200).send('Mensaje enviado a Kafka con éxito');
  } catch (error) {
    console.error('Error al enviar mensaje a Kafka:', error);
    res.status(500).send('Error al enviar mensaje a Kafka');
  }
};

module.exports = { postUserMovieData };