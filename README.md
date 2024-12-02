# Fastest API Fetcher

This project is a Go-based CLI application designed to fetch address data from two APIs concurrently. It determines which API responds the fastest, displays the result, and includes a 1-second timeout limit. If neither API responds within the time limit, a timeout error is shown.

---

## Features

- **Concurrent API Requests**: Sends requests to both APIs simultaneously.
- **Timeout Management**: Limits the response time to 1 second.
- **Winner Selection**: Displays the result from the faster API and discards the slower response.
- **Command-line Output**: Outputs the address data and the source API in the terminal.

---

## APIs Used

1. [BrasilAPI](https://brasilapi.com.br/docs#tag/CEP/operation/cepV1Controller_findByCEP)
2. [ViaCEP](https://viacep.com.br/)

---

## Folder Structure
``` bash
project/
├── cmd/
│   └── main.go     
├── internal/
│   ├── apiclient/       
│   │   └── client.go
├── go.mod               
└── go.sum             
```


### Explanation of Folders

- **`cmd/`**  
  Contains the entry point for the application (`main.go`).

- **`internal/apiclient/`**  
  Includes logic for making HTTP requests, managing concurrency, and handling timeout limits.

- **`go.mod` and `go.sum`**  
  Define and manage the project's dependencies.

---

## Requirements

- Go version 1.20 or higher

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/math-schenatto/go-multithreading
   cd go-multithreading


## Execution

To execute the project, follow these steps:

1. **Run the application**:
   Navigate to the project directory and execute the following command:
   ```bash
   go run cmd/main.go
