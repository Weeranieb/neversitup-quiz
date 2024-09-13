# Folder Structure Conventions

> This is my latest projects directory

### File Directory

    .
    ├── controllers                   # Handles incoming requests and returns responses
    ├── middlewares                   # Middleware functions for request processing
    ├── migrations                    # Database schema migrations
    ├── pkg                           # combine all service, repo, etc.
        ├── models                    # Data models for the application
        ├── processors                # Business logic that access more than 1 domain
        ├── repositories              # Data access layer
        ├── services                  # Business logic and services
    ├── utils                     # Utility functions and helpers
    ├── main.go
    └── README.md
