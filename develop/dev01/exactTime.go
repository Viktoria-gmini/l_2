package exactTime

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func ExactTime() {
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting NTP time: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Текущее точное время: %v\n", ntpTime.Format(time.RFC3339))
}
