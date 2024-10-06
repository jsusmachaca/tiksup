import { readdirSync } from 'fs'
import { join } from 'path'

const routesLoader = (app) => {
  const dirsPath = join(process.cwd(), 'src', 'routes')

  readdirSync(dirsPath).forEach(async (file) => {
    if (file !== 'routeLoader.js' && file.endsWith('.js')) {
      const pathName = file.replace('.js', '')
      const path = await import(join(dirsPath, file))
      app.use(`/${pathName.toLowerCase()}`, path.default)
    }
  })
}

export default routesLoader
