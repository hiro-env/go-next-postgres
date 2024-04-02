FROM node:18-alpine as builder

WORKDIR /front

COPY package.json .

RUN yarn install

COPY . .

RUN yarn build

EXPOSE 3000

# CMD ["yarn", "start"]
CMD ["yarn", "dev"]

