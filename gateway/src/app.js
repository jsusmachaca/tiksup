import express, { json } from 'express'
import cargarRutas from './routes/Loader.js'
import { iniciarProductor } from './config/kafka.js'
import cors from 'cors'
import 'dotenv/config'

const app = express()
<<<<<<< HEAD
const port = process.env.PORT || 3000
=======
const port = process.env.PORT || 3005
>>>>>>> 07f6615dc8ee75effecab9a511a1eba9ad85afce

app.use(json())
app.use(cors())

cargarRutas(app)

app.get('/', (req, res) => {
  res.send('¡Hello, world!')
})

app.listen(port, async () => {
  console.log(`Express server running on http://localhost:${port}`)
  await iniciarProductor()
})
