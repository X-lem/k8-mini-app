FROM node:current-alpine3.17 AS build

WORKDIR /client

COPY package.json .
RUN npm install
COPY . .
RUN npm run build

FROM nginx

COPY --from=build /client/build /usr/share/nginx/html

# Overwrite the default NGINX configuration file with our custom file
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
