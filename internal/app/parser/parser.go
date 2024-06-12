package parser

import (
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

func ParseLogFile() {
	return
}
