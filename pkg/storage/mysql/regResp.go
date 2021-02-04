type RegRespService struct {
	DB *gorm.DB
}

// RegResp retorna la regResp pedida en la BD
func (s *RegRespService) RegResp(id uint) (*models.RegResp, error) {
	var regResp models.RegResp
	s.DB.First(&regResp, id)
	return &regResp, nil
}

// RegResps retorna todas las regResps de la BD
func (s *RegRespService) RegResps() (*[]models.RegResp, error) {
	var regResps []models.RegResp
	s.DB.Find(&regResps)
	return &regResps, nil
}

// CreateRegResp agrega una regResp a la BD
func (s *RegRespService) CreateRegResp(r *models.RegResp) (*models.RegResp, error) {
	s.DB.Create(r)
	return r, nil
}

// UpdateRegResp actualiza una regResp en la BD
func (s *RegRespService) UpdateRegResp(id uint, r *models.RegResp) (*models.RegResp, error) {
	var regResp models.RegResp
	s.DB.Find(&regResp, id)
	s.DB.Model(&regResp).Updates(r)
	return &regResp, nil
}

// DeleteRegResp elimina una regResp de la BD
func (s *RegRespService) DeleteRegResp(id uint) error {
	s.DB.Delete(&models.RegResp{}, id)
	return nil
}
