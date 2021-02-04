type RegTiempoService struct {
	DB *gorm.DB
}

// RegTiempo retorna la regTiempo pedida en la BD
func (s *RegTiempoService) RegTiempo(id uint) (*models.RegTiempo, error) {
	var regTiempo models.RegTiempo
	s.DB.First(&regTiempo, id)
	return &regTiempo, nil
}

// RegTiempos retorna todas las regTiempos de la BD
func (s *RegTiempoService) RegTiempos() (*[]models.RegTiempo, error) {
	var regTiempos []models.RegTiempo
	s.DB.Find(&regTiempos)
	return &regTiempos, nil
}

// CreateRegTiempo agrega una regTiempo a la BD
func (s *RegTiempoService) CreateRegTiempo(r *models.RegTiempo) (*models.RegTiempo, error) {
	s.DB.Create(r)
	return r, nil
}

// UpdateRegTiempo actualiza una regTiempo en la BD
func (s *RegTiempoService) UpdateRegTiempo(id uint, r *models.RegTiempo) (*models.RegTiempo, error) {
	var regTiempo models.RegTiempo
	s.DB.Find(&regTiempo, id)
	s.DB.Model(&regTiempo).Updates(r)
	return &regTiempo, nil
}

// DeleteRegTiempo elimina una regTiempo de la BD
func (s *RegTiempoService) DeleteRegTiempo(id uint) error {
	s.DB.Delete(&models.RegTiempo{}, id)
	return nil
}
