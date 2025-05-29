package interfaces

import xtt "github.com/faelmori/xtui/types"

type IUserRepo interface {
	Create(u User) (User, error)
	FindOne(where ...interface{}) (User, error)
	FindAll(where ...interface{}) ([]User, error)
	Update(u User) (User, error)
	Delete(id string) error
	Close() error
	List(where ...interface{}) (xtt.TableHandler, error)
}
