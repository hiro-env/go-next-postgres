FROM node:18-alpine

WORKDIR /front

COPY package.*json ./
COPY yarn.lock ./

RUN yarn install

EXPOSE 3000

CMD ["yarn", "dev"]
