const Joi = require('joi');

const streamDataSchema = Joi.object({
    user_id: Joi.string().required(),
    video_id: Joi.string().required(),
    watching_time: Joi.number().required(),
    watching_repeat: Joi.number().required(),
    data: Joi.object({
      genre: Joi.array().items(
        Joi.string()
      ).required(),
      protagonist: Joi.string().required(),
      director: Joi.string().required()
    }).required(),
  });

module.exports = {streamDataSchema};