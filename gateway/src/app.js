import express, { json } from 'express'
import cargarRutas from './routes/Loader.js'
import { iniciarProductor } from './config/kafka.js'
import cors from 'cors'
import 'dotenv/config'

const app = express()
const port = process.env.PORT || 3005

app.use(json())
app.use(cors())

cargarRutas(app)

app.get('/', (req, res) => {
  res.send('¡Hola, mundo!')
})

app.listen(port, async () => {
  console.log(`Servidor Express corriendo en http://localhost:${port}`)
  await iniciarProductor()
})
