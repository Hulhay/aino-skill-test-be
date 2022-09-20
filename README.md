# aino-skill-test-be

## Register

```
curl --location --request POST 'http://localhost:8080/api/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "hulhay",
    "email": "hulhay@gmail.com",
    "password": "123456"
}'
```

## Login
```
curl --location --request POST 'http://localhost:8080/api/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "hulhay@gmail.com",
    "password": "123456"
}'
```

## Logout
```
curl --location --request POST 'http://localhost:8080/api/auth/logout' \
--header 'Authorization: FILL_THE_TOKEN_HERE'
```