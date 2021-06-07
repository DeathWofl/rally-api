package regresp

import (
	"gorm.io/gorm"
)

type store interface {
	ReadByID(id uint) (*RegResp, error)
	All() (*[]RegResp, error)
	Create(r *RegResp) (*RegResp, error)
	Update(id uint, r *RegResp) (*RegResp, error)
	Delete(id uint) error
}

type regRespStore struct {
	db *gorm.DB
}

func (s *regRespStore) ReadByID(id uint) (*RegResp, error) {
	var regResp RegResp
	s.db.First(&regResp, id)
	return &regResp, nil
}

func (s *regRespStore) All() (*[]RegResp, error) {
	var regResps []RegResp
	s.db.Find(&regResps)
	return &regResps, nil
}

func (s *regRespStore) Create(r *RegResp) (*RegResp, error) {
	s.db.Create(r)
	return r, nil
}

func (s *regRespStore) Update(id uint, r *RegResp) (*RegResp, error) {
	var regResp RegResp
	s.db.Find(&regResp, id)
	s.db.Model(&regResp).Updates(r)
	return &regResp, nil
}

func (s *regRespStore) Delete(id uint) error {
	s.db.Delete(&RegResp{}, id)
	return nil
}

func newStore(db *gorm.DB) (*regRespStore, error) {
	return &regRespStore{
		db: db,
	}, nil
}
