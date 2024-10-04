import Joi from 'joi'

export const registerUserSchema = Joi.object({
  first_name: Joi.string().required(),
  username: Joi.string().required(),
  password: Joi.string().required()
})

export const loginUserSchema = Joi.object({
  username: Joi.string().required(),
  password: Joi.string().required()
})
