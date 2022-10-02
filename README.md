<h1 align="center">Go-Book-API</h1>

<p align="center">
Reads from JSON and displays a list of books. 
</p>

## Environment Variables
The environment variables are key in relation to the mysql database.

```bash
MYSQL_ROOT_PASSWORD=
MYSQL_ROOT_USER=
MYSQL_DATABASE=
MYSQL_HOST=
```
## How to run

```bash
go build . 

go run . 
```
### Using Docker

```bash
docker compose up
```


Navigate to `http://localhost:10000/`