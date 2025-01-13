package uuid

import (
	"github.com/google/uuid"
	"github.com/nnniyaz/blog/pkg/core"
)

var Nil = UUID(uuid.Nil)

type UUID uuid.UUID

func NewUUID() UUID {
	return UUID(uuid.New())
}

func (id UUID) String() string {
	return uuid.UUID(id).String()
}

func UUIDFromString(id string) (UUID, error) {
	parsedUUID, err := uuid.Parse(id)
	if uuid.IsInvalidLengthError(err) {
		return Nil, core.NewI18NError(core.EINVALID, core.TXT_UUID_INVALID_LENGTH)
	}
	if err != nil {
		return Nil, err
	}
	return UUID(parsedUUID), nil
}

func UUIDFromBytes(id []byte) (UUID, error) {
	parsedUUID, err := uuid.FromBytes(id)
	if err != nil {
		return Nil, err
	}
	return UUID(parsedUUID), nil
}
