package services

import (
	"github.com/MNU/exam-go"
	"github.com/pkg/errors"
)

// UserService is
type UserService struct {
	db *DB
}

var _ goexam.UserService = &UserService{}

// NewUserService is return UserService *
func NewUserService(db *DB) *UserService {
	return &UserService{
		db: db,
	}
}

// Login is 用户登陆
func (u *UserService) Login(account string, password string) error {
	user := new(goexam.User)
	err := u.db.Where("account = ? and password = ?", account, password).Find(user).Error
	return err
}

// Create is 添加用户
func (u *UserService) Create(user *goexam.User) error {
	if user.Name == "" {
		return errors.New("name was required")
	}

	if user.Role == "" {
		return errors.New("role was required")
	}

	if user.Account == "" {
		return errors.New("account was required")
	}

	if user.Password == "" {
		return errors.New("password was required")
	}

	if user.ClassID == 0 {
		return errors.New("class_id was required")
	}

	err := u.db.Create(user).Error
	return err
}

// Delete is 删除用户
func (u *UserService) Delete(ID uint) error {
	user := new(goexam.User)
	err := u.db.Where("id = ?", ID).Delete(user).Error
	return err
}

// Update is 更改用户
func (u *UserService) Update(user *goexam.User) error {
	err := u.db.Model(user).Updates(user).Error
	return err
}

// Get is 获取用户信息
func (u *UserService) Get(ID uint) (*goexam.User, error) {
	user := new(goexam.User)
	err := u.db.Preload("Class").First(user, ID).Error
	return user, err
}

// GetList is 获取用户列表
func (u *UserService) GetList(userFilter *goexam.UserFilter) (userList []*goexam.User, err error) {
	userList = make([]*goexam.User, 0)
	query := u.db.Model(&goexam.User{}).Preload("Class")
	if userFilter.Page != 0 {
		query = query.Offset(userFilter.Page * userFilter.Limit)
	}
	query = query.Limit(userFilter.Limit)
	err = query.Find(&userList).Error
	return userList, err
}
