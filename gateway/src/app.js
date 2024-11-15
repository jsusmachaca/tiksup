import express, { json } from 'express'
import routesLoader from './routes/Loader.js'
import { initProductor } from './config/kafka.js'
import cors from 'cors'
import 'dotenv/config'

const app = express()

const port = process.env.PORT || 3000

app.use(json())
app.use(cors())

routesLoader(app)

app.get('/', (req, res) => {
  res.send('¡Hello, world!')
})

app.listen(port, async () => {
  console.log(`Express server running on http://localhost:${port}`)
  await initProductor()
})
