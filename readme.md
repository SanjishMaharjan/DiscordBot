# Discord Bot with Trivia Functionality

This project is a simple Discord bot implemented in Go that includes trivia functionality. The bot responds to commands and provides trivia questions with answers after a delay.

## Project Structure

discord-bot/
│
├── main.go        # Main bot logic
├── handlers/
│   └── trivia.go  # Trivia logic
└── go.mod         # Go module file
└── .env           # Environment variables


## Prerequisites

- [Go](https://golang.org/doc/install) installed on your machine.
- A Discord bot token. You can obtain one by creating a bot on the [Discord Developer Portal](https://discord.com/developers/applications).

## Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/SanjishMaharjan/DiscordBot.git
   cd discord-bot
   ```

2. **Install Dependencies**

   ```bash
   go mod tidy
   ```

3. **Create a `.env` File**

   ```bash
   TOKEN=<your-discord-bot-token>
   ```

4. **Run the Bot**

   ```bash
   go run main.go
   ```

## Usage

- **Commands**

  - `/trivia`: Fetches a trivia question and answers.
  - `/help`: Displays a list of available commands.
  
- **Example**

  ```bash
  !trivia
  ```




