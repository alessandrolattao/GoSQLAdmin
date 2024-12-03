# GoSQLAdmin

## A Lightweight and Intuitive SQL Web Client Built with Go, HTMX, and DaisyUI

---

![GoSQLAdmin Logo](web/static/images/logo.png)

**⚠️ WARNING: This project is currently under development and is not ready for production use.**

GoSQLAdmin is an evolving web client designed to provide a simple and user-friendly interface for managing SQL databases. Built with the power of **Golang**, enhanced interactivity of **HTMX**, and the modern styling of **DaisyUI**, this project aims to deliver a sleek and efficient experience. However, it is **not functional yet**. Stay tuned as we work toward creating a fully functional tool.

### What to Expect (Eventually):
- Lightweight and fast SQL database management.
- Seamless user interactions powered by HTMX.
- Beautiful and responsive UI styled with DaisyUI.
- Secure connections and easy configuration.

## Environment Variables

| **Variable**          | **Description**                                                                                           | **Type**        | **Default**    |
|-----------------------|-----------------------------------------------------------------------------------------------------------|-----------------|----------------|
| `USER`          | Username for the SQL database.                                                                          | `string`        | -              |
| `PASSWORD`      | Password for the SQL database.                                                                          | `string`        | -              |
| `HOST`          | Hostname or IP address of the SQL server.                                                              | `string`        | -              |
| `PORT`          | Port number on which the SQL server is listening.                                                      | `string`        | `3306`         |
| `CONN_TIMEOUT`  | Timeout for establishing a connection to the SQL server (in seconds).                                  | `time.Duration` | `10s`          |
| `READ_TIMEOUT`  | Timeout for reading data from the SQL server (in seconds).                                             | `time.Duration` | `30s`          |
| `WRITE_TIMEOUT` | Timeout for writing data to the SQL server (in seconds).                                               | `time.Duration` | `30s`          |
| `MAX_OPEN_CONNS`      | Maximum number of open connections to the database.                                                      | `int`           | `100`          |
| `MAX_IDLE_CONNS`      | Maximum number of idle connections in the connection pool.                                               | `int`           | `10`           |
| `CONN_MAX_LIFETIME`   | Maximum lifetime of a connection in the pool before being closed (in seconds).                           | `time.Duration` | `300s`         |

## Docker Compose Example

Below is an example of a `docker-compose.yml` file for running the container:

```yaml
services:
  gosqladmin:
    image: alessandrolattao/gosqladmin
    container_name: gosqladmin
    environment:
      USER: myuser
      PASSWORD: mypassword
      HOST: mydbhost.com
      PORT: 3306
      CONN_TIMEOUT: 10s
      READ_TIMEOUT: 30s
      WRITE_TIMEOUT: 30s
      MAX_OPEN_CONNS: 100
      MAX_IDLE_CONNS: 10
      CONN_MAX_LIFETIME: 300s
    ports:
      - "8080:8080" # Replace 8080 with the correct exposed port if needed
    restart: unless-stopped
```

### Current Status:
The project is in the early development stage, and no features are operational yet. Contributions and feedback are welcome, but **please refrain from using this project in any production environment**.

---

### Stay Tuned:
Follow this repository for updates as development progresses. Exciting features are on the horizon!

---

### Contribution:
If you wish to contribute, feel free to open issues or submit pull requests. Your ideas and efforts can help shape GoSQLAdmin into a great tool for developers.

---

### Resources:
https://go.dev/

https://htmx.org/

https://daisyui.com/

https://docs.docker.com/

https://docs.docker.com/compose/

---
