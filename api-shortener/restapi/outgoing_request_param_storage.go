package restapi

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type IOutgoingRequestParamDAO interface {
	Create(api *OutgoingRequestParam) error
	Get(id uint) (*OutgoingRequestParam, error)
	GetAllByConfigID(configID uint) ([]*OutgoingRequestParam, error)
	Update(api *OutgoingRequestParam) error
	Delete(id uint) error
}

type OutgoingRequestParamDAO struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (dao *OutgoingRequestParamDAO) Create(param *OutgoingRequestParam) error {
	err := dao.validate.Struct(param)
	if err != nil {
		return err
	}
	return dao.db.Create(param).Error
}

func (dao *OutgoingRequestParamDAO) Get(id uint) (*OutgoingRequestParam, error) {
	result := &OutgoingRequestParam{}
	takeResult := dao.db.Where("ID = ?", id).Take(result)
	return result, takeResult.Error
}

func (dao *OutgoingRequestParamDAO) GetAllByConfigID(configID uint) ([]*OutgoingRequestParam, error) {
	var result []*OutgoingRequestParam
	takeResult := dao.db.Where("Outgoing_Request_Config_ID = ?", configID).Find(&result)
	return result, takeResult.Error
}

func (dao *OutgoingRequestParamDAO) Update(param *OutgoingRequestParam) error {
	err := dao.validate.Struct(param)
	if err != nil {
		return err
	}
	return dao.db.Updates(param).Error
}

func (dao *OutgoingRequestParamDAO) Delete(id uint) error {
	param, err := dao.Get(id)
	if err != nil {
		return err
	}
	return dao.db.Unscoped().Delete(param).Error
}

func NewOutgoingRequestParamDAO(conn *gorm.DB, validate *validator.Validate) IOutgoingRequestParamDAO {
	return &OutgoingRequestParamDAO{db: conn, validate: validate}
}
