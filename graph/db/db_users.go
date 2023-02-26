package db

func (source *SecurityDatabase) CreateUser(name string) error {
	err := source.Database.Create(&DatabaseUser{
		Name: name,
	}).Error
	return err
}

func (source *SecurityDatabase) GetUser(name string) (DatabaseUser, error) {
	result := DatabaseUser{}

	err := source.Database.Model(&DatabaseUser{}).
		First(&result, "name = ?", name).Error

	return result, err
}

func (source *SecurityDatabase) UpdateUser(name string, p DatabaseUser) error {
	currentUser, err := source.GetUser(name)
	if err != nil {
		return err
	}
	err = source.Database.Model(&currentUser).Updates(p).Error
	return err
}

func (source *SecurityDatabase) DeleteUser(name string) error {
	currentUser, err := source.GetUser(name)
	if err != nil {
		return err
	}
	err = source.Database.Delete(&currentUser, currentUser.ID).Error
	return err
}

func (source *SecurityDatabase) AddUserRole(project DatabaseSecurityProject, name string, role string) error {
	user, err := source.GetUser(name)
	if err != nil {
		return err
	}
	err = source.Database.Model(&project).Association("Users").Error
	if err != nil {
		return err
	}
	err = source.Database.Model(&project).Association("Users").Append(&DatabaseUserRole{
		Role:   role,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	return nil
}
