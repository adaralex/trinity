package db

import "gorm.io/gorm"

type DatabaseSecurityProject struct {
	gorm.Model
	Name            string                       `json:"name" gorm:"uniqueIndex"`
	Users           []DatabaseUserRole           `json:"users" gorm:"many2many:pj_user_roles;"`
	Analysis        []DatabaseScannerAnalysis    `json:"analysis" gorm:"many2many:pj_scanner_analysis;"`
	Vulnerabilities []DatabaseSecurityIssue      `json:"vulnerabilities" gorm:"many2many:pj_vulns"`
	ProjectAssets   []DatabaseProjectAssets      `json:"assets" gorm:"many2many:pj_assets"`
	Credentials     []DatabaseProjectCredentials `json:"credentials" gorm:"many2many:pj_creds"`
	Params          []DatabaseProjectParameters  `json:"params" gorm:"many2many:pj_params"`
}
type DatabaseScanner struct {
	gorm.Model
	Name    string `json:"name"`
	Install string `json:"install"`
	Run     string `json:"run"`
	Report  string `json:"report"`
	Type    string `json:"type"`
}
type DatabaseScannerAnalysis struct {
	gorm.Model
	ScannerID uint
	Scanner   DatabaseScanner `gorm:"foreignKey:ScannerID"`
	Cron      string          `json:"cron"`
	//Params     []string `json:"params"`
	Timeout int `json:"timeout"`
}
type DatabaseProjectAssets struct {
	gorm.Model
	Details   string `json:"details"`
	TypeAsset string `json:"type_asset"`
}
type DatabaseProjectParameters struct {
	gorm.Model
	Label string `json:"label"`
	Value string `json:"value"`
}
type DatabaseProjectCredentials struct {
	gorm.Model
	Label string `json:"label"`
	Value string `json:"value"`
}
type DatabaseUserRole struct {
	gorm.Model
	Role   string `json:"role"`
	UserID uint
	User   DatabaseUser `gorm:"foreignKey:UserID"`
}
type DatabaseUser struct {
	gorm.Model
	Name string `json:"name"`
}
type DatabaseSecurityIssue struct {
	gorm.Model
	OriginalCvss float64 `json:"original_cvss"`
	RevisedCvss  float64 `json:"revised_cvss"`
	AnalysisDate string  `json:"analysis_date"`
	Scanner      string  `json:"scanner"`
	Cve          string  `json:"cve"`
	Cwe          string  `json:"cwe"`
	Vex          string  `json:"vex"`
	Infos        string  `json:"infos"`
	Status       string  `json:"status"`
	Origin       string  `json:"origin"`
}
