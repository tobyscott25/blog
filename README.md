# tobyscott.dev

## About

A personal website built using:
- [Next.js](https://nextjs.org/) for SSG & SEO,
- [TypeScript](https://www.typescriptlang.org/) for static typing,
- and [Docker](https://www.docker.com/) with [Nginx](https://www.nginx.com/) base to serve the statically generated site in production.


## Installation and setup

Install dependencies:
```bash
$ npm install
```

## Development setup

Start the development server:
```bash
$ npm run dev
```

## Production deployment

Build the Docker image:
```bash
$ docker build -t <username>/tobyscott.dev .
```

Run the Docker container:
```bash
$ docker run -p 8080:80 -d <username>/tobyscott.dev
```