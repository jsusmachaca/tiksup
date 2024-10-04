import { readdirSync } from 'fs'
import { join } from 'path'

const cargarRutas = (app) => {
  const rutasDir = join(process.cwd(), 'src', 'routes')

  readdirSync(rutasDir).forEach(async (archivo) => {
    if (archivo !== 'routeLoader.js' && archivo.endsWith('.js')) {
      const rutaNombre = archivo.replace('.js', '')
      const ruta = await import(join(rutasDir, archivo))
      app.use(`/${rutaNombre.toLowerCase()}`, ruta.default)
    }
  })
}

export default cargarRutas
