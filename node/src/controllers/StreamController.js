const { preferences } = require('joi');
const { producer } = require('../config/kafka');
const { streamDataSchema } = require('../schemas/StreamDataSchema');

const postUserMovieData = async (req, res) => {
  const { user_id, video_id, watching_time, watching_repeat, data } = req.body;

  const { error } = streamDataSchema.validate({ user_id, video_id, watching_time, watching_repeat, data });
  
  if (error) {
    return res.status(400).send(`Error de validación: ${error.details[0].message}`);
  }

  let preferences = {
    genre_score: [],
    protagonist_score: { name: "", score: 0.0 },
    director_score: { name: "", score: 0.0 }
  };

  data.genre.forEach(genre => {
    preferences.genre_score.push({ name: genre, score: 0.0 });
  });

  if (watching_time >= 15) {
    preferences.genre_score.forEach(item => item.score += 1.0);
    preferences.protagonist_score.score += 1.0;
    preferences.director_score.score += 1.0;
  } else if (watching_time >= 10) {
    preferences.genre_score.forEach(item => item.score += 0.5);
    preferences.protagonist_score.score += 0.5;
    preferences.director_score.score += 0.5;
  } else if (watching_time < 5) {
    preferences.genre_score.forEach(item => item.score -= 0.5);
    preferences.protagonist_score.score -= 0.5;
    preferences.director_score.score -= 0.5;
  }

  if (watching_repeat > 1) {
    const repeatBonus = 0.5 * (watching_repeat - 1);
    preferences.genre_score.forEach(item => item.score += repeatBonus);
    preferences.protagonist_score.score += repeatBonus;
    preferences.director_score.score += repeatBonus;
  }

  preferences.protagonist_score.name = data.protagonist;
  preferences.director_score.name = data.director;

  try {
    const mensajeJson = {
      user_id,
      video_id,
      watching_time,
      watching_repeat,
      preferences,
      next: false,
    };

    const mensajeString = JSON.stringify(mensajeJson);

    await producer.send({
      topic: 'tiksup-user-data', 
      messages: [{ value: mensajeString }],
    });

    res.status(200).json({ message: 'Mensaje enviado a Kafka con éxito' });
  } catch (error) {
    console.error('Error al enviar mensaje a Kafka:', error);
    res.status(500).json({ error: 'Error al enviar mensaje a Kafka' });
  }
};

module.exports = { postUserMovieData };