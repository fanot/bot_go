# Telegram Bot

This is a Telegram bot built using the Go programming language and the `go-telegram-bot-api` library. The bot is designed to interact with users, allowing them to save and retrieve personal information.

## Features

- Start a conversation with the bot.
- Add facts, location, and bio.
- Retrieve saved information.

## Prerequisites

- Go 1.22.2 or later
- Docker (optional, for containerized deployment)
- A Telegram bot token

## Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/telegram-bot.git
   cd telegram-bot
   ```

2. **Set the Telegram token:**

   Create a `.env` file in the root directory and add your Telegram bot token:

   ```env
   TELEGRAM_TOKEN=your-telegram-bot-token
   ```

3. **Build and run the bot:**

   You can run the bot directly or use Docker.

   **Directly:**

   ```bash
   go build -o bot
   ./bot
   ```

   **Using Docker:**

   Build and run the Docker container:

   ```bash
   docker-compose up --build
   ```

## Testing

To run the tests, use the following command:

```bash
cd tests
go test
```

## Project Structure

- `main.go`: Entry point of the application.
- `bot/`: Contains the bot logic and handlers.
- `tests/`: Contains test files for the bot.

## Contributing

Feel free to submit issues or pull requests if you find any bugs or have suggestions for improvements.

## License

This project is licensed under the MIT License. 
