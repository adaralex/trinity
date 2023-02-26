package graph

import (
	"fmt"
	"github.com/adaralex/trinity/graph/db"
	"github.com/adaralex/trinity/graph/model"
)

func ProjectToGQL(p db.DatabaseSecurityProject) model.Project {
	userList := make([]*model.UserRole, 0)
	for _, u := range p.Users {
		userList = append(userList, &model.UserRole{
			Role: u.Role,
			Name: u.User.Name,
		})
	}
	analysisList := make([]*model.ScannerAnalysis, 0)
	for _, a := range p.Analysis {
		analysisList = append(analysisList, &model.ScannerAnalysis{
			Scanner: a.Scanner.Name,
			Cron:    a.Cron,
			Params:  nil,
			Timeout: &a.Timeout,
		})
	}

	result := model.Project{
		IDProject:       fmt.Sprintf("%d", p.ID),
		Name:            p.Name,
		Users:           userList,
		Analysis:        analysisList,
		Vulnerabilities: make([]*model.Vulnerability, 0),
		ProjectAssets:   make([]*model.ProjectAssets, 0),
		Credentials:     make([]*model.ProjectCredentials, 0),
		Params:          make([]*model.Parameters, 0),
	}
	return result
}
