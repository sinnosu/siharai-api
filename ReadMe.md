# set on locally
 * build on devcontainer 
 * setup DB
   * create table
   ```
   $ make sqldef-setup
   $ make db-dry-run
   $ make db-apply
   ```
   * insert test data
   ```
   $ make db-connect
   > password: password
   ```
   - Please RUN insert SQL from ./db/testdata/clients.sql and companies.sql

## Start server on locally
GO_ENV=dev go run src/main.go

## Try API By Curl
 * signup
```
curl -X POST http://localhost:8080/signup -H "Content-Type: application/json" -H "X-CSRF-Token: your-csrf-token-here" -d '{"email": "example@example.com", "password": "yourpassword", "name": "your name", "company_id": 1}'
```

 * login
```
curl -X POST http://localhost:8080/login   -H "Content-Type: application/json"   -H "X-CSRF-Token: your-csrf-token-here"   -d '{"email": "example@example.com", "password": "yourpassword"}'  -c cookies.txt
```

 * create invoices
 ```
 curl -iX POST http://localhost:8080/api/invoices \
-H "Content-Type: application/json" -H "X-CSRF-Token: your-csrf-token-here" \
-d '{
    "invoices": [
        {
            "company_id": 1,
            "client_id": 1,
            "issue_date": "2024-01-01",
            "payment_amount": 10000,
            "payment_due_date": "2024-01-31"
        },
        {
            "company_id": 1,
            "client_id": 2,
            "issue_date": "2024-02-15",
            "payment_amount": 15000,
            "payment_due_date": "2024-03-31"
        }
    ]
}' -b cookies.txt
 ```

 * get invoices
 ```
 curl -iX GET "http://localhost:8080/api/invoices?fromDate=2024-01-01&toDate=2024-02-28" -b cookies.txt
 ```


## 残課題
- テストコードを書く。