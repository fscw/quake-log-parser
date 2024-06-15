package parser

import (
	"bufio"
	"fmt"
	"os"
	"quake_log_parser/pkg/models"
	"regexp"
	"strconv"
	"strings"
)

const worldKillerID = 1022

func meansOfDeath(modID int) string {
	if modName, ok := models.MeansOfDeath[modID]; ok {
		return modName
	}

	return models.MeansOfDeath[models.MOD_UNKNOWN]
}

func parsePlayer(line string) *models.Player {
	parts := strings.Split(line, " ")
	playerID, _ := strconv.Atoi(parts[2])
	infoParts := strings.Split(parts[3], "\\")
	playerName := infoParts[1]

	return &models.Player{ID: playerID, Name: playerName}
}

func parseKill(logLine string) *models.Kill {
	killDataRegex := regexp.MustCompile(`Kill:\s*(\d+)\s+(\d+)\s+(\d+):`)

	killData := killDataRegex.FindStringSubmatch(logLine)
	if len(killData) < 4 {
		fmt.Printf("Kill data not found on line: %s\n", logLine)
		return nil
	}

	killerID, err := strconv.Atoi(killData[1])
	if err != nil {
		fmt.Printf("Error converting killerID: %s\n", killData[1])
		return nil
	}

	victimID, err := strconv.Atoi(killData[2])
	if err != nil {
		fmt.Printf("Error converting victimID: %s\n", killData[2])
		return nil
	}

	modID, err := strconv.Atoi(killData[3])
	if err != nil {
		fmt.Printf("Error converting modID: %s\n", killData[3])
		return nil
	}

	return &models.Kill{KillerID: killerID, VictimID: victimID, ModID: modID}
}

func processKill(match *models.Match, kill *models.Kill) {
	match.TotalKills++
	modName := meansOfDeath(kill.ModID)

	if kill.KillerID == worldKillerID {
		match.Kills[kill.VictimID]--
	} else {
		match.Kills[kill.KillerID]++
	}

	match.KillsByMod[modName]++
}

func ParseLogFile(logFilePath string) ([]models.Match, error) {
	logFile, err := os.Open(logFilePath)

	if err != nil {
		return nil, err
	}
	defer logFile.Close()

	matches := []models.Match{}
	currentMatch := models.Match{
		Players:    make(map[int]*models.Player),
		Kills:      make(map[int]int),
		KillsByMod: make(map[string]int),
	}
	scanner := bufio.NewScanner(logFile)

	for scanner.Scan() {
		logLine := scanner.Text()
		if strings.Contains(logLine, "InitGame") {
			if currentMatch.TotalKills > 0 {
				matches = append(matches, currentMatch)
			}
			currentMatch = models.Match{
				ID:         len(matches) + 1,
				Players:    make(map[int]*models.Player),
				Kills:      make(map[int]int),
				KillsByMod: make(map[string]int),
			}
		} else if strings.Contains(logLine, "Kill") {
			kill := parseKill(logLine)
			processKill(&currentMatch, kill)
		} else if strings.Contains(logLine, "ClientUserinfoChanged") {
			player := parsePlayer(logLine)
			if player != nil {
				currentMatch.Players[player.ID] = player
			}
		}
	}

	if currentMatch.TotalKills > 0 {
		matches = append(matches, currentMatch)
	}

	return matches, scanner.Err()
}
