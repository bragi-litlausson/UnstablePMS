package core

import "os"

func WriteTextFile(path string, data string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil {
		panic(err)
	}
}
