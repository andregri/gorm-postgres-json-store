# RESTful API using Go, Gorilla Mux, GORM, PostgreSQL

```
curl -X POST \http://localhost:8000/v1/user \
    -H 'cache-control: no-cache' \
    -H 'content-type: application/json' \
    -d '{ \
        "username": "andrea", \
        "email_address": "a.b@c.com", \
        "first_name": "Andrea", \
        "last_name": "Grillo" \
    }'
```