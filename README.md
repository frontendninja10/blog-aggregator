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

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You can then install `gator` with:

```bash
go install github.com/frontendninja10/blog-aggregator/cmd/gator@latest
```

## Configuration

Set up your PostgreSQL database and create a configuration file:
```bash
# Create configuration file in your home directory
echo '{"db_url": "postgres://username:password@localhost/dbname?sslmode=disable"}' > ~/.gatorconfig.json
```

Replace `username`, `password`, and `dbname` with your PostgreSQL credentials and database name.

## Usage

The application provides several commands for managing users and RSS feeds:

Create a new user:
```bash
gator register <username>
```

Add a feed:
```bash
gator addfeed <feed_name> <feed_url>
```

Start aggregating feeds:
```bash
gator aggregate <time_between_requests>
```
Time between requests is in the format of `1s`, `1m`, `1h`, etc.

View the posts:
```bash
gator browse [limit]
```

Other commands you might need:
- `gator login <username>`: Login as an existing user
- `gator users`: List all users
- `gator reset`: Reset the database (deletes all users)
- `gator feeds`: List all feeds
- `gator follow <feed_url>`: Follow an existing feed
- `gator unfollow <feed_url>`: Unfollow an existing feed
- `gator following`: List all followed feeds

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

3. **Configuration file not found**: Make sure `~/.gatorconfig.json` exists in your home directory with valid database credentials.

4. **RSS feed parsing errors**: Some feeds may have non-standard formats. The application includes HTML unescaping for better compatibility.
