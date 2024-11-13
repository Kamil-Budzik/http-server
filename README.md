Documentation for the HTTP Server Application
Overview
This application implements a basic HTTP server that handles GET and POST requests. It provides a variety of endpoints to interact with and perform operations, such as echoing messages, retrieving user-agent information, and handling file operations. The server is built in Go and is designed to learn how to handle network connections, implement request handlers, and test Go applications with 100% test coverage.

Key Features
GET Requests:
/: Responds with a 200 status and an empty body.
/echo/{message}: Returns the message parameter passed in the path.
/user-agent: Responds with the User-Agent header of the request.
/files/{filename}: Retrieves the contents of the specified file from the server's directory.
POST Requests:
/files/{filename}: Accepts a file upload and writes the content to a file on the server in the specified directory.
How to Use
Running the Server:

The server listens on localhost:4221 by default. You can change the directory where files are stored by passing the -directory flag when starting the server.
go
Copy code
go run main.go -directory "/path/to/directory"
GET Endpoints:

GET /: Returns an empty response body with a 200 OK status.
GET /echo/{message}: Echoes back the {message} part of the URL as the body. For example:
bash
Copy code
GET /echo/hello -> "hello"
GET /user-agent: Returns the value of the User-Agent header sent with the request.
GET /files/{filename}: Returns the contents of the specified file. If the file is not found, it returns a 404 Not Found error.
POST Endpoints:

POST /files/{filename}: Accepts a file upload and writes the content to a file in the server's specified directory. For example:
bash
Copy code
POST /files/test.txt with body: "Hello, world!"
This will create a test.txt file in the specified directory with the content Hello, world!.
Implementation
Connection Handling:
The server accepts incoming TCP connections, reads the request, parses it, and routes it to the appropriate handler (either GET or POST). The server also sends the appropriate HTTP response based on the request.

Request Handlers:

GET Request Handlers: The server supports several GET endpoints, including the root, echo, user-agent, and file retrieval.
POST Request Handlers: The server supports file uploads via the POST method. It writes the received file contents to the server's directory.
Server Lifecycle:
The server runs continuously, handling multiple concurrent connections. It can be gracefully shut down via the ShutdownChan channel, which signals the server to close the listener and clean up.

Testing
Test Coverage:
The application has 100% test coverage. The tests cover the request handling logic for different endpoints, including GET and POST requests. This helps ensure that the server works as expected and responds with the correct status and body.

Learning Testing in Go:
By implementing tests for this project, you will learn how to:

Write unit tests for HTTP handlers.
Mock network connections and simulate requests.
Ensure the server responds correctly based on the request method and path.
What I Learned
Working with TCP and HTTP:
I gained practical experience with handling raw TCP connections, parsing HTTP request lines, and constructing HTTP responses in Go.

Request and Response Handling:
I implemented different request handlers based on the HTTP method (GET/POST) and learned how to manage the body and headers of HTTP requests and responses.

File Operations:
I learned how to handle file operations (read and write) in Go, which is useful for server-based file upload and download functionality.

Concurrency:
By using goroutines to handle multiple client connections concurrently, I learned how to handle concurrent network requests in Go.

Testing in Go:
The project was an opportunity to learn how to write unit tests for Go code, ensuring that my HTTP server behaves correctly under various scenarios. This practice helped me understand how to mock data and test different functionalities in Go applications.

Conclusion
This project provided a great introduction to building an HTTP server in Go, handling network connections, processing HTTP requests, and testing the server with 100% coverage. It improved my understanding of networking, file handling, and server-side programming in Go while emphasizing best practices like concurrent connection handling and testing.

