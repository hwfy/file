package file

import (
	"testing"
)

func TestSave(t *testing.T) {
	dir := "library"
	data := map[string]string{
		"book_1.go": "hello book_1",
		"book_2.go": "hello book_2",
		"food_1.go": "hello food_1",
	}
	for name, content := range data {
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
	dir := "library"
	name := "food_1.go"

	err := AppendContent(dir+"/"+name, "\n test append")
	if err != nil {
		t.Error(err)
	}
	content, err := Get(dir, name)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("the contents of the file %s:\n %s", name, content)
	}
}

func TestRemoveLine(t *testing.T) {
	dir := "library"
	name := "food_1.go"

	err := RemoveLine(dir+"/"+name, "\n test append")
	if err != nil {
		t.Error(err)
	}
	content, err := Get(dir, name)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("the contents of the file %s:\n %s", name, content)
	}
}
