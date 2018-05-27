# Gocatpus

<div style="text-align:center"><img src ="https://preview.ibb.co/g4botc/Polpo_disegno.png" /></div>

Goctapus is a wrapper around Echo that makes it easier to build REST APIs and applications that are based on them.

For a sample application built on Goctapus, check [Kamaropoulos/goctapus-example](https://github.com/Kamaropoulos/goctapus-example)

There is also a fork ([Kamaropoulos/goctapus-mongo](https://github.com/Kamaropoulos/goctapus-mongo)) that uses MongoDB as the Database Backend but it is currently outdated and should be soon merged to this one.

This repository was originaly a fork of [ezynda3/go-echo-vue](https://github.com/ezynda3/go-echo-vue) in order to make it use MySQL as a database backend but it's purpose has changed a lot since then.

## Getting started

Goctapus makes it really easy and fast to get started and build your own REST APIs.

All you have to do is:

0. ### Install Go if you haven't already,
1. ### Clone or copy the files from [Kamaropoulos/goctapus-blank-template](https://github.com/Kamaropoulos/goctapus-blank-template)
2. ### Connect to a database.

    Goctapus will connect to the MySQL server using the information passed to the command line. You can open connections to one or multiple MySQL databases using `ConnectDB()`. For example:

    ```golang
    goctapus.ConnectDB("goapp")
    ```

    will connect to the databased named `goapp` on the MySQL server provided on the command line.

3. ### Create your API endpoints:
    You can create endpoints using `AddEndpoint()`. For example:
    
    ```golang
    goctapus.AddEndpoint(goctapus.Route{
		Method:  "GET",
		Path:    "/tasks",
		Handler: handlers.GetTasks(goctapus.Databases["goapp"]),
		Rate:    10,
	}, myAwesomeMiddleware1(), myAwesomeMiddleware2())
    ```

    will add a GET endpoint on `/tasks` using the handler `GetTasks()` with a rate limit of 10 requests per second. It will also register two middlewares.
    
    You can also serve static files using `AddStatic()`. For example:

    ```golang
    goctapus.AddStatic(goctapus.Route{
		Path: "/",
		File: "public/index.html",
    })
    ```
    
    will serve the file `index.html` when the root path of the server is requested.

4. ### Create your handlers and models

    Create your handler and model files on the `handlers` and `models` folders. For more information and examples you can refer to the Echo documentation and the [Kamaropoulos/goctapus-example](https://github.com/Kamaropoulos/goctapus-example) application.

5. ### Run your application!

    To execute your application you can run:
    ```bash
    go run main.go [application_Port] [DB_Username] [DB_Password] [DB_Hostname/IP] [DB_PORT]
    ```

    where you'll have to replace:
    - `[application_Port]` with the port you want the application to listen to,
    - `[DB_Username]` with the username the application is going to connect to the Database with,
    - `[DB_Password]` with the password the application is going to connect to the Database with,
    - `[DB_Hostname/IP]` with the hostname or the IP the Database is running on,
    - `[DB_Port]` with the Port the Database is listening to.

### To try the application, point your browser to localhost:[applicationPort]