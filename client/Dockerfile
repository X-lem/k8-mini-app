FROM node:current-alpine3.17 AS build

WORKDIR /client

COPY package.json .
RUN npm install
COPY . .
RUN npm run build

FROM nginx
EXPOSE 80
COPY --from=build /client/build /usr/share/nginx/html