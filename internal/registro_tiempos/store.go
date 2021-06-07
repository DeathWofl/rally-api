package regtieps

import (
	"gorm.io/gorm"
)

type store interface {
	ReadByID(id uint) (*RegTiempo, error)
	All() (*[]RegTiempo, error)
	Create(r *RegTiempo) (*RegTiempo, error)
	Update(id uint, r *RegTiempo) (*RegTiempo, error)
	Delete(id uint) error
}

type regTiempoStore struct {
	db *gorm.DB
}

func (s *regTiempoStore) ReadByID(id uint) (*RegTiempo, error) {
	var regTiempo RegTiempo
	s.db.First(&regTiempo, id)
	return &regTiempo, nil
}

func (s *regTiempoStore) All() (*[]RegTiempo, error) {
	var regTiempos []RegTiempo
	s.db.Find(&regTiempos)
	return &regTiempos, nil
}

func (s *regTiempoStore) Create(r *RegTiempo) (*RegTiempo, error) {
	s.db.Create(r)
	return r, nil
}

func (s *regTiempoStore) Update(id uint, r *RegTiempo) (*RegTiempo, error) {
	var regTiempo RegTiempo
	s.db.Find(&regTiempo, id)
	s.db.Model(&regTiempo).Updates(r)
	return &regTiempo, nil
}

func (s *regTiempoStore) Delete(id uint) error {
	s.db.Delete(&RegTiempo{}, id)
	return nil
}

func newStore(db *gorm.DB) (*regTiempoStore, error) {
	return &regTiempoStore{
		db: db,
	}, nil
}
