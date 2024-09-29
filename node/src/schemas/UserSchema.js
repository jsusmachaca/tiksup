const Joi = require('joi');

const registerUserSchema = Joi.object({
  first_name: Joi.string().required(),
  username: Joi.string().required(),
  email: Joi.string().email().required(),
  password: Joi.string().required()
});

const loginUserSchema = Joi.object({
  username: Joi.string().required(),
  password: Joi.string().required()
});

module.exports = { registerUserSchema, loginUserSchema };
