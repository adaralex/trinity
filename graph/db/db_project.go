package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ============================
// ===== Project handling =====
// ============================

func (source *SecurityDatabase) CreateProject(name string) error {
	err := source.Database.Create(&DatabaseSecurityProject{
		Name: name,
	}).Error
	return err
}

func (source *SecurityDatabase) GetProject(id string) (DatabaseSecurityProject, error) {
	result := DatabaseSecurityProject{}

	err := source.Database.Model(&DatabaseSecurityProject{}).
		//Preload("Analysis").
		Preload("Analysis.Scanner").
		Preload("Users.User").
		//Preload("Users").
		//Preload("Vulnerabilities").
		//Preload("ProjectAssets").
		//Preload("Credentials").
		//Preload("Params").
		Preload(clause.Associations).
		First(&result, "id = ?", id).Error

	return result, err
}

func (source *SecurityDatabase) FindProject(name string) (DatabaseSecurityProject, error) {
	result := DatabaseSecurityProject{}

	err := source.Database.Model(&DatabaseSecurityProject{}).
		//Preload("Analysis").
		Preload("Analysis.Scanner").
		Preload("Users.User").
		//Preload("Users").
		//Preload("Vulnerabilities").
		//Preload("ProjectAssets").
		//Preload("Credentials").
		//Preload("Params").
		Preload(clause.Associations).
		First(&result, "name = ?", name).Error

	return result, err
}

func (source *SecurityDatabase) UpdateProject(name string, p DatabaseSecurityProject) error {
	currentProject, err := source.GetProject(name)
	if err != nil {
		return err
	}
	err = source.Database.Model(&currentProject).Updates(p).Error
	return err
}

func (source *SecurityDatabase) DeleteProject(name string) error {
	currentProject, err := source.GetProject(name)
	if err != nil {
		return err
	}
	err = source.Database.Delete(&currentProject, currentProject.ID).Error
	return err
}

func (source *SecurityDatabase) AddScannerAnalysis(project DatabaseSecurityProject, name string, cron string, timeout int) error {
	scanner, err := source.GetScanner(name)
	if err != nil {
		return err
	}
	err = source.Database.Model(&project).Association("Analysis").Error
	if err != nil {
		return err
	}
	err = source.Database.Model(&project).Association("Analysis").Append(&DatabaseScannerAnalysis{
		ScannerID: scanner.ID,
		Cron:      cron,
		Timeout:   timeout,
	})
	if err != nil {
		return err
	}
	return nil
}

func (source *SecurityDatabase) DeleteScannerAnalysis(project DatabaseSecurityProject, analysis DatabaseScannerAnalysis) error {
	err := source.Database.Model(&project).Association("Analysis").Error
	if err != nil {
		return err
	}
	err = source.Database.Model(&project).Association("Analysis").Delete(&DatabaseScannerAnalysis{
		Model: gorm.Model{ID: analysis.ID},
	})
	return err
}

func (source *SecurityDatabase) AddProjectAsset(project DatabaseSecurityProject, detail string, typeAsset string) error {
	err := source.Database.Model(&project).Association("ProjectAssets").Error
	if err != nil {
		return err
	}
	err = source.Database.Model(&project).Association("ProjectAssets").Append(&DatabaseProjectAssets{
		Details:   detail,
		TypeAsset: typeAsset,
	})
	if err != nil {
		return err
	}
	return nil
}

func (source *SecurityDatabase) DeleteProjectAsset(project DatabaseSecurityProject, asset DatabaseProjectAssets) error {
	err := source.Database.Model(&project).Association("ProjectAssets").Error
	if err != nil {
		return err
	}
	err = source.Database.Model(&project).Association("ProjectAssets").Delete(&DatabaseProjectAssets{
		Model: gorm.Model{ID: asset.ID},
	})
	return err
}

func (source *SecurityDatabase) AddProjectCredential(project DatabaseSecurityProject, name string, value string) error {
	err := source.Database.Model(&project).Association("Credentials").Error
	if err != nil {
		return err
	}
	err = source.Database.Model(&project).Association("Credentials").Append(&DatabaseProjectCredentials{
		Label: name,
		Value: value,
	})
	if err != nil {
		return err
	}
	return nil
}

func (source *SecurityDatabase) DeleteProjectCredential(project DatabaseSecurityProject, cred DatabaseProjectCredentials) error {
	err := source.Database.Model(&project).Association("Credentials").Error
	if err != nil {
		return err
	}
	err = source.Database.Model(&project).Association("Credentials").Delete(&DatabaseProjectCredentials{
		Model: gorm.Model{ID: cred.ID},
	})
	return err
}

func (source *SecurityDatabase) AddProjectParam(project DatabaseSecurityProject, name string, value string) error {
	err := source.Database.Model(&project).Association("Params").Error
	if err != nil {
		return err
	}
	err = source.Database.Model(&project).Association("Params").Append(&DatabaseProjectParameters{
		Label: name,
		Value: value,
	})
	if err != nil {
		return err
	}
	return nil
}

func (source *SecurityDatabase) DeleteProjectParam(project DatabaseSecurityProject, param DatabaseProjectParameters) error {
	err := source.Database.Model(&project).Association("Params").Error
	if err != nil {
		return err
	}
	err = source.Database.Model(&project).Association("Params").Delete(&DatabaseProjectParameters{
		Model: gorm.Model{ID: param.ID},
	})
	return err
}

func (source *SecurityDatabase) AddVulnerability(
	project DatabaseSecurityProject,
	cvss float64,
	date string,
	scanner string,
	cve string,
	cwe string,
	vex string,
	infos string,
	status string,
	origin string) error {
	err := source.Database.Model(&project).Association("Vulnerabilities").Error
	if err != nil {
		return err
	}
	err = source.Database.Model(&project).Association("Vulnerabilities").Append(&DatabaseSecurityIssue{
		OriginalCvss: cvss,
		RevisedCvss:  cvss,
		AnalysisDate: date,
		Scanner:      scanner,
		Cve:          cve,
		Cwe:          cwe,
		Vex:          vex,
		Infos:        infos,
		Status:       status,
		Origin:       origin,
	})
	if err != nil {
		return err
	}
	return nil
}

func (source *SecurityDatabase) DeleteVulnerability(project DatabaseSecurityProject, param DatabaseSecurityIssue) error {
	err := source.Database.Model(&project).Association("Vulnerabilities").Error
	if err != nil {
		return err
	}
	err = source.Database.Model(&project).Association("Vulnerabilities").Delete(&DatabaseSecurityIssue{
		Model: gorm.Model{ID: param.ID},
	})
	return err
}
