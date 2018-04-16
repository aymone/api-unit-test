package storage

import (
	"github.com/aymone/api-unit-test/domain"
	mgo "gopkg.in/mgo.v2"
)

const collectionUser string = "users"

type UserStorage struct {
	session        *mgo.Session
	databaseName   string
	collectionName string
}

func NewUserStorage(session *mgo.Session) (*UserStorage, error) {
	return &UserStorage{
		session:        session,
		databaseName:   DBName,
		collectionName: collectionUser,
	}, nil
}

func (u *UserStorage) Insert(user *domain.User) (*domain.User, error) {
	session := u.session.Copy()
	defer session.Close()

	if err := session.DB(u.databaseName).C(u.collectionName).Insert(user); err != nil {
		if mgo.IsDup(err) {
			return nil, domain.ErrDuplicatedKey
		}
		return nil, err
	}

	return user, nil
}

func (u *UserStorage) Get(id string) (*domain.User, error) {
	session := u.session.Copy()
	defer session.Close()

	user := &domain.User{}
	if err := session.DB(u.databaseName).C(u.collectionName).FindId(id).One(user); err != nil {
		return nil, err
	}

	return user, nil
}
