package db

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"mtui/types"
	"time"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/google/uuid"
	"github.com/minetest-go/dbutil"
)

type OauthAppRepository struct {
	DB dbutil.DBTx
}

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go/22892986#22892986
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (r *OauthAppRepository) Set(m *types.OauthApp) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.Secret == "" {
		m.Secret = randSeq(16)
	}
	if m.Created == 0 {
		m.Created = time.Now().Unix()
	}
	return dbutil.InsertOrReplace(r.DB, m)
}

func (r *OauthAppRepository) GetAll() ([]*types.OauthApp, error) {
	return dbutil.SelectMulti(r.DB, func() *types.OauthApp { return &types.OauthApp{} }, "")
}

func (r *OauthAppRepository) GetByName(name string) (*types.OauthApp, error) {
	f, err := dbutil.Select(r.DB, &types.OauthApp{}, "where name = $1", name)
	if err == sql.ErrNoRows {
		return nil, nil
	} else {
		return f, err
	}
}

func (r *OauthAppRepository) GetByID(id string) (*types.OauthApp, error) {
	f, err := dbutil.Select(r.DB, &types.OauthApp{}, "where id = $1", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else {
		return f, err
	}
}

func (r *OauthAppRepository) Delete(id string) error {
	return dbutil.Delete(r.DB, &types.OauthApp{}, "where id = $1", id)
}

type OAuthAppStore struct {
	Repo *OauthAppRepository
}

func (s *OAuthAppStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	// id == name
	app, err := s.Repo.GetByName(id)
	if err != nil {
		return nil, err
	} else if !app.Enabled {
		return nil, errors.New("app not enabled")
	} else {
		return app, nil
	}
}
