# IEEE RVCE URL Shortner

Made using Golang and Redis (Key Value Store) to enable fast access to sb records and low latency redirection.

## How to run

- Bring Up Redis Instance

```bash
docker-compose up -d
```

- Build the artifacts

```bash
go build .
```

- Run the server

```bash
./github.com/IEEE-RVCE/url-shortner
```
