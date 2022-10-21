# IEEE RVCE URL Shortner

Made using Golang and Redis (Key Value Store) to enable fast access to db records and low latency redirection.

## How to run

make sure you have go installed and docker installed

- Bring Up Redis Instance

```bash
docker-compose up -d
```

- Download all the dependencies

```bash
go mod download
```

- Build the artifact

```bash
go build .
```

- Run the server

```bash
./url-shortner
```

## How this works

TBD

## API DOC

There are 3 endpoints,

Development Host is `http://localhost:8080`

Production Host is hided to avoid spam

- POST /api/encode

Payload

```json
{
  "long_url": "https://ieee-rvce.org",
  "user_id": "UwQPr3aIf9dM5x7r",
  "custom_text": "main_website"
}
```

Response

```json
{
  "short_url": "http://localhost:8080/main_website"
}
```

You might ask why user_id is required, it is required to construct the short url. When the custom_text is not provided, the short url is constructed using the user_id and the custom_text and the short url is unique for each user and custom_text combination.

- GET /api/decode/:short_url

Response

```json
{
  "long_url": "https://ieee-rvce.org",
  "short_url": "http://localhost:8080/main_website"
}
```

- GET /api/:short_url

Redirects to the long url, passing 301 status code.

## TBD

[] Dockerise the whole thing;
[] Check If the short_url is valid iff redirect; if not then display a page with proper error message
[] Right Now we are storing the data for 6 hours, should see into making it store permanently
[] Should add a way to delete the data
[] Should add a way to update the data
[] Analytics? (Is it possible to track no of clicks, location , etc...)
[] Maybe a UI to add Data
