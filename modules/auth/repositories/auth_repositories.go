package repositories

import (
	"fmt"
	"integration-auth-service/modules/auth/entities"
	"integration-auth-service/pkg/utils"

	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
)

type AuthRepository interface {
	GetAndInitCache(k string, dataKey string, getFromDB func() map[string]string) string
	GetClientSystemSourceByClientId(k string) string
	GetClientSecretByClientId(k string) string
}

type authRepo struct {
	Db *sqlx.DB
	C  *cache.Cache
}

func NewAuthRepository(db *sqlx.DB, c *cache.Cache) AuthRepository {
	return &authRepo{
		Db: db,
		C:  c,
	}
}

func (r *authRepo) GetAndInitCache(k string, dataKey string, getFromDB func() map[string]string) string {
	m, found := r.C.Get(dataKey)

	times := utils.GetTimeMinsToNewDay()
	if !found {
		m = getFromDB()
		r.C.Set(dataKey, m, times)
	}

	x, found := m.(map[string]string)[k]

	if !found {
		return k
	}

	return x
}

func (r *authRepo) GetClientSystemSourceByClientId(k string) string {
	x := r.GetAndInitCache(k, "clientSystemSource", func() map[string]string {
		xs := []entities.ClientData{}
		err := r.Db.Select(&xs, "SELECT * FROM client c Where c.is_active = true")

		if err != nil {
			fmt.Println("GetApplicationType Err: ", err)
		}

		m := make(map[string]string)

		for _, x := range xs {
			m[x.ClientID+x.GrantType] = x.SystemSource
		}
		return m
	})

	return x
}

func (r *authRepo) GetClientSecretByClientId(k string) string {
	x := r.GetAndInitCache(k, "clientSecret", func() map[string]string {
		xs := []entities.ClientData{}
		err := r.Db.Select(&xs, "SELECT * FROM client c Where c.is_active = true")
		if err != nil {
			fmt.Println("Get Client secret Err: ", err)
		}

		m := make(map[string]string)

		for _, x := range xs {
			m[x.ClientID+x.GrantType] = x.ClientSecret
		}
		return m
	})

	return x
}
