package jobs

import (
	"fmt"
	"mtui/db"
	"time"
)

func logCleanup(r *db.LogRepository) {
	for {
		ts := time.Now().AddDate(0, 0, -30)
		err := r.DeleteBefore(ts.UnixMilli())
		if err != nil {
			fmt.Printf("Log cleanup error: %s\n", err.Error())
		}

		// re-schedule every minute
		time.Sleep(time.Minute)
	}
}
