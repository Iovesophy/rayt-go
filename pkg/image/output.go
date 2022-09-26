package image

import "os"

func (img Elements) CreateFile(filename string, header string, body string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(header)
	if err != nil {
		return err
	}
	_, err = file.WriteString(body)
	if err != nil {
		return err
	}
	return nil
}
