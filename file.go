//Copyright (c) 2017, hwfy

//Permission to use, copy, modify, and/or distribute this software for any
//purpose with or without fee is hereby granted, provided that the above
//copyright notice and this permission notice appear in all copies.

//THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
//WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
//MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
//ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
//WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
//ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
//OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package file

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// MkDir if the directory does not exist it is created,
// you can specify multiple directory names
func MkDir(paths ...string) {
	for _, path := range paths {
		if _, err := os.Stat(path); err != nil {
			os.Mkdir(path, os.ModePerm)
		}
	}
}

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

// Write append content to file if the content is repeated will be ignored
func Write(fileName, content string) error {
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

// Read gets the contents of the specified file
func Read(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

// RemoveSame delete all matching files,such as Del(library, book)
//
// -library ( before )
//|____book_1.go
//|____book_2.go
//|____food_1.go
// -library ( after )
//|____food_1.go
func RemoveSame(dirName, fileName string) error {
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

// Remove delete the specified file
func Remove(path string) error {
	return os.Remove(filepath.Clean(path))
}

// RemoveLine delete the specified row according to the character line
func RemoveLine(fileName, line string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	//escapes the special character \.+*?()|[]{}^$
	line = strings.TrimSpace(regexp.QuoteMeta(line))

	//(. *) can match the beginning or end of the line
	reg, err := regexp.Compile("(.*)" + line + "(.*)")
	if err != nil {
		return fmt.Errorf("Expression error,%s", err)
	}
	text := reg.ReplaceAllString(string(data), "")

	return ioutil.WriteFile(fileName, []byte(text), 0777)
}

// Names if the current directory exists to return a map key is the file name,
// value is to remove the underscore first character in the file name
// -library
//|____book_a.go
//|____book_b.go
//|____food_c.go
//
// Names(library): map[string]string{"book_a":"BookA","book_b":"BookB","food_c":"FoodC"}
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
		for i := 0; i < len(items); i++ {
			items[i] = strings.Title(items[i])
		}
		list[prefix[0]] = strings.Join(items, "")
	}
	return list, nil
}

// Generator generates a go file based on name and source code and
// automatically introduces the package name
func Generator(path, code string) error {
	// format code
	cmd := exec.Command("gofmt")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	io.WriteString(stdin, code)
	stdin.Close()

	bytes, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Failed to generate file：" + path + " syntax error")
	}
	// write file
	shortPath := filepath.Clean(path + ".go")

	file, err := os.Create(shortPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}
	// import package
	return exec.Command("goimports", "-w", shortPath).Run()
}
