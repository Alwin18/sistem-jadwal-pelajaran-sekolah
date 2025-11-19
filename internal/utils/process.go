package utils

import (
	"time"

	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/models"
)

const JPDuration = 40 // menit per JP

func HitungTotalJP(scheduleList []models.Schedule) int {
	totalJP := 0

	for _, s := range scheduleList {

		// Parse start
		start, err := time.Parse("15:04:05", s.TimeStart)
		if err != nil {
			// fallback jika hanya "HH:MM"
			start, err = time.Parse("15:04", s.TimeStart)
			if err != nil {
				continue
			}
		}

		// Parse end
		end, err := time.Parse("15:04:05", s.TimeEnd)
		if err != nil {
			// fallback jika hanya "HH:MM"
			end, err = time.Parse("15:04", s.TimeEnd)
			if err != nil {
				continue
			}
		}

		duration := end.Sub(start)
		minutes := int(duration.Minutes())

		if minutes <= 0 {
			continue
		}

		jp := minutes / JPDuration
		totalJP += jp
	}

	return totalJP
}
