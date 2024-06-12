## Golang Service Sample

A Web application sample implemented with Go/Fiber.

## Project Structure

```bash
## project structure
.
├── README.md
├── go.mod            // defines your module and its dependencies
├── go.sum            // ensures the integrity and reproducibility of your dependencies
├── handlers
│   ├── about.go      // "/about" about page
│   ├── home.go       // "/" home page
│   └── home_test.go  // home unit test - go test ./...
├── main.go           // entry point file
└── routes
└── routes.go         // application routes
```

## Unit Test

testify: https://github.com/stretchr/testify
