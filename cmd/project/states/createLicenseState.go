package states

import (
	"os"

	"github.com/AlecAivazis/survey/v2"

	"github.com/bragi-litlausson/UnstablePM/data"
)

func DefineLicenseType() string {
	licenseTypes := getLicenseNames()
	licensePrompt := &survey.Select{
		Message: "Choose license type",
		Options: licenseTypes,
	}
	licenseType := ""

	err := survey.AskOne(licensePrompt, &licenseType)
	if err != nil {
		panic(err)
	}

	licenseData := data.LicenseTypes[licenseType]
	createLicenseFile(licenseData.Template)

	return licenseType
}

func getLicenseNames() []string {
	keys := make([]string, 0, len(data.LicenseTypes))
	for k, _ := range data.LicenseTypes {
		keys = append(keys, k)
	}

	return keys
}

func createLicenseFile(template string) {
	f, err := os.Create("LICENSE")
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(template)
	if err != nil {
		panic(err)
	}
	f.Close()
}
