package file

import (
	"strings"
	"testing"
)

func TestFile(t *testing.T) {
	dir := "library"
	path := "library/food_1.go"
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

	err := Del(dir, "book")
	if err != nil {
		t.Error(err)
	}

	names, err := Names(dir)
	if err != nil {
		t.Error(err)
	}
	if names["food_1"] != "Food1" {
		t.Errorf("The value of food_1 should be Food1, but it is %s", names["food_1"])
	}

	err = AppendContent(path, " test append\n")
	if err != nil {
		t.Error(err)
	}

	err = RemoveLine(path, " test append")
	if err != nil {
		t.Error(err)
	}

	content, err := GetContent(path)
	if err != nil {
		t.Error(err)
	}
	if strings.TrimSpace(string(content)) != "hello food_1" {
		t.Errorf("The contents of the file %s should be 'hello food_1', but it is '%s'", path, content)
	}
}
