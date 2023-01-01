package core

import (
	"fmt"

	"github.com/bragi-litlausson/UnstablePM/core/license"
)

func GetLicenseNames() []string {
	return GetMapKeys(license.LicenseTypes)
}

func CreateLicenseFile(licenseName string) error {
	data := license.LicenseTypes[licenseName]
	err := CreateFile("LICENSE", data.Template)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("Created %s LICENSE", licenseName)
	fmt.Println(msg)

	return nil
}
