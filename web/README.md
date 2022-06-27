# Workflow

## Getting Started

```bash
yarn install
```

```bash
yarn dev
```

## Using Docker

```bash
# build the image
docker build -t workflow .
# and run it
docker run -d -p 3000:3000 --name workflow workflow:latest
```
