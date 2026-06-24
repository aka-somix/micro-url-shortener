# Mirco URL-Shortener
Summary here

## Commands

```
go tool air
```

## Deployment

See [docs/k8s-deploy.md](docs/k8s-deploy.md) for the full guide to deploy on k3s with Cloudflare.

## Docker

```
docker build -f build/Dockerfile -t micro-url-shortener:latest .
```

```
docker run -p 8080:8080 \
  -e BASE_URL=http://localhost:8080 \
  -e REDIS_URL=<host>:6379 \
  -e REDIS_PASSWORD=<password> \
  micro-url-shortener:latest
```