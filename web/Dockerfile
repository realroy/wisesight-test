FROM node:13.11.0-alpine3.11 AS build

WORKDIR /usr/app
COPY . .
RUN yarn --prod
RUN yarn build:static

FROM nginx:1.17.9-alpine

COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf
COPY --from=build /usr/app/out /usr/share/nginx/html
RUN chown nginx.nginx /usr/share/nginx/html/ -R

CMD [ "nginx", "-g", "daemon off;" ]
