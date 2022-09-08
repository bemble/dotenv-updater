# bemble/dotenv-updater

[![gh_last_release_svg]][gh_last_release_url]
[![tippin_svg]][tippin_url]

[gh_last_release_svg]: https://img.shields.io/github/v/release/bemble/dotenv-updater?sort=semver
[gh_last_release_url]: https://github.com/bemble/dotenv-updater/releases/latest

[tippin_svg]: https://img.shields.io/badge/donate-BuyMeACoffee-ffdd00?logo=buymeacoffee&style=flat
[tippin_url]: https://www.buymeacoffee.com/bemble

> Sometimes you just need a simple API to update a dotenv file, dotenv-updater provides it. It does not provide a read access to dotenv, only update/delete.

**Classic use case:** you tag your production images using real versions (`v1.0.0` etc), not `latest`, to control what is in production. You have your `CD` pipeline but how to update your container?
You could use portainer business or give a try to a simple `docker-compose`:

```yml
version: "3"

services:
    my-app:
        image: user/my-app:${MYAPP_VERSION}
        container_name: dotenv-updater
        restart: unless-stopped

    dotenv-updater:
        image: ghcr.io/bemble/dotenv-updater:latest
        container_name: dotenv-updater
        restart: unless-stopped
        ports:
            - 5000:5000
        environment:
            API_KEY: my-api-key
        volumes:
            - "${PWD}/.env:/app/data/.env"
            - /etc/localtime:/etc/localtime:ro

    watchtower:
        image: containrrr/watchtower
        volumes:
            - /var/run/docker.sock:/var/run/docker.sock
        command: --http-api-update
        environment:
            - WATCHTOWER_HTTP_API_TOKEN=${WATCHTOWER_TOKEN}
        ports:
            - 8080:8080
```

Then, in your `CD` pipeline just call:

```
[POST] http://server:5000/env/MYAPP_VERSION {"value": "v1.2.3"} -> update tag
[POST] http://server:8080/v1/update                             -> update container
```

## Running

### Docker

```bash
docker run -v "portainer/compose/2/stack.env:/app/data/.env" -p5000:5000 -e API_KEY=my-api-key  ghcr.io/bemble/dotenv-updater:latest 
```

### Docker compose

```yml
  dotenv-updater:
    image: ghcr.io/bemble/dotenv-updater:latest
    container_name: dotenv-updater
    restart: unless-stopped
    ports:
      - 5000:5000
    environment:
      API_KEY: my-api-key
      BASE_PATH: "/updater"
      DELETE_ALLOWED: 1
      DEBUG: 1
    volumes:
      - "${PWD}/portainer/compose/2/stack.env:/app/data/.env"
      - /etc/localtime:/etc/localtime:ro
```

## Configuration

### Volumes

- `/app/data/.env`: dotenv file to modify

### Environment variables

- `API_KEY`: api key used to make calls (default `CHANGE-IT-ASAP`)
- `BASE_PATH`: web base path (default `/`)
- `DELETE_ALLOWED`: allow delete vars (default: `false`, anything different of [`1`, `true` or `True`] will set it to `false`)
- `DEBUG`: display debug info (default: `false`, anything different of [`1`, `true` or `True`] will set it to `false`)

## Using the API

Don't forget the `X-Api-Key` header!

- `[GET] /status`: get status of the app
- `[POST] /env/{VAR_NAME}`: set `VAR_NAME` variable value
  
  Body:
```json
{"value": "the new value"}
```
  Replies `204 - No Content` when OK
- `[DELETE] /env/{VAR_NAME}`: delete `VAR_NAME` (if `DELETE_ALLOWED` is `false`, does nothing)
  
  Replies `204 - No Content` when OK

## Development

### Requirements

* Golang version 1.17 minimum must be installed

### Running tests

```bash
go test ./...
```

### Running app

Server will run on port `5000`:

```bash
go run main.go
```

Alternatively, you can use `nodemon` to get server restart on file change:

```bash
npm i -g nodemon
nodemon --exec go run main.go --signal SIGTERM
```

### Build docker image

```bash
docker build . -f Dockerfile -t dotenv-updater
```

## TODO

- [X] disable `DELETE`
- [ ] restrict update to given variables