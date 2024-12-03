
# GoSQLAdmin

## A Lightweight and Intuitive SQL Web Client Built with Go, HTMX, and DaisyUI

---

![GoSQLAdmin Logo](web/static/images/logo.png)

**⚠️ WARNING: This project is currently under development and is not ready for production use.**

GoSQLAdmin is an evolving web client designed to provide a simple and user-friendly interface for managing SQL databases. Built with the power of **Golang**, enhanced interactivity of **HTMX**, and the modern styling of **DaisyUI**, this project aims to deliver a sleek and efficient experience. However, it is **not functional yet**. Stay tuned as we work toward creating a fully functional tool.

### What to Expect (Eventually)
- Lightweight and fast SQL database management.
- Support for multiple SQL database types.
- Seamless user interactions powered by HTMX.
- Beautiful and responsive UI styled with DaisyUI.
- Secure connections and easy configuration.

---

## Supported Databases

GoSQLAdmin aims to support the following databases:

- **MySQL**
- **PostgreSQL**
- **SQLite**
- **Microsoft SQL Server**
- **Snowflake**
- **ClickHouse**

---

## Environment Variables

| **Variable**           | **Description**                                                                                           | **Type**        | **Default**    |
|-------------------------|-----------------------------------------------------------------------------------------------------------|-----------------|----------------|
| `USER`                 | Username for the SQL database.                                                                            | `string`        | -              |
| `PASSWORD`             | Password for the SQL database.                                                                            | `string`        | -              |
| `HOST`                 | Hostname or IP address of the SQL server.                                                                | `string`        | -              |
| `PORT`                 | Port number on which the SQL server is listening.                                                        | `string`        | `3306`         |
| `DATABASE`             | Name of the initial database to connect to.                                                              | `string`        | -              |
| `DRIVER`               | The database driver to use (`mysql`, `postgres`, `sqlite`, `sqlserver`, `snowflake`, `clickhouse`).      | `string`        | `mysql`        |
| `CONN_TIMEOUT`         | Timeout for establishing a connection to the SQL server (in seconds).                                    | `time.Duration` | `10s`          |
| `READ_TIMEOUT`         | Timeout for reading data from the SQL server (in seconds).                                               | `time.Duration` | `30s`          |
| `WRITE_TIMEOUT`        | Timeout for writing data to the SQL server (in seconds).                                                 | `time.Duration` | `30s`          |
| `MAX_OPEN_CONNS`       | Maximum number of open connections to the database.                                                      | `int`           | `100`          |
| `MAX_IDLE_CONNS`       | Maximum number of idle connections in the connection pool.                                               | `int`           | `10`           |
| `CONN_MAX_LIFETIME`    | Maximum lifetime of a connection in the pool before being closed (in seconds).                           | `time.Duration` | `300s`         |
| `SSL_MODE`             | SSL mode for PostgreSQL (e.g., `disable`, `require`).                                                    | `string`        | `disable`      |
| `SNOWFLAKE_WAREHOUSE`  | Snowflake warehouse name (applicable for Snowflake only).                                                | `string`        | -              |
| `SNOWFLAKE_SCHEMA`     | Snowflake schema name (applicable for Snowflake only).                                                   | `string`        | -              |

---

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
      DATABASE: mydatabase
      DRIVER: mysql
      CONN_TIMEOUT: 10s
      READ_TIMEOUT: 30s
      WRITE_TIMEOUT: 30s
      MAX_OPEN_CONNS: 100
      MAX_IDLE_CONNS: 10
      CONN_MAX_LIFETIME: 300s
      SSL_MODE: disable
    ports:
      - "8080:8080" # Replace 8080 with the correct exposed port if needed
    restart: unless-stopped
