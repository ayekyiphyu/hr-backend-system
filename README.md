# Product Structure


project
├── cmd                    # Command-related files
│   └── app                # Application entry point
│       └── main.go        # Main application logic
├── internal               # Internal codebase
│   ├── user               # Domain 'user'
│   │   ├── handler.go     # User-specific handler
│   │   ├── service.go     # User-specific service
│   │   ├── repository.go  # User-specific repository
│   │   └── user.go        # User model
│   └── product            # Domain 'product'
│       ├── handler.go     # Product-specific handler
│       ├── service.go     # Product-specific service
│       └── repository.go  # Product-specific repository
├── pkg/                   # Shared utilities or helpers
│   └── logger.go          # Logging utilities
├── configs/               # Configuration files (YAML, JSON, etc.)
│   └── config.yaml
├── go.mod                 # Go module definition
└── go.sum                 # Go module checksum file
