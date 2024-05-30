# User CRUD

Simple Go Grud application to practice:

- Go Syntax
- Gin Framework
- JWT
- MongoDB
- Authorization Model

## To run the project
1. Run a MongoDb container
```bash
docker build -t mongo-image
docker run -d -p 27017:27017 --name mongo-container mongo-image
```
2. Run the Go Project
```bash
go run ./main.go
```

## Testing Funcionalities

### Curls
````bash
### SignUp
curl -s --location --request POST 'http://localhost:8080/users/signup' \
 --data-raw '{
 "Username":"daniel",
 "Name": "daniel",
 "Email":"daniel@abc.com",
 "Password":"daniel",
 "User_type":"ADMIN"
 }' ```

### Log in

```bash
curl -s --location --request POST 'http://localhost:8080/users/login' \
 --data-raw '{
 "Email":"daniel@abc.com",
 "Password":"daniel"
 }'
```

## Get user details (Only ADMIN is allowed)
```bash
curl -s --location --request GET "http://localhost:8080/users/6658c3c83087b9381e7cede1" \
--header "token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImRhbmllbEBhYmMuY29tIiwiTmFtZSI6ImRhbmllbCIsIlVzZXJuYW1lIjoiZGFuaWVsIiwiVWlkIjoiNjY1OGMzYzgzMDg3YjkzODFlN2NlZGUxIiwiVXNlcl90eXBlIjoiQURNSU4iLCJleHAiOjE3MTcxMzY1NjF9.hD8gvE1Xim0P4ETGRrwJdOt8QXWalv-Ajgoqsl82W34" \
--data-raw ''
```


## Get all users (only ADMIN is allowed)
curl -s --location --request GET 'http://localhost:8080/users' \
--header "token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImRhbmllbEBhYmMuY29tIiwiTmFtZSI6ImRhbmllbCIsIlVzZXJuYW1lIjoiZGFuaWVsIiwiVWlkIjoiNjY1OGMzYzgzMDg3YjkzODFlN2NlZGUxIiwiVXNlcl90eXBlIjoiQURNSU4iLCJleHAiOjE3MTcxMzY1NjF9.hD8gvE1Xim0P4ETGRrwJdOt8QXWalv-Ajgoqsl82W34" \
--data-raw ''
````
