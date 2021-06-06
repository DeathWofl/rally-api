package regtieps

import (
	"gorm.io/gorm"
)

// RegTiempoService metodos disponibles para RegTiempo
type store interface {
	RegTiempo(id uint) (*RegTiempo, error)
	RegTiempos() (*[]RegTiempo, error)
	CreateRegTiempo(r *RegTiempo) (*RegTiempo, error)
	UpdateRegTiempo(id uint, r *RegTiempo) (*RegTiempo, error)
	DeleteRegTiempo(id uint) error
}

// RegTiempoService implementando metodos
type regTiempoStore struct {
	db *gorm.DB
}

// RegTiempo retorna la regTiempo pedida en la BD
func (s *regTiempoStore) RegTiempo(id uint) (*RegTiempo, error) {
	var regTiempo RegTiempo
	s.db.First(&regTiempo, id)
	return &regTiempo, nil
}

// RegTiempos retorna todas las regTiempos de la BD
func (s *regTiempoStore) RegTiempos() (*[]RegTiempo, error) {
	var regTiempos []RegTiempo
	s.db.Find(&regTiempos)
	return &regTiempos, nil
}

// CreateRegTiempo agrega una regTiempo a la BD
func (s *regTiempoStore) CreateRegTiempo(r *RegTiempo) (*RegTiempo, error) {
	s.db.Create(r)
	return r, nil
}

// UpdateRegTiempo actualiza una regTiempo en la BD
func (s *regTiempoStore) UpdateRegTiempo(id uint, r *RegTiempo) (*RegTiempo, error) {
	var regTiempo RegTiempo
	s.db.Find(&regTiempo, id)
	s.db.Model(&regTiempo).Updates(r)
	return &regTiempo, nil
}

// DeleteRegTiempo elimina una regTiempo de la BD
func (s *regTiempoStore) DeleteRegTiempo(id uint) error {
	s.db.Delete(&RegTiempo{}, id)
	return nil
}

func newService(db *gorm.DB) (*regTiempoStore, error) {
	return &regTiempoStore{
		db: db,
	}, nil
}
