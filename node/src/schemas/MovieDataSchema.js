const Joi = require('joi');

const movieDataSchema = Joi.object({
    user_id: Joi.number().required(),
    video_id: Joi.number().required(),
    waching_time: Joi.string().pattern(/^\d+s$/).required(),
    waching_repeat: Joi.number().required(),
    preferences: Joi.object({
      genre_scores: Joi.array().items(
        Joi.object().pattern(Joi.string(), Joi.number().required())
      ).required(),
      actors_scores: Joi.array().items(
        Joi.object().pattern(Joi.string(), Joi.number().required())
      ).required(),
      director_scores: Joi.array().items(
        Joi.object().pattern(Joi.string(), Joi.number().required())
      ).required()
    }).required(),
    next: Joi.boolean().required()
  });

module.exports = {movieDataSchema};
