# Date Time Server
------------------

## Description:
minimal dateTime server, returns current time

### Endpoints:
- `GET /datetime`
Response : Current date and time

### Installition:
```
git clone https://github.com/codescalersinternships/datetime-server-salmaelsoly
```
- Individual Docker Image
    - Build Docker Image
    ```
    docker build -f Dockerfile.http -t http-server .
    docker build -f Dockerfile.gin -t gin-server .
    ```
    - Run Docker Image:
        - With docker randomly picking ports
            ```
            docker run -d -P http-server
            docker run -d -P gin-server
            ```
        - With specfied port mapping
            ```
            docker run -d -p <host-port>:<container-port> image
            ```
            Example:
            ```
            docker run -d -p 3000:8080 http-server
            docker run -d -p 5000:8080 gin-server
            ```

- Use docker-compose:
    ```
    docker compuse up
    ```
 - Test The Api:
    - curl the ip:port/datetime
        ```
        curl -q -O - 172.17.0.2:3000/datetime
        curl localhost:3000/datetime
        ```