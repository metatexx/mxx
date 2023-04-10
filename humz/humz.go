package humz

import (
	"fmt"
	"math"
	"time"
)

func Bytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

func Count(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %c",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func Duration(b time.Duration) string {
	if b < 1000 {
		return fmt.Sprintf("%d ns", b)
	}
	us := float64(b) / 1000.0 // micro
	if us < 1000 {
		return fmt.Sprintf("%.3f µs", us)
	}
	ms := float64(us) / 1000.0 // milli
	if ms < 1000 {
		return fmt.Sprintf("%.3f ms", ms)
	}
	s := ms / 1000 // seconds
	if s < 60 {
		return fmt.Sprintf("%.3f s", s)
	}
	m := math.Floor(s / 60) // minutes
	if m < 60 {
		return fmt.Sprintf("%.0f m %.1f s", m, s-m*60)
	}
	h := math.Floor(m / 60) // stunden
	return fmt.Sprintf("%.0f h %.0f m %.0f s", h, m-h*60, s-m*60)
}
