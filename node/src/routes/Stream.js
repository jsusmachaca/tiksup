const express = require('express');
const { postUserMovieData } = require('../controllers/StreamController');

const router = express.Router();

router.post('/sendMovieData', postUserMovieData);

module.exports = router;