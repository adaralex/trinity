package db

import (
	"fmt"
	"testing"
)

func TestProcess(t *testing.T) {
	databaseConnector := &SecurityDatabase{}
	err := databaseConnector.Open()
	if err != nil {
		t.Errorf("error loading database: %+v", err.Error())
	}

	err = databaseConnector.CreateProject("Flexwitch")
	if err != nil {
		t.Errorf("error creating project: %+v", err.Error())
	}
	err = databaseConnector.CreateScanner("Nuclei")
	if err != nil {
		t.Errorf("error creating scanner: %+v", err.Error())
	}
	scanner, err := databaseConnector.GetScanner("Nuclei")
	if err != nil {
		t.Errorf("error creating scanner: %+v", err.Error())
	}
	if scanner.Name != "Nuclei" {
		t.Errorf("scanner not stored correctly")
	}
	project, err := databaseConnector.GetProject("Flexwitch")
	if err != nil {
		t.Errorf("error getting project: %+v", err.Error())
	}
	err = databaseConnector.AddScannerAnalysis(project, "Nuclei", "* * * * * *", 12)
	if err != nil {
		t.Errorf("error adding project analysis: %+v", err.Error())
	}
	project, err = databaseConnector.GetProject("Flexwitch")
	if err != nil {
		t.Errorf("error getting project: %+v", err.Error())
	}
	if project.Analysis == nil || len(project.Analysis) == 0 {
		t.Errorf("analysis not stored correctly")
	} else {
		if project.Analysis[0].Scanner.Name != "Nuclei" {
			t.Errorf("wrongful scanner definition")
		}
	}

	err = databaseConnector.UpdateScanner("Nuclei", DatabaseScanner{
		Install: "go get nuclei",
		Run:     "go run nuclei",
		Report:  "cat report.json",
		Type:    "URL",
	})
	if err != nil {
		t.Errorf("error updating scanner: %+v", err.Error())
	}

	err = databaseConnector.CreateUser("pano")
	if err != nil {
		t.Errorf("error creating user: %+v", err.Error())
	}
	user, err := databaseConnector.GetUser("pano")
	if err != nil {
		t.Errorf("error getting user: %+v", err.Error())
	}
	fmt.Printf("User: %+v\n", user)

	err = databaseConnector.AddUserRole(project, "pano", "admin")
	if err != nil {
		t.Errorf("error adding user role: %+v", err.Error())
	}

	project, err = databaseConnector.GetProject("Flexwitch")
	if err != nil {
		t.Errorf("error getting project: %+v", err.Error())
	}
	fmt.Printf("Project: %+v\n", project)

	err = databaseConnector.DeleteScannerAnalysis(project, project.Analysis[0])
	if err != nil {
		t.Errorf("error removing analysis: %+v", err.Error())
	}
	project, err = databaseConnector.GetProject("Flexwitch")
	if err != nil {
		t.Errorf("error getting project: %+v", err.Error())
	}
	fmt.Printf("Project: %+v\n", project)

}
