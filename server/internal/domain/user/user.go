package user

import (
	"github.com/nnniyaz/blog/internal/domain/base/email"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/internal/domain/user/valueobject"
	"time"
)

type User struct {
	id        uuid.UUID
	email     email.Email
	password  valueobject.Password
	role      valueobject.Role
	isDeleted bool
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(rawEmail, password, role string) (*User, error) {
	convertedEmail, err := email.NewEmail(rawEmail)
	if err != nil {
		return nil, err
	}

	convertedPassword, err := valueobject.NewPassword(password)
	if err != nil {
		return nil, err
	}

	convertedRole, err := valueobject.NewRole(role)
	if err != nil {
		return nil, err
	}

	return &User{
		id:        uuid.NewUUID(),
		email:     convertedEmail,
		password:  convertedPassword,
		role:      convertedRole,
		isDeleted: false,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func (u *User) GetId() uuid.UUID {
	return u.id
}

func (u *User) GetEmail() email.Email {
	return u.email
}

func (u *User) GetPassword() valueobject.Password {
	return u.password
}

func (u *User) GetRole() valueobject.Role {
	return u.role
}

func (u *User) GetIsDeleted() bool {
	return u.isDeleted
}

func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) UpdateEmail(rawEmail string) error {
	convertedEmail, err := email.NewEmail(rawEmail)
	if err != nil {
		return err
	}
	u.email = convertedEmail
	u.updatedAt = time.Now()
	return nil
}

func (u *User) UpdatePassword(password string) error {
	convertedPassword, err := valueobject.NewPassword(password)
	if err != nil {
		return err
	}

	u.password = convertedPassword
	u.updatedAt = time.Now()

	return nil
}

func (u *User) UpdateRole(role string) error {
	convertedRole, err := valueobject.NewRole(role)
	if err != nil {
		return err
	}

	u.role = convertedRole
	u.updatedAt = time.Now()

	return nil
}

func UnmarshalUserFromDatabase(id uuid.UUID, rawEmail string, password valueobject.Password, role string, isDeleted bool, createdAt, updatedAt time.Time) *User {
	return &User{
		id:        id,
		email:     email.Email(rawEmail),
		password:  password,
		role:      valueobject.Role(role),
		isDeleted: isDeleted,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
