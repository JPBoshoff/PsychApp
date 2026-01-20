# Copilot Instructions for PsychApp

## Project Overview

PsychApp is a Go-based API service using chi router, Viper for configuration, and Zap for structured logging.

## Code Style & Conventions

### Go Best Practices

- Follow standard Go formatting (use `gofmt` and `goimports`)
- Use Go 1.25+ features
- Prefer `context.Context` for cancellation and timeouts
- Always handle errors explicitly - never ignore them
- Use structured logging with `zap.Logger`

### Naming Conventions

- **Packages**: lowercase, single word (e.g., `entries`, `health`, `config`)
- **Files**: lowercase with underscores if needed (e.g., `handler.go`, `router.go`)
- **Variables**: camelCase (e.g., `userID`, `configPath`)
- **Constants**: camelCase or UPPER_SNAKE_CASE for exported constants
- **Interfaces**: suffix with `-er` when appropriate (e.g., `Handler`, `Runner`)

### Project Structure

Follow the existing structure:

```
services/api/
├── cmd/api/          # Application entry points
├── internal/         # Private application code
│   ├── app/         # Application setup and lifecycle
│   ├── config/      # Configuration management
│   ├── entries/     # Domain handlers
│   ├── health/      # Health check endpoints
```

## Architecture Patterns

### Handler Pattern

- Handlers should be in their own package under `internal/`
- Each handler should have a `handler.go` file
- HTTP handlers should accept `http.ResponseWriter` and `*http.Request`
- Use dependency injection for logger, config, and services

### Configuration

- Use Viper for configuration management
- Support environment variables
- Configuration should be loaded in `config.Load()`
- Validate configuration on startup

### Logging

- Use `zap.Logger` for all logging
- Use structured logging with fields (e.g., `zap.String("key", "value")`)
- Log levels: Debug, Info, Warn, Error, Fatal
- Include context in logs (request ID, user ID, etc.)

### Error Handling

- Return errors up the call stack
- Wrap errors with context using `fmt.Errorf("context: %w", err)`
- Log errors at the appropriate level
- Return appropriate HTTP status codes

## API Development

### Router (chi)

- Define routes in `router.go`
- Group related routes with `chi.Router.Route()`
- Use middleware for common concerns (logging, auth, CORS)
- Keep handler logic separate from routing

### HTTP Responses

- Use standard HTTP status codes
- Return JSON for API responses
- Handle errors gracefully with appropriate status codes
- Include proper headers (Content-Type, etc.)

### Graceful Shutdown

- Use `context` for graceful shutdown
- Listen for `os.Interrupt` and `syscall.SIGTERM`
- Clean up resources (connections, files) on shutdown

## Dependencies

Core libraries in use:

- **chi/v5**: HTTP router
- **viper**: Configuration management
- **zap**: Structured logging

## Testing

- Write tests for all handlers
- Use table-driven tests where appropriate
- Mock external dependencies
- Test files should be named `*_test.go`
- Use `t.Parallel()` for independent tests

## Security

- Validate all user input
- Use environment variables for sensitive data
- Never commit secrets or credentials
- Use HTTPS in production

## Code Generation Guidelines

When generating new code:

1. Follow the existing package structure
2. Include proper error handling
3. Add structured logging
4. Use dependency injection
5. Include docstrings for exported functions
6. Consider graceful shutdown implications

## Comments & Documentation

- Add docstrings for all exported types and functions
- Use `//` for single-line comments
- Keep comments concise and meaningful
- Update comments when code changes
