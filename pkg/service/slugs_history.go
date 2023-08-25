package service

import (
	"fmt"
	"strings"
	"time"
)

const header = "slug_name,type,date\n"

func getType(removed bool) string {
	if removed {
		return "slug_removed"
	} else {
		return "slug_added"
	}
}

func (s *Service) GetSlugHistoryCsv(userId int, startDate time.Time, endDate time.Time) (string, error) {
	history, err := s.repository.GetSlugHistory(userId, startDate, endDate)
	if err != nil {
		return "", err
	}
	csv := strings.Builder{}
	csv.WriteString(header)
	for _, entry := range history {
		csv.WriteString(fmt.Sprintf(
			"%s,%s,%s\n",
			entry.SlugName,
			getType(entry.Removed),
			entry.DoneAt.Format(time.DateTime),
		))
	}
	return csv.String(), nil
}
