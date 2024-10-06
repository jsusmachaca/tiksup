import { Kafka } from 'kafkajs'
import 'dotenv/config'

const kafka = new Kafka({
  clientId: 'myapp',
  brokers: [process.env.KAFKA_SERVER]
})

export const producer = kafka.producer()

export const iniciarProductor = async () => {
  await producer.connect()
  console.log('Kafka producer connected')
}
