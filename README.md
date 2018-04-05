# Go TODO

This is a simple todo web app written in Go and using the Echo Framework

This repository is a fork of [ezynda3/go-echo-vue](https://github.com/ezynda3/go-echo-vue) with the difference of using MySQL instead of SQLite3 for the Database

To get it running using Docker:

#### 1. Clone this repository
```bash
go get github.com/Kamaropoulos/go-echo-vue-mysql
```


#### 2. Create a new Docker network
```bash
docker network create todo-network
```


#### 3. Create a new MySQL container
```bash
docker run --name=todo-mysql --network=todo-network -p 3306:3306 -e MYSQL_ROOT_PASSWORD=rootpass -e MYSQL_DATABASE=todos -e MYSQL_USER=todouser -e MYSQL_PASSWORD=todopass -d mysql:latest 
```


#### 4. Build the application image
```bash
docker build -t go-echo-vue-mysql -f Dockerfile .
```

#### 5. Wait for the MySQL Server to startup


#### 6. Create a Docker container for the TODO app
```bash
docker run --name=todo-app --network=todo-network -p 8000:8000 -e TODODBHOST=todo-mysql -e TODODBUSER=todouser -e TODODBPASS=todopass -d go-echo-vue-mysql
```

#### To try the application, point your browser to http://localhost:8000
