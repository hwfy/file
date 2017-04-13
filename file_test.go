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
