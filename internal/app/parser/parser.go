package parser

import (
	"bufio"
	"os"
	"quake_log_parser/pkg/models"
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

func parseKill(line string) models.Kill {
	parts := strings.Split(line, " ")
	killerID, _ := strconv.Atoi(parts[3])
	victimID, _ := strconv.Atoi(parts[4])
	modID, _ := strconv.Atoi(parts[5])

	return models.Kill{KillerID: killerID, VictimID: victimID, ModID: modID}
}

func processKill(match *models.Match, kill models.Kill) {
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
		line := scanner.Text()
		if strings.Contains(line, "InitGame") {
			if currentMatch.TotalKills > 0 {
				matches = append(matches, currentMatch)
			}
			currentMatch = models.Match{
				ID:         len(matches) + 1,
				Players:    make(map[int]*models.Player),
				Kills:      make(map[int]int),
				KillsByMod: make(map[string]int),
			}
		} else if strings.Contains(line, "Kill") {
			kill := parseKill(line)
			processKill(&currentMatch, kill)
		} else if strings.Contains(line, "ClientUserinfoChanged") {
			player := parsePlayer(line)
			currentMatch.Players[player.ID] = player
		}
	}

	if currentMatch.TotalKills > 0 {
		matches = append(matches, currentMatch)
	}

	return matches, scanner.Err()
}
