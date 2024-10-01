import { Kafka } from 'kafkajs';
import 'dotenv/config'

console.log(process.env.KAFKA_BROKER)
const kafka = new Kafka({
  clientId: 'myapp',
  brokers: [process.env.KAFKA_BROKER],
});

export const producer = kafka.producer();

export const iniciarProductor = async () => {
  await producer.connect();
  console.log('Productor Kafka conectado');
};
