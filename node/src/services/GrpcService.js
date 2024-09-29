const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const path = require('path');

const PROTO_PATH = path.join(__dirname, '../proto/service.proto');
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const myappProto = grpc.loadPackageDefinition(packageDefinition).myapp;

const movieClient = new myappProto.MovieService(process.env.GRPC_HOST, grpc.credentials.createInsecure());
const authClient = new myappProto.AuthService(process.env.GRPC_HOST, grpc.credentials.createInsecure());

module.exports = {
  movieClient,
  authClient,
};