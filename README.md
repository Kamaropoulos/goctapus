# Gocatpus

![goctapus logo](https://preview.ibb.co/g4botc/Polpo_disegno.png)

Goctapus is a wrapper around Echo that makes it easier to build REST APIs and applications that are based on them.

It implements a TODO App for demonstration purposes but it can be easily converted into any kind of application.

This repository was originaly a fork of [ezynda3/go-echo-vue](https://github.com/ezynda3/go-echo-vue) but it is now using MySQL instead of SQLite3 for the Database and it's purpose is to abstract as many things as possible to make the development of web applications as fast and easy as possible.

### To get it running using Docker:

#### 1. Clone this repository
```bash
git clone https://github.com/Kamaropoulos/goctapus.git && cd goctapus
```


#### 2. Build and run it using Docker Compose
```bash
docker-compose up -d
```

### If you don't have Docker installed:

#### 1. Clone this repository
```bash
git clone https://github.com/Kamaropoulos/goctapus.git && cd goctapus
```

#### 2. Run the server using your MySQL Database connection information
```bash
go run main.go [application_Port] [DB_Username] [DB_Password] [DB_Hostname/IP] [DB_PORT]
```

where you'll have to replace:
- `[application_Port]` with the port you want the application to listen to,
- `[DB_Username]` with the username the application is going to connect to the Database with,
- `[DB_Password]` with the password the application is going to connect to the Database with,
- `[DB_Hostname/IP]` with the hostname or the IP the Database is running on,
- `[DB_Port]` with the Port the Database is listening to.

#### To try the application, point your browser to http://localhost:[applicationPort]
