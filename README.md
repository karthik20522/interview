# League - API Assignment

##  About
Basic web server to perform matrix calculations

### Prerequisites
- git
- go

### Installation

```
$ git clone -b league https://github.com/karthik20522/interview.git
```

### Commands

-   Run test cases
```
$ go test .
```

-   Run
```
$ go run .
```

-   API
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/v1/echo"
curl -F 'file=@/path/matrix.csv' "localhost:8080/v1/invert"
curl -F 'file=@/path/matrix.csv' "localhost:8080/v1/multiply"
curl -F 'file=@/path/matrix.csv' "localhost:8080/v1/flatten"
curl -F 'file=@/path/matrix.csv' "localhost:8080/v1/sum"
```

## Swagger Documentation
```
http://localhost:8080/swagger/index.html#/
```
