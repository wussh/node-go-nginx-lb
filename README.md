## Node.js-Go-Nginx-Load-Balancing

This repository provides a comprehensive setup for load balancing Node.js and Go applications using Nginx. Below, you'll find detailed information on the features, usage instructions, and configuration details:

### Features

1. **Load Balancer**: Nginx efficiently distributes incoming requests among multiple Node.js and Go instances.
2. **Load Balancing Strategy**: Implements three strategies for load balancing:
   - Round Robin
   - IP Hash
   - Least Connections
3. **Rate Limiting**: Requests are restricted to a predefined rate using Nginx's `limit_req_zone`.
4. **Caching**: Nginx is configured to cache responses, enhancing overall performance.
5. **Circuit Breaker**: Although not explicitly mentioned, a circuit-breaking mechanism can be implemented using Nginx directives.
6. **Logging**: Detailed logging is set up with a custom log format in Nginx, aiding in monitoring and debugging.

### Usage

#### SSH into Nginx VM (Virtual Machine)

```bash
vagrant ssh nginx
```

#### Node.js

1. Start the Node.js application using PM2:

```bash
pm2 start index.js --name "{app-name}" -- {port} {app-name}
```

Example:

```bash
pm2 start index.js --name "app-01" -- 3022 app-1
pm2 start index.js --name "app-02" -- 3023 app-2
pm2 start index.js --name "app-03" -- 3024 app-3
```

Three servers are set up for load balancing.

#### Golang

1. Run the Go application:

```bash
go run main.go <port> <app-name>
```

Example:

```bash
go run main.go 3022 app-1
go run main.go 3023 app-2
go run main.go 3024 app-3
```

### Configuration Details

- **Nginx Configuration**: The `nginx.conf` file located in the `nginx` directory contains essential settings:
  - Defines upstream servers.
  - Specifies load balancing strategies, rate limiting, caching, logging formats, and server blocks for request routing.
- **Node.js Configuration**: Dockerfiles and docker-compose files in the `nodejs` directory provide configurations for running Node.js applications.
- **Go Configuration**: Dockerfiles and docker-compose files in the `go` directory offer configurations for running Go applications.

### Notes

- Ensure to configure and adjust parameters according to your application's specific requirements.
- Additional features like SSL termination, health checks, and security measures can be incorporated as needed.

This setup provides a robust foundation for load balancing Node.js and Go applications, enhancing scalability, reliability, and performance.