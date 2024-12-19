package repository

import (
	"database/sql"

	"github.com/fajarihsan21/blog-be/database"
	"github.com/fajarihsan21/blog-be/model"
	"github.com/google/uuid"
)

type usrRepository struct {
	db *database.Postgres
}

type UserRepository interface{
	GetUserByUsername(username string) (*model.User, error)
	InsertUser(data model.User) (uuid.UUID, error)
	InsertUserProfile(data model.UserProfile) (uuid.UUID, error)
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

func (r usrRepository) InsertUser(data model.User) (uuid.UUID, error) {
	query := `INSERT INTO public.tb_users (id, username, "password", is_active, created_at, updated_at) 
			VALUES(:id, :username, :password, :is_active, :created_at, :updated_at)
			RETURNING id;`

	stmt, err := r.db.GetActiveDB().PrepareNamed(query)
	if err != nil {
		return uuid.Nil, err
	}

	var id uuid.UUID
    err = stmt.Get(&id, data)
    if err != nil {
        return uuid.Nil, err
    }

    return id, nil
}

func (r usrRepository) InsertUserProfile(data model.UserProfile) (uuid.UUID, error) {
	query := `INSERT INTO public.tb_users_profile
			(id, user_id, "name", "role", email, phone, is_active, created_at, updated_at)
			VALUES(:id, :user_id, :name, :role, :email, :phone, :is_active, :created_at, :updated_at);`
	
	stmt, err := r.db.GetActiveDB().PrepareNamed(query)
	if err != nil {
		return uuid.Nil, err
	}

	var profileId uuid.UUID
    err = stmt.Get(&profileId, data)
    if err != nil {
        return uuid.Nil, err
    }

    return profileId, nil
}