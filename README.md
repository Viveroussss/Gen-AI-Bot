# Telegram Bot with Message Reversal and Story Generation

This is a Telegram bot written in Go that provides two main functionalities:
1. Reverses any incoming message (e.g., "hello" becomes "olleh")
2. Generates short Sci-fi stories (up to 400 characters) using OpenAI's API

## Features

- Message reversal functionality
- Story generation using OpenAI API
- Docker support for easy deployment
- Unit tests for core functionality

## Prerequisites

- Go 1.16 or higher
- Docker and Docker Compose
- OpenAI API key

## Environment Variables

Create a `.env` file in the root directory with the following variables:
```
TELEGRAM_BOT_TOKEN=your_telegram_bot_token
OPENAI_API_KEY=your_openai_api_key
```

## Running the Bot

### Using Docker (Recommended)

1. Make sure you have Docker and Docker Compose installed
2. Create the `.env` file with your credentials
3. Run the bot:
```bash
docker-compose up -d
```

### Manual Setup

1. Install dependencies:
```bash
go mod download
```

2. Run the bot:
```bash
go run main.go
```

## Testing

Run the tests with:
```bash
go test ./...
```

## Project Structure

```
.
├── main.go              # Main application entry point
├── bot/
│   └── bot.go          # Bot implementation
├── utils/
│   └── string.go       # String manipulation utilities
├── tests/
│   └── string_test.go  # Unit tests
├── Dockerfile
├── docker-compose.yml
├── .env
└── README.md
```

## API Endpoints

The bot responds to the following commands:
- Any text message: Returns the reversed version of the message
- `/story`: Generates and returns a short Sci-fi story

## License

MIT