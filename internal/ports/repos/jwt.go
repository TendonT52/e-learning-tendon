package repos

import "time"

type JwtDB interface {
	InsertJwtToken(time.Time) (string, error)
	CheckJwtToken(string) error
	DeleteJwtToken(string) error
}
