package file

import (
	"testing"
)

func TestSave(t *testing.T) {
	dir := "library"
	files := map[string]string{
		"book_1.go": "hello book_1\n",
		"book_2.go": "hello book_2\n",
		"food_1.go": "hello food_1\n",
	}
	for name, content := range files {
		_, err := Save(dir, name, []byte(content))
		if err != nil {
			t.Error(err)
		}
	}
	names, err := Names(dir)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%s contains the file name: %s", dir, names)
	}

}

func TestDel(t *testing.T) {
	dir := "library"
	err := Del(dir, "book")
	if err != nil {
		t.Error(err)
	}
	names, err := Names(dir)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%s contains the file name: %s", dir, names)
	}
}

func TestAppendContent(t *testing.T) {
	path := "library/food_1.go"

	err := AppendContent(path, " test append\n")
	if err != nil {
		t.Error(err)
	}
	content, err := GetContent(path)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("the contents of the file %s:\n %s", path, content)
	}
}

func TestRemoveLine(t *testing.T) {
	path := "library/food_1.go"

	err := RemoveLine(path, " test append")
	if err != nil {
		t.Error(err)
	}
	content, err := GetContent(path)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("the contents of the file %s:\n %s", path, content)
	}
}
