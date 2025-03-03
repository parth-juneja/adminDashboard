# API Test Curls

- Health Check
```
curl --location 'http://localhost:8080/api/v1/health'
```

-GET ALL users
```
curl --location 'http://localhost:8080/api/v1/customers'
```

- GET user by UserId
```
curl --location 'http://localhost:8080/api/v1/customers/1'
```

- UPDATE customer details
```
curl -X PUT "http://localhost:8080/api/v1/customers/1" \
     -H "Content-Type: application/json" \
     -d '{
           "name": "John Doe Updated",
           "start_date": "2025-03-01T00:00:00Z",
           "remaining_days": 60,
           "subscription": "3_months",
           "payment_amount": 150.75,
           "last_paid": "2025-03-02T00:00:00Z"
         }'
```

- POST create User API test 
```
curl --location 'http://localhost:8080/api/v1/customers' \
--header 'Content-Type: application/json' \
--data '{
           "name": "John Doe",
           "start_date": "2025-03-01T00:00:00Z",
           "remaining_days": 30,
           "subscription": "1_month",
           "payment_amount": 100.50,
           "last_paid": "2025-03-01T00:00:00Z"
         }'
```

- Delete Customer by ID
```
curl --location --request DELETE 'http://localhost:8080/api/v1/customers/2'
```