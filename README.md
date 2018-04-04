# Go TODO

This is a simple todo web app written in Go and using the Echo Framework

This repository is a fork of [ezynda3/go-echo-vue](https://github.com/ezynda3/go-echo-vue) with the difference of using MySQL instead of SQLite3 for the Database

To get it running:

#### 1. Clone this repository and Echo
```bash
go get github.com/Kamaropoulos/go-echo-vue-mysql
go get github.com/labstack/echo
```


#### 2. MySQL Database
If you have MySQL up and running you can skip this step

If not, the following commands will start a Docker MySQL container

Don't forget to change your password, you're going to need it for the next step
```bash
docker run -p 3306:3306 --name todo-mysql -e MYSQL_ROOT_PASSWORD=<PASSWORD> -d mysql:latest
```


#### 3. Set MySQL username and password
You will have to change the username and password on the database connection string on [todo.go](todo.go#L13)


#### 4. Run the server
```bash
go run todo.go
```


##### To see the application running, point your browser to http://localhost:8000
