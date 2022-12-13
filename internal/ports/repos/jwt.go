package repos

import "time"

type JwtDB interface {
	InsertJwtToken(exp time.Time) (string, error)
	CheckJwtToken(hexId string) error
	DeleteJwtToken(hexId string) error
}
