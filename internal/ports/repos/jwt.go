package repos

import "time"

type JwtDB interface {
	InsertJwt(exp time.Time) (string, error)
	CheckJwt(hexId string) error
	DeleteJwt(hexId string) error
}
