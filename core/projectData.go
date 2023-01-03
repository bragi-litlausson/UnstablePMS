package core

type ProjectData struct {
	ProjectName string
	ProjectType ProjectType
	LicenseType string
	ShellData   NixShellData
}
