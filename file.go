package file

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Save save the file to the specified directory and create it if the file or directory does not exist
func Save(dirName, fileName string, fileBuf []byte) (int64, error) {
	err := os.MkdirAll(dirName, 0777)
	if err != nil {
		return 0, err
	}
	file, err := os.OpenFile(dirName+"/"+fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	buf := bytes.NewReader(fileBuf)

	return io.Copy(file, buf)
}

// Del delete all matching files,such as Del(library, book)
//
// -library ( before )
//|____book_1.go
//|____book_2.go
//|____food_1.go
// -library ( after )
//|____food_1.go
func Del(dirName, fileName string) error {
	err := filepath.Walk(dirName, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if strings.Contains(f.Name(), fileName) {
			return os.Remove(path)
		}
		return nil
	})
	return err
}

// AppendContent append content to file if the content is repeated will be ignored
func AppendContent(fileName, content string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	//exist, _ := regexp.Match(content, data)

	//if the added content does not exist
	if !strings.Contains(string(data), content) {
		_, err = file.Write([]byte(content))
		if err != nil {
			return err
		}
	}
	return nil
}

// GetContent gets the contents of the specified file
func GetContent(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

// RemoveLine delete the specified row according to the character line
func RemoveLine(fileName, line string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	//escapes the special character \.+*?()|[]{}^$
	line = regexp.QuoteMeta(line)

	//(. *) can match the beginning or end of the line
	reg, err := regexp.Compile("(.*)" + line + "(.*)")
	if err != nil {
		return fmt.Errorf("Expression error,%s", err)
	}
	text := reg.ReplaceAllString(string(data), "")

	return ioutil.WriteFile(fileName, []byte(text), 0777)
}

// MkDir if the directory does not exist it is created, you can specify multiple directory names
func MkDir(paths ...string) {
	for _, path := range paths {
		if _, err := os.Stat(path); err != nil {
			os.Mkdir(path, os.ModePerm)
		}
	}
}

// Names if the current directory exists to return a map
// key is the file name, value is to remove the underscore first character in the file name
// -library
//|____book_a.go
//|____book_b.go
//|____food_c.go
//
// Names(library): map[string]string{"book_a":"Booka","book_b":"Bookb","food_c":"Foodc"}
func Names(dirName string) (map[string]string, error) {
	list := make(map[string]string)

	file, err := os.OpenFile(dirName, os.O_RDONLY, 0644)
	if err != nil {
		return list, err
	}
	names, err := file.Readdirnames(0)
	if err != nil {
		return list, err
	}
	for _, name := range names {
		//Remove the extension
		prefix := strings.SplitN(name, ".", 2)
		//Remove the underline
		items := strings.Split(prefix[0], "_")
		//Combination of strings
		name = strings.Join(items, "")

		list[prefix[0]] = strings.Title(name)
	}
	return list, nil
}
