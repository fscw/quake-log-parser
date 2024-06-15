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

   Ensure you have a Quake 3 Arena server log file you want to parse.

2. **Run the parser:**

   Modify the `main.go` file or create a new file to call the `ParseLogFile` function.

   ```go
   package main

   import (
       "fmt"
       "quake_log_parser/internal/app/parser"
   )

   func main() {
       logFilePath := "path/to/your/logfile.log"
       matches, err := parser.ParseLogFile(logFilePath)
       if err != nil {
           fmt.Printf("Error parsing log file: %s\n", err)
           return
       }

       for _, match := range matches {
           fmt.Printf("Match %d:\n", match.ID)
           fmt.Printf("Total Kills: %d\n", match.TotalKills)
           fmt.Println("Players:")
           for _, player := range match.Players {
               fmt.Printf(" - %s (ID: %d)\n", player.Name, player.ID)
           }
           fmt.Println("Kills by Means of Death:")
           for mod, count := range match.KillsByMod {
               fmt.Printf(" - %s: %d\n", mod, count)
           }
           fmt.Println()
       }
   }
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
