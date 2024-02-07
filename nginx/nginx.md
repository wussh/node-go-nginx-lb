```nginx
events {
    worker_connections 1024;  # Maximum number of simultaneous connections that each worker process can handle.
}
```
- `events`: This block specifies parameters related to Nginx's event processing model.
- `worker_connections 1024;`: Sets the maximum number of simultaneous connections that each worker process can handle to 1024.

```nginx
http {
    # Define upstream servers for each load balancing strategy.

    # Round Robin load balancing strategy
    upstream appRoundRobin {
        server 127.0.0.1:3021;
        server 127.0.0.1:3022;
        server 127.0.0.1:3023;
    }

    # IP Hash load balancing strategy
    upstream appIPHash {
        ip_hash;  # Ensures that requests from the same IP are always routed to the same server.
        server 127.0.0.1:3021;
        server 127.0.0.1:3022;
        server 127.0.0.1:3023;
    }

    # Least Connections load balancing strategy
    upstream appLeastConn {
        least_conn;  # Routes each request to the server with the fewest active connections.
        server 127.0.0.1:3021;
        server 127.0.0.1:3022;
        server 127.0.0.1:3023;
    }
```
- `http`: This block specifies parameters related to the HTTP protocol.
- `upstream appRoundRobin`: Defines an upstream block named `appRoundRobin` for the Round Robin load balancing strategy, listing the backend servers with their respective ports.
- `upstream appIPHash`: Defines an upstream block named `appIPHash` for the IP Hash load balancing strategy, enabling `ip_hash` to ensure requests from the same IP are always routed to the same server.
- `upstream appLeastConn`: Defines an upstream block named `appLeastConn` for the Least Connections load balancing strategy, enabling `least_conn` to route each request to the server with the fewest active connections.

```nginx
    # Define a rate limiter to restrict requests to a certain rate.
    limit_req_zone $binary_remote_addr zone=ourRateLimiter:10m rate=15r/s; # 1m = 16000 IP

    # Configure proxy cache to cache responses for improved performance.
    proxy_cache_path /var/cache/nginx levels=1:2 keys_zone=cache_one:5m inactive=10m;

    # Define a custom log format for upstream requests.
    log_format upstreamlog 'remote_addr: $remote_addr | '
        'remote_user: $remote_user | '
        'time_local: $time_local | '
        'request: $request | '
        'status: $status | '
        'body_bytes_sent: $body_bytes_sent | '
        'http_referer: $http_referer | '
        'upstream_addr: $upstream_addr | '
        'upstream_response_time: $upstream_response_time | '
        'request_time: $request_time | '
        'msec: $msec | '
        'http_user_agent: $http_user_agent';
```
- `limit_req_zone $binary_remote_addr zone=ourRateLimiter:10m rate=15r/s;`: Defines a rate limiter named `ourRateLimiter`, using the client's IP address (`$binary_remote_addr`) to track the limit. It allows 15 requests per second per IP address.
- `proxy_cache_path /var/cache/nginx levels=1:2 keys_zone=cache_one:5m inactive=10m;`: Configures the proxy cache to store cached responses in `/var/cache/nginx`, with a cache size of 5 MB (`5m`). Cached items are considered inactive after 10 minutes of no usage.
- `log_format upstreamlog ...`: Defines a custom log format named `upstreamlog` for logging upstream requests, capturing various request and response details.

```nginx
    server {
        listen 80;  # Listen for incoming connections on port 80.

        # Define different locations and their corresponding configurations.

        location / {
            # Apply rate limiting to requests in this location.
            limit_req zone=ourRateLimiter;
            limit_req_status 429;  # Return 429 status code for requests exceeding the rate limit.
            
            # Set a custom header before passing the request to the upstream server.
            proxy_set_header X-Custom-Header this-was-custom-header;
            
            # Route requests to the upstream server using the Round Robin strategy.
            proxy_pass http://appRoundRobin;
        }
```
- `server { listen 80; ... }`: Defines a server block to handle requests on port 80.
- `location / { ... }`: Defines a location block for requests to the root URI.
- `limit_req zone=ourRateLimiter;`: Applies rate limiting to requests in this location using the previously defined rate limiter.
- `limit_req_status 429;`: Returns a status code of 429 (Too Many Requests) for requests exceeding the rate limit.
- `proxy_set_header X-Custom-Header this-was-custom-header;`: Sets a custom header `X-Custom-Header` before passing the request to the upstream server.
- `proxy_pass http://appRoundRobin;`: Routes requests to the upstream server `appRoundRobin` using the Round Robin strategy.

```nginx
        location /ipHash {
            # Log detailed information about requests using the custom log format.
            access_log /var/log/app/app.log upstreamlog;
            
            # Route requests to the upstream server using the IP Hash strategy.
            proxy_pass http://appIPHash;
        }

        location /leastConn {
            # Set a custom header before passing the request to the upstream server.
            proxy_set_header X-Custom-Header this-was-custom-header;
            
            # Route requests to the upstream server using the Least Connections strategy.
            proxy_pass http://appLeastConn;
        }

        location /metadata {
            # Configure proxy cache for caching responses.
            proxy_cache cache_one;
            proxy_cache_min_uses 5;
            proxy_cache_methods HEAD GET;
            proxy_cache_valid 200 304 30s;
            proxy_cache_key $uri;
            
            # Route requests to the upstream server without specifying a load balancing strategy.
            proxy_pass http://app;
        }
    }
}
```
- `location /ipHash { ... }`: Defines a location block for requests to the `/ipHash` URI.


- `access_log /var/log/app/app.log upstreamlog;`: Logs detailed information about requests to `/ipHash` using the custom log format `upstreamlog`.
- `proxy_pass http://appIPHash;`: Routes requests to the upstream server `appIPHash` using the IP Hash strategy.

- `location /leastConn { ... }`: Defines a location block for requests to the `/leastConn` URI.
- `proxy_pass http://appLeastConn;`: Routes requests to the upstream server `appLeastConn` using the Least Connections strategy.

- `location /metadata { ... }`: Defines a location block for requests to the `/metadata` URI.
- `proxy_cache cache_one;`: Configures proxy caching using the `cache_one` zone defined earlier.
- `proxy_cache_min_uses 5;`: Specifies that an item should be cached after it has been requested at least 5 times.
- `proxy_cache_methods HEAD GET;`: Specifies that only `HEAD` and `GET` requests are eligible for caching.
- `proxy_cache_valid 200 304 30s;`: Specifies that cached items with a `200` or `304` status code are considered valid for 30 seconds.
- `proxy_cache_key $uri;`: Defines the cache key based on the requested URI.
- `proxy_pass http://app;`: Routes requests to the upstream server `app` without specifying a load balancing strategy.
