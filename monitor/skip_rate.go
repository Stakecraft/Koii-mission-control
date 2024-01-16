package monitor

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Stakecraft/koii-mission-control/alerter"
	"github.com/Stakecraft/koii-mission-control/config"
	"github.com/Stakecraft/koii-mission-control/types"
	"github.com/Stakecraft/koii-mission-control/utils"
)

var (
	solanaBinaryPath = os.Getenv("KOII_BINARY_PATH")
)

func SkipRate(cfg *config.Config) (float64, float64, error) {
	var valSkipped, netSkipped, totalSkipped float64

	if solanaBinaryPath == "" {
		solanaBinaryPath = "solana"
	}

	log.Printf("Koii binary path : %s", solanaBinaryPath)

	cmd := exec.Command(solanaBinaryPath, "validators", "--output", "json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while running Koii validators cli command %v", err)
		return valSkipped, netSkipped, err
	}

	var result types.SkipRate
	err = json.Unmarshal(out, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return valSkipped, netSkipped, err
	}

	for _, val := range result.Validators {
		if val.IdentityPubkey == cfg.ValDetails.PubKey {
			valSkipped = val.SkipRate
		}
		totalSkipped = totalSkipped + val.SkipRate
	}

	voteAccounts, err := GetVoteAccounts(cfg, utils.Network)
	if err != nil {
		log.Printf("Error while getting vote accounts : %v", err)
	}

	if &voteAccounts.Result != nil {
		currentVal := len(voteAccounts.Result.Current)

		netSkipped = totalSkipped / float64(currentVal)
	}

	log.Printf("VAL skip rate : %f, Network skip rate : %f", valSkipped, netSkipped)

	return valSkipped, netSkipped, nil
}

func SkipRateAlerts(cfg *config.Config) error {
	valSkipped, netSkipped, err := SkipRate(cfg)
	if err != nil {
		log.Printf("Error while getting SkipRate: %v", err)
	}

	if valSkipped > netSkipped && (valSkipped > float64(cfg.AlertingThresholds.SkipRateThreshold)) {
		if strings.EqualFold(cfg.AlerterPreferences.SkipRateAlerts, "yes") {
			err = alerter.SendTelegramAlert(fmt.Sprintf("SKIP RATE ALERT ::  Your validator SKIP RATE : %f has exceeded network SKIP RATE : %f", valSkipped, netSkipped), cfg)
			if err != nil {
				log.Printf("Error while sending skip rate alert to telegram: %v", err)
			}
			err = alerter.SendEmailAlert(fmt.Sprintf("Your validator SKIP RATE : %f has exceeded network SKIP RATE : %f", valSkipped, netSkipped), cfg)
			if err != nil {
				log.Printf("Error while sending skip rate alert to email: %v", err)
			}
		}
	}
	return nil
}
