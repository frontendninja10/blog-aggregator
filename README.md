# Blog Aggregator

A command-line RSS feed aggregator built in Go that allows users to manage RSS feeds and fetch blog posts from various sources.

## Features

- **User Management**: Register, login, and manage multiple users
- **RSS Feed Parsing**: Fetch and parse RSS feeds from any URL
- **Database Storage**: PostgreSQL database for persistent user data
- **Configuration Management**: JSON-based configuration with user sessions
- **CLI Interface**: Simple command-line interface for all operations

## Prerequisites

- Go 1.24.4 or later
- PostgreSQL database
- [SQLC](https://sqlc.dev/) for database code generation (development)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/frontendninja10/blog-aggregator.git
cd blog-aggregator
```

2. Install dependencies:
```bash
go mod download
```

3. Set up your PostgreSQL database and create a configuration file:
```bash
# Create configuration file in your home directory
echo '{"db_url": "postgres://username:password@localhost/dbname?sslmode=disable", "current_username": ""}' > ~/.gatorconfig.json
```

4. Run database migrations:
```bash
# Apply the schema from sql/schema/001_users.sql to your database
```

## Usage

The application provides several commands for managing users and fetching RSS feeds:

### User Management

#### Register a new user
```bash
go run . register <username>
```
Creates a new user and automatically logs them in.

#### Login as an existing user
```bash
go run . login <username>
```
Switch to an existing user account.

#### List all users
```bash
go run . users
```
Display all registered users, with the current user marked with an asterisk.

#### Reset database
```bash
go run . reset
```
⚠️ **Warning**: This deletes all users from the database.

### RSS Feed Operations

#### Fetch RSS feed
```bash
gator . agg <feed_url>
```
Fetch and display RSS feed content from the specified URL.

Example:
```bash
go run . agg https://blog.boot.dev/index.xml
```

## Project Structure

```
blog-aggregator/
├── main.go                 # Application entry point
├── commands.go             # Command registration and routing
├── register_handler.go     # User registration logic
├── login_handler.go        # User login logic
├── users_handler.go        # List users functionality
├── reset_handler.go        # Database reset functionality
├── agg.go                  # RSS feed aggregation
├── rss.go                  # RSS parsing utilities
├── internal/
│   ├── config/             # Configuration management
│   └── database/           # Generated database code (SQLC)
├── sql/
│   ├── queries/            # SQL queries for SQLC
│   └── schema/             # Database schema migrations
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
└── sqlc.yaml              # SQLC configuration
```

## Configuration

The application uses a JSON configuration file stored at `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://username:password@localhost/dbname?sslmode=disable",
  "current_username": "your_username"
}
```

- `db_url`: PostgreSQL connection string
- `current_username`: Currently logged-in user (managed automatically)

## Database Schema

The application uses a simple user table:

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT UNIQUE NOT NULL
);
```

## Development

### Database Code Generation

This project uses [SQLC](https://sqlc.dev/) to generate type-safe Go code from SQL queries:

```bash
# Install SQLC
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Generate database code
sqlc generate
```

### Adding New Commands

1. Create a handler function with the signature: `func(s *state, cmd command) error`
2. Register the command in `main.go`:
```go
cmds.register("your_command", yourHandler)
```

## Dependencies

- [github.com/lib/pq](https://github.com/lib/pq) - PostgreSQL driver
- [github.com/google/uuid](https://github.com/google/uuid) - UUID generation

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source. Please check the repository for license details.

## Troubleshooting

### Common Issues

1. **Database connection errors**: Ensure PostgreSQL is running and the connection string in your config file is correct.

2. **User already exists**: The register command will fail if you try to create a user that already exists.

3. **Configuration file not found**: Make sure `~/.gatorconfig.json` exists with valid database credentials.

4. **RSS feed parsing errors**: Some feeds may have non-standard formats. The application includes HTML unescaping for better compatibility.

## Examples

```bash
# Complete workflow example
go run . register alice          # Register user 'alice'
go run . users                   # List users (alice will be marked as current)
go run . agg https://example.com/feed.xml  # Fetch RSS feed
go run . register bob            # Register another user 'bob'
go run . login alice             # Switch back to alice
go run . users                   # List users (alice marked as current)
```
