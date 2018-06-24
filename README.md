# go
This is a development setup for golang. Idea got from [article](https://mikemadisonweb.github.io/2018/03/06/go-autoreload)

## Getting Started

### Features
* Live reload [realize](https://github.com/oxequa/realize)
* Dependency management [dep](https://golang.github.io/dep)

### What you’ll need
* Docker

### Folder structure
```
go/
├── docker/ (Docker setup with golang environment)
│   ├── entrypoint.sh
│   ├── Dockerfile
├── src (Golang projects)
├── .env
├── .gitignore
├── docker-compose.yml
```

### Run Command
From **project** root execute:

    $ docker-compose up
