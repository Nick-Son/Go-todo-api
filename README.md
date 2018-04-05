# RESTful JSON API

*Coded in Go*

A simple project to get familiar with building web servers in Go. More info to come...

*Note:*
As the project has been split into multiple files,  `go run main.go` will return a  `# command-line-arguments` error. To run the project, you need to build it into an executable by running:

```bash
go build
```

This will create a go script, which can be executed by running:

```bash
./*scriptname*
```

## Adding todo's
As there currently is no front end, to see if the ability to create todos is functional, we can send a post request with curl:
```bash
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
```

If you refresh the  `http://localhost:8080/todos` page, you will see the new todo entry. 