package db

// =============================
// ===== Scannner handling =====
// =============================

func (source *SecurityDatabase) CreateScanner(name string) error {
	err := source.Database.Create(&DatabaseScanner{
		Name: name,
	}).Error
	return err
}

func (source *SecurityDatabase) GetScanner(name string) (DatabaseScanner, error) {
	result := DatabaseScanner{}

	err := source.Database.Model(&DatabaseScanner{}).First(&result, "name = ?", name).Error

	return result, err
}

func (source *SecurityDatabase) UpdateScanner(name string, s DatabaseScanner) error {
	currentScanner, err := source.GetScanner(name)
	if err != nil {
		return err
	}
	err = source.Database.Model(&currentScanner).Updates(s).Error
	return err
}

func (source *SecurityDatabase) DeleteScanner(name string) error {
	currentScanner, err := source.GetScanner(name)
	if err != nil {
		return err
	}
	err = source.Database.Delete(&currentScanner, currentScanner.ID).Error
	return err
}
