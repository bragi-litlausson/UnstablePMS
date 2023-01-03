package core

import (
	"fmt"
)

func GetLicenseNames() []string {
	return GetMapKeys(LicenseTypes)
}

func CreateLicenseFile(licenseName string) error {
	data := LicenseTypes[licenseName]
	err := CreateFile("LICENSE", data.Template)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("Created %s LICENSE", licenseName)
	fmt.Println(msg)

	return nil
}
