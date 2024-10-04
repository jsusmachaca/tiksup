import { createClient } from 'redis'
import 'dotenv/config'

export const client = createClient({
  url: `redis://${process.env.REDIS_HOST}:${process.env.REDIS_PORT}`
})

client.on('error', err => console.log('Redis Client Error', err))

async function connectRedis () {
  await client.connect()
}

connectRedis()
