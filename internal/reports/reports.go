package reports

import (
	"fmt"
	"quake_log_parser/pkg/models"
)

func GenerateReport(match models.Match) string {
	report := fmt.Sprintf("Match %d:\n", match.ID)
	report += fmt.Sprintf("Total Kills: %d\n", match.TotalKills)
	report += "Players:"

	for _, player := range match.Players {
		report += player.Name + ", "
	}

	report += "\n Kills:\n"

	for id, kills := range match.Kills {
		if player, ok := match.Players[id]; ok {
			report += fmt.Sprintf("%s: %d\n", player.Name, kills)
		}
	}

	report += "Kills by Means of Death:\n"

	for mod, count := range match.KillsByMod {
		report += fmt.Sprintf("%s: %d\n", mod, count)
	}

	report += "\n"

	return report
}
