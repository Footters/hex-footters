package redisdb

import (
	"encoding/json"

	"github.com/Footters/hex-footters/pkg/auth"
	"github.com/go-redis/redis"
)

const table = "users"

type userRepository struct {
	con *redis.Client
}

// NewRedisUserRepository Constructor
func NewRedisUserRepository(connection *redis.Client) auth.UserRepository {
	return &userRepository{
		con: connection,
	}
}

func (r *userRepository) Create(user *auth.User) error {
	encoded, err := json.Marshal(user)
	if err != nil {
		return err
	}

	r.con.HSet(table, user.Email, encoded)

	return nil
}

func (r *userRepository) FindByEmail(email string) (*auth.User, error) {

	b, err := r.con.HGet(table, email).Bytes()
	if err != nil {
		return nil, err
	}

	u := new(auth.User)
	err = json.Unmarshal(b, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
