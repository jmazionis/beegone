# Beegone

Beegone is a CRUD app built to manage vehicle licence plate registrations.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.


### Prerequisites

- Install [Node.js](https://nodejs.org/en/)
- Install [Yarn](https://yarnpkg.com/lang/en/docs/install/)
- Install [Docker](https://docs.docker.com/install/)
- Install [Docker Compose](https://docs.docker.com/compose/install/)
- You will need [Go v1.12+ installed and configured](https://golang.org/dl/).
- Make sure you have the [GOPATH environmental variable set up](https://github.com/golang/go/wiki/SettingGOPATH) 

### Installing

- Clone the project to any directory outside GOPATH

```
$ git clone https://github.com/ICanHaz/beegone.git
```

## Project's structure

The project's structure follows conventions as described in [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

App is split into 2 binaries with entry points located at:
```
cmd/
├── api
│   └── main.go
├── server
    └── main.go
```

- **api** acts as a REST api server accessible on http://localhost:9090/ by default
- **server** acts as a static content server (serves SPA in this case) accessible on http://localhost:9000/ by default

### Api endpoints

Api consists of 5 endpoints:

- (GET) /api/carplates - Gets all carplates
```
curl -X GET \
  http://localhost:9090/api/carplates \
  -H 'Content-Type: application/json'
```


- (POST) /api/carplates - Adds a new carplate
```
curl -X POST \
  http://localhost:9090/api/carplates \
  -H 'Content-Type: application/json'
  -d '{
	"plateId": "AAA-300",
	"modelName": "Batmobile Outback",
	"modelYear": "1990",
	"owner": "Driver 2"
}
```

- (GET) /api/carplates/:id - Retrieves carplate with given id
```
curl -X GET \
  http://localhost:9090/api/carplates/1Q4GX5yPW3KxDuJ9k9XIplGwCTh \ 
  -H 'Content-Type: application/json' 
```

- (PUT) /api/carplates - Updates carplate
```
curl -X PUT \
  http://localhost:9090/api/carplates \
  -H 'Content-Type: application/json' 
  -d '{
    "id": "1PLaFaOOuyXv3ja3eLfUPyDy3de",
    "plateId": "AAA-201",
    "modelName": "Old car",
    "modelYear": "2000",
    "owner": "Driver 2"
}'
```

- (DELETE) /api/carplates/:id

```
curl -X DELETE \
  http://localhost:9090/api/carplates/1PLaC8WRVwz08V0ZUEekSSQViUP \
  -H 'Content-Type: application/json'  
```

# Running

The simplest way to run the app is with Docker Compose:

```
$ docker-compose up -d
```

Go to http://localhost:9000/ to see results 

## Testing

- To run integration tests, run:

```
make run-integration-tests
```


- To run unit tests, run:

```
make run-unit-tests
```

## Built With

* [GO](https://golang.org/)
* [Beego](https://beego.me/)
* [React](https://reactjs.org/)
* [Typescript](https://www.typescriptlang.org/)
* [Docker](https://www.docker.com/)
