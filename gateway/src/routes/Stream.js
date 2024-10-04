import { Router } from 'express'
import { postUserMovieData } from '../controllers/StreamController.js'

const router = Router()

router.post('/sendMovieData', postUserMovieData)

export default router
