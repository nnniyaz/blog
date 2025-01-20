package valueobject

import "github.com/nnniyaz/blog/internal/domain/user/exceptions"

type Role string

const (
	RoleAdmin     Role = "admin"
	RoleModerator Role = "moderator"
)

func NewRole(role string) (Role, error) {
	switch role {
	case "admin":
		return RoleAdmin, nil
	case "moderator":
		return RoleModerator, nil
	default:
		return "", exceptions.ErrUserRoleInvalid
	}
}

func (r Role) String() string {
	return string(r)
}
