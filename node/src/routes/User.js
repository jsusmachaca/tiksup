const express = require('express');
const { postUserMovieData } = require('../controllers/UserController');

const router = express.Router();

router.post('/sendMovieData', postUserMovieData);

module.exports = router;