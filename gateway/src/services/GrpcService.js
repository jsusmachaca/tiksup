import { loadPackageDefinition, credentials } from '@grpc/grpc-js'
import { loadSync } from '@grpc/proto-loader'
import { join } from 'path'

const PROTO_PATH = join(__dirname, '../proto/service.proto')
const packageDefinition = loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true
})

const myappProto = loadPackageDefinition(packageDefinition).myapp

const movieClient = new myappProto.MovieService(process.env.GRPC_HOST, credentials.createInsecure())
const authClient = new myappProto.AuthService(process.env.GRPC_HOST, credentials.createInsecure())

export default {
  movieClient,
  authClient
}
