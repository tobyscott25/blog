# Stage 1 - Build with Node
FROM node:latest as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY ./ .
RUN npm run build

# Stage 2 - Host with Nginx
FROM nginx as production-stage
RUN mkdir /app
COPY --from=build-stage /app/out /app
COPY nginx.conf /etc/nginx/nginx.conf