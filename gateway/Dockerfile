FROM node:22-alpine3.19

WORKDIR /app

COPY package*.json .

RUN npm install

COPY src src

ENTRYPOINT [ "node", "src/app.js" ]
