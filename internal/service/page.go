package service

import "strconv"

const defaultPageSize = 100

func fromPageToken(token string) (int, error) {
	if len(token) == 0 {
		return 0, nil
	}
	return strconv.Atoi(token)
}

func fromPageOffset(offset int) string {
	return strconv.Itoa(offset)
}
