# Quake Log Parser

This Go application parses Quake 3 Arena server log files to extract match information, including players and kill events.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Example](#example)
- [License](#license)

## Features

- Parse Quake 3 Arena log files
- Extract player information
- Extract and process kill events
- Aggregate match statistics

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/quake_log_parser.git
   cd quake_log_parser
   ```

2. **Install dependencies:**

   Ensure you have Go installed. If not, download and install it from [here](https://golang.org/dl/).

   ```bash
   go mod tidy
   ```

## Usage

1. **Prepare your log file:**

   Ensure you have a Quake 3 Arena server log file you want to parse. Access folder `logs`, open file `qgames.log` then paste the log you want to be analyzed.

2. **Run the parser:**

   On `cmd/quakeparser` folder, run:

   ```bash
   go run main.go
   ```

3. **Build and run the application:**

   ```bash
   go build -o quakeparser
   ./quakeparser
   ```

## Example

Below is an example of how the output might look:

```
Match 1:
Total Kills: 11
Players:
 - Player1 (ID: 1)
 - Player2 (ID: 2)
Kills:
 - Player1: 5
 - Player2: 6
Kills by Means of Death:
 - MOD_UNKNOWN: 11
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
