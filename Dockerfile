FROM ubuntu:22.04 as builder
WORKDIR /app
COPY . .
RUN apt update && apt install -y hugo
RUN hugo

FROM nginx:1.23.2-alpine as server
RUN mkdir /public_html
COPY --from=builder /app/public /public_html
COPY nginx.conf /etc/nginx/nginx.conf