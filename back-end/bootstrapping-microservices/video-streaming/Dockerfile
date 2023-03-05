FROM node:18.14-alpine

WORKDIR /usr/src/app

COPY package*.json ./
RUN npm install --only=production

COPY ./src ./src
COPY ./localResource ./localResource

CMD npm start