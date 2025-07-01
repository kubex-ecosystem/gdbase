package user

import (
	"fmt"

	"github.com/google/uuid"
	is "github.com/rafa-mori/gdbase/internal/services"
	gl "github.com/rafa-mori/gdbase/logger"
	t "github.com/rafa-mori/gdbase/types"
	l "github.com/rafa-mori/logz"
	xtt "github.com/rafa-mori/xtui/types"
	"gorm.io/gorm"
)

type IUserRepo interface {
	// TableName returns the name of the table in the database.
	TableName() string
	Create(u IUser) (IUser, error)
	FindOne(where ...interface{}) (IUser, error)
	FindAll(where ...interface{}) ([]IUser, error)
	Update(u IUser) (IUser, error)
	Delete(id string) error
	Close() error
	List(where ...interface{}) (xtt.TableHandler, error)
	GetContextDbService() t.DBService
}

type UserRepo struct {
	// g is the gorm.DB instance used for database operations.
	g *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	if db == nil {
		gl.Log("error", "UserModel repository: gorm DB is nil")
		return nil
	}
	return &UserRepo{db}
}

func (ur *UserRepo) TableName() string {
	return "users"
}

func (ur *UserRepo) Create(um IUser) (IUser, error) {
	if um == nil {
		return nil, fmt.Errorf("UserModel repository: UserModel is nil")
	}
	if uModel := um.GetUserObj(); uModel == nil {
		return nil, fmt.Errorf("UserModel repository: UserModel is not of type *UserModel")
	} else {
		if uModel.GetID() != "" {
			if _, err := uuid.Parse(uModel.GetID()); err != nil {
				return nil, fmt.Errorf("UserModel repository: UserModel ID is not a valid UUID: %w", err)
			}
		} else {
			uModel.SetID(uuid.New().String())
		}

		err := ur.g.Create(uModel.GetUserObj()).Error
		if err != nil {
			return nil, fmt.Errorf("UserModel repository: failed to create UserModel: %w", err)
		}
		return uModel, nil
	}
}
func (ur *UserRepo) FindOne(where ...interface{}) (IUser, error) {
	var um UserModel
	err := ur.g.Where(where[0], where[1:]...).First(&um).Error
	if err != nil {
		return nil, fmt.Errorf("UserModel repository: failed to find UserModel: %w", err)
	}
	return &um, nil
}
func (ur *UserRepo) FindAll(where ...interface{}) ([]IUser, error) {
	var ums []UserModel
	err := ur.g.Where(where[0], where[1:]...).Find(&ums).Error
	if err != nil {
		return nil, fmt.Errorf("UserModel repository: failed to find all users: %w", err)
	}
	ius := make([]IUser, len(ums))
	for i, usr := range ums {
		ius[i] = &usr
	}
	return ius, nil
}
func (ur *UserRepo) Update(um IUser) (IUser, error) {
	if um == nil {
		return nil, fmt.Errorf("UserModel repository: UserModel is nil")
	}
	uModel := um.GetUserObj()
	if uModel == nil {
		return nil, fmt.Errorf("UserModel repository: UserModel is not of type *UserModel")
	}
	err := ur.g.Save(uModel).Error
	if err != nil {
		return nil, fmt.Errorf("UserModel repository: failed to update UserModel: %w", err)
	}
	return uModel, nil
}
func (ur *UserRepo) Delete(id string) error {
	err := ur.g.Delete(&UserModel{}, id).Error
	if err != nil {
		return fmt.Errorf("UserModel repository: failed to delete UserModel: %w", err)
	}
	return nil
}
func (ur *UserRepo) Close() error {
	sqlDB, err := ur.g.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
func (ur *UserRepo) List(where ...interface{}) (xtt.TableHandler, error) {
	var users []UserModel
	err := ur.g.Where(where[0], where[1:]...).Find(&users).Error
	if err != nil {
		return xtt.TableHandler{}, fmt.Errorf("UserModel repository: failed to list users: %w", err)
	}
	tableHandlerMap := make([][]string, 0)
	for i, usr := range users {
		tableHandlerMap = append(tableHandlerMap, []string{
			fmt.Sprintf("%d", i+1),
			usr.GetID(),
			usr.GetName(),
			usr.GetUsername(),
			usr.GetEmail(),
			usr.GetPhone(),
			fmt.Sprintf("%t", usr.GetActive()),
		})
	}
	return xtt.TableHandler{Rows: tableHandlerMap}, nil
}
func (ur *UserRepo) GetContextDbService() t.IDBService {
	dbService, dbServiceErr := is.NewDatabaseService(t.NewDBConfigWithDBConnection(ur.g), l.GetLogger("GodoBase"))
	if dbServiceErr != nil {
		gl.Log("error", fmt.Sprintf("UserModel repository: failed to get context DB service", dbServiceErr))
		return nil
	}
	return dbService
}
