package repository

import (
	"database/sql"

	"github.com/fajarihsan21/blog-be/database"
	"github.com/fajarihsan21/blog-be/model"
)

type usrRepository struct {
	db *database.Postgres
}

type UserRepository interface{
	GetUserByUsername(username string) (*model.User, error)
}


func (r usrRepository) GetUserByUsername(username string) (*model.User, error)  {
	result := &model.User{}
	query := `SELECT tu.id, tu.username, tu."password", tup."name", tup."role", tup.email, tup.phone, tu.is_active, tu.created_at, tu.updated_at FROM public.tb_users tu 
	LEFT JOIN tb_users_profile tup on tup.user_id = tu.id`
	err := r.db.GetActiveDB().Get(result, query+" WHERE tu.username = $1", username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}