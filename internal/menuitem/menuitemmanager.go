package menuitem

import (
	"gorm.io/gorm"
)

type MenuItemManager interface {
	GetMenuItems() (GetResponse, error)
	GetMenuItemWithID(GetWithIDRequest) (GetWithIDResponse, error)
	CreateMenuItem(CreateRequest) (CreateResponse, error)
	DeleteMenuItem(DeleteRequest) (DeleteResponse, error)
}

type MenuItemManagerImpl struct {
	database *gorm.DB
}

func NewMenuItemManager(database *gorm.DB) MenuItemManager {
	return &MenuItemManagerImpl{
		database: database,
	}
}

func (m *MenuItemManagerImpl) GetMenuItems() (resp GetResponse, err error) {
	return
}

func (m *MenuItemManagerImpl) GetMenuItemWithID(req GetWithIDRequest) (resp GetWithIDResponse, err error) {
	return
}

func (m *MenuItemManagerImpl) CreateMenuItem(req CreateRequest) (resp CreateResponse, err error) {
	newMenuItem := &MenuItem{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CreatedBy:   req.UserId,
		Ingredients: req.Ingredients,
	}

	err = m.database.Create(&newMenuItem).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
	}
	return
}

func (m *MenuItemManagerImpl) DeleteMenuItem(req DeleteRequest) (resp DeleteResponse, err error) {
	return
}
