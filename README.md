# go
This is a development setup for golang. Idea got from [article](https://mikemadisonweb.github.io/2018/03/06/go-autoreload)

## Getting Started
Project setup using [realize](https://github.com/oxequa/realize) with live reload feature.

### What you’ll need
* Docker

### Folder structure
```
go/
├── docker/ (Docker setup with golang environment)
│   ├── root/
│   │   ├── entrypoint.sh
│   ├── Dockerfile
├── src (Golang projects)
├── .env
├── docker-compose.yml
```

### Run Command
From **project** root execute:

    $ docker-compose up
