### Rest API for timing Enduro/Bike races

### Configure Environment Variables

Open .env file and edit the values if you need to. This project uses [godotenv](https://github.com/joho/godotenv) to read and use .env file. 

Database config string is formatted in [go-sql-driver format](https://github.com/go-sql-driver/mysql#parameters).

## Start Project

```
$ go run main.go
```

To explicitly compile the code before you run the server:

```
$ go build main.go
$ ./main
```
