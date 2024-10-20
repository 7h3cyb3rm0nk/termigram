# Termigram

Termigram is a Telegram bot that enables remote command execution and script management through a secure interface. It allows authorized users to execute system commands, manage directories, and run predefined scripts via Telegram messages.

## Features

- Remote command execution via Telegram
- Directory navigation and management
- Execution logging with timestamp and user tracking
- Predefined script management and execution
- Access control through user authorization
- Command history viewing

## Installation

1. Clone the repository:
```bash
git clone https://github.com/7h3cyb3rm0nk/termigram.git
cd termigram
```

2. Install dependencies:
```bash
go mod download
```

3. Configure the bot:
   - Copy `config.json.example` to `config.json`
   - Add your Telegram Bot API token
   - Add authorized user IDs
   - Configure scripts (optional)

Example configuration:
```json
{
  "bot_token": "YOUR_TELEGRAM_BOT_API_KEY",
  "allowed_users": [YOUR_TELEGRAM_USER_ID],
  "scripts": [
    {
      "id": "example_script",
      "comment": "print hello world",
      "command": "example.sh"
    }
  ]
}
```

## Usage

1. Start the bot:
```bash
go run main.go
```

2. Available commands in Telegram:
   - `/start` - Display help message and user ID
   - `/getlogs` - View command execution history
   - `/listscripts` - List all configured scripts
   - `/runscript script_name` - Execute a predefined script
   - Any other message will be treated as a system command

## Security Features

- User authorization through Telegram User IDs
- Command logging with timestamps
- Restricted access to authorized users only
- Working directory tracking

## Logs

All commands are logged to `bot_commands.log` with the following information:
- Timestamp
- User ID
- Executed command

## Contributing

Feel free to submit issues and enhancement requests.

## License

Copyright (C) 2024 Termigram

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

## Note

This bot provides system-level access through Telegram. Ensure proper security measures are in place and only authorize trusted users.
