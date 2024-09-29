const express = require('express');
const { getMovies } = require('../controllers/MovieController');

const router = express.Router();

router.get('/', getMovies);

module.exports = router;