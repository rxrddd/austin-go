package cls

var ClsAccount = struct {
	ID        string
	CreatedBy string
	CreatedAt string
	UpdatedBy string
	UpdatedAt string
	Username  string
	Password  string
	Salt      string
	RoleID    string
	Nickname  string
	IsDelete  string
}{
	ID:        "id",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
	Username:  "username",
	Password:  "password",
	Salt:      "salt",
	RoleID:    "role_id",
	Nickname:  "nickname",
	IsDelete:  "is_delete",
}
