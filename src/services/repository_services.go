package services

import (
	"database/sql"
	"time"

	"github.com/Hamid-Ba/bama/common"
	"github.com/Hamid-Ba/bama/domain/models"
	"github.com/Hamid-Ba/bama/infrastructure/db"
	"github.com/Hamid-Ba/bama/pkg/logging"
	"gorm.io/gorm"
)

type RepositroyService[TEntity, TCreate, TUpdate, TResponse any] struct {
	db *gorm.DB
}

func NewRepositoryService[TEntity, TCreate, TUpdate, TResponse any]() *RepositroyService[TEntity, TCreate, TUpdate, TResponse] {
	return &RepositroyService[TEntity, TCreate, TUpdate, TResponse]{
		db: db.GetDb(),
	}
}

func (repo RepositroyService[TEntity, TCreate, TUpdate, TResponse]) GetBy(id int) (TResponse, error) {
	model := new(TEntity)
	res := new(TResponse)

	if err := repo.db.Model(model).Where("id = ? AND IsActive = ?", id, true).First(&res).Error; err != nil {
		return *res, err
	}

	return *res, nil
}

func (repo RepositroyService[TEntity, TCreate, TUpdate, TResponse]) GetList() (*[]TResponse, error) {
	model := new(TEntity)
	res := new([]TResponse)

	if err := repo.db.Model(model).Where("IsActive = ?", true).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (repo RepositroyService[TEntity, TCreate, TUpdate, TResponse]) Create(req TCreate) (TResponse, error) {
	var response TResponse

	entity, err := common.TypeConverter[TEntity](req)

	if err != nil {
		return response, err
	}

	tx := repo.db.Statement.Begin()

	if err := tx.Create(entity).Error; err != nil {
		tx.Rollback()
		logging.Log.Error(err.Error())
		return response, err
	}

	tx.Commit()

	res, _ := common.TypeConverter[models.BaseModel](entity)

	return repo.GetBy(res.Id)
}

func (repo RepositroyService[TEntity, TCreate, TUpdate, TResponse]) Update(id int, req TUpdate) (TResponse, error) {
	var model TEntity
	var response TResponse

	entity, err := common.TypeConverter[map[string]interface{}](req)

	if err != nil {
		return response, err
	}

	entity["Updated_at"] = sql.NullTime{Time: time.Now().UTC(), Valid: true}

	tx := repo.db.Statement.Begin()
	if err := tx.Model(model).Where("id = ?", id).Updates(entity).Error; err != nil {
		tx.Rollback()
		logging.Log.Error(err.Error())
		return response, err
	}
	tx.Commit()

	return repo.GetBy(id)
}

func (repo RepositroyService[TEntity, TCreate, TUpdate, TResponse]) Delete(id int) error {
	model := new(TEntity)

	tx := repo.db.Statement.Begin()
	if err := tx.Model(model).Where("id = ?", id).Update("IsActive", false).Error; err != nil {
		tx.Rollback()
		logging.Log.Error(err.Error())
		return err
	}

	tx.Commit()

	return nil
}
