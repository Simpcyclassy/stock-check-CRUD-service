``` go
r.HandleFunc("/create", handlers.CreateStockHandler) // CREATE
r.HandleFunc("/read", handlers.ReadStockHandler) // READ
r.HandleFunc("/update", handlers.UpdateStockHandler) // PUT
r.HandleFunc("/delete", handlers.DeleteStockHandler)


func ReadStockHandler(w http.ResponseWriter, r *http.Request) {
    //
}
```

OR we could design the program like:

```go
r.HandleFunc("/myapp", handlers.StockHandler)

func StockHandler(w http.ResponseWriter, r *http.Request) {
    switch r.method {
        case "PUT":
            // do some action

        case "CREATE":
            // do some action


        case "DELETE":
            // do some action

    }
}
```