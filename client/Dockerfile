# stage1 - build react app
FROM node:16-alpine as build
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY ./package.json /app/
COPY ./yarn.lock /app/
#COPY ./.env /app/
RUN yarn --network-timeout 100000
COPY . /app
RUN yarn build

# stage 2 - build the final image and copy the react build files
FROM nginx:stable-alpine
COPY --from=build /app/build /usr/share/nginx/html
#COPY --from=build /app/.env /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
