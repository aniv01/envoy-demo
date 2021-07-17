# envoy-demo

### Envoy Architecture
---
- Downstream/Upstream
- Listeners
    - Listens on a port for downstream clients
    - Network filters are applied to Listeners
        - TCP/HTTP Filters
        - Transport Sockerts TLS
- Network Filters
    - How Envoy maps Listeners and Clusters
    - TCP Proxy Network Filter
        - `envoy.filters.network.tcp_proxy`
    - HTTP Proxy Network Filter
        - `envoy.filters.network.http_connection_manager`
- Clusters
- Threading Model
    - Single process multi-threaded model
    - Each thread is bound to a single connection
    - No coordination between threads
    - Command to run envoy process with a single thread `envoy -c envoy.yaml --concurrency 1`.
    ![Screenshot 2021-07-17 at 12.18.34 AM](https://i.imgur.com/2JGJkvo.png)

- Connection Pools
    - Each host in a cluster gets 1 or more connection pools
    - Each protocol gets a pool HTTP 1.1, HTTP/2
    - More pools allocated per priority or socket options
    - Connection pools are per worker thread
    ![Screenshot 2021-07-17 at 12.12.07 AM](https://i.imgur.com/XSMrhOK.png)

- **Note** Envoy recommends using HTTP/2

http://app.demo.com:8080/
http://app.demo.com:4443/
