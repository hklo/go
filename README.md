# go
This is a development setup for golang

## Getting Started
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
