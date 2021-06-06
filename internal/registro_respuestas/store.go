package regresp

import (
	"gorm.io/gorm"
)

// RegRespService metodos disponibles para RegRes
type store interface {
	RegResp(id uint) (*RegResp, error)
	RegResps() (*[]RegResp, error)
	CreateRegResp(r *RegResp) (*RegResp, error)
	UpdateRegResp(id uint, r *RegResp) (*RegResp, error)
	DeleteRegResp(id uint) error
}

// RegRespService implementando metodos
type regRespStore struct {
	db *gorm.DB
}

// RegResp retorna la regResp pedida en la BD
func (s *regRespStore) RegResp(id uint) (*RegResp, error) {
	var regResp RegResp
	s.db.First(&regResp, id)
	return &regResp, nil
}

// RegResps retorna todas las regResps de la BD
func (s *regRespStore) RegResps() (*[]RegResp, error) {
	var regResps []RegResp
	s.db.Find(&regResps)
	return &regResps, nil
}

// CreateRegResp agrega una regResp a la BD
func (s *regRespStore) CreateRegResp(r *RegResp) (*RegResp, error) {
	s.db.Create(r)
	return r, nil
}

// UpdateRegResp actualiza una regResp en la BD
func (s *regRespStore) UpdateRegResp(id uint, r *RegResp) (*RegResp, error) {
	var regResp RegResp
	s.db.Find(&regResp, id)
	s.db.Model(&regResp).Updates(r)
	return &regResp, nil
}

// DeleteRegResp elimina una regResp de la BD
func (s *regRespStore) DeleteRegResp(id uint) error {
	s.db.Delete(&RegResp{}, id)
	return nil
}
