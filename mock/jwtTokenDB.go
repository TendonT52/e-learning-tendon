package mock

import (
	"time"
)

type MockJwtDB struct {
	Insert error
	Check  error
	Delete error
}

func (jw *MockJwtDB) InsertJwtToken(exp time.Time) (string, error) {
	switch jw.Insert {
	case nil:
		return "mock hex id", nil
	default:
		return "", jw.Insert
	}
}

func (jw *MockJwtDB) CheckJwtToken(hexId string) error {
	switch jw.Check {
	case nil:
		return nil
	default:
		return jw.Check
	}
}

func (jw *MockJwtDB) DeleteJwtToken(hexId string) error {
	switch jw.Delete {
	case nil:
		return nil
	default:
		return jw.Delete
	}
}
