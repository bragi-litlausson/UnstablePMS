package core

func ProjectFileExists() bool {
	return FileExists(".upm")
}
