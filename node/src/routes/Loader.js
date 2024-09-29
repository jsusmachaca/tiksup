const fs = require('fs');
const path = require('path');

const cargarRutas = (app) => {
  const rutasDir = __dirname; 
  fs.readdirSync(rutasDir).forEach((archivo) => {
    if (archivo !== 'routeLoader.js' && archivo.endsWith('.js')) {
      const rutaNombre = archivo.replace('.js', '');
      const ruta = require(path.join(rutasDir, archivo));
      app.use(`/${rutaNombre}`, ruta);
    }
  });
};

module.exports = cargarRutas;