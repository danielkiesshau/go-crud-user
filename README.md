## Testing Funcionalities

### Sign Up

````bash
curl -s --location --request POST 'http://localhost:8080/users/login' \
 --data-raw '{
 "Username":"daniel",
 "Email":"daniel@abc.com",
 "Password":"daniel",
 "User_type":"ADMIN"
 }' ```

### Log in

```bash
curl -s --location --request POST 'http://localhost:8080/users/login' \
 --data-raw '{
 "Email":"daniel@abc.com",
 "Password":"abc"
 }'
````