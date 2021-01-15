package handlers

import "testing"

func TestParseMail(t *testing.T) {
	var tests = []struct {
		r   dataItem
		exp bool
	}{
		{dataItem{Item: "John Daggett, 341 King Road, Plymouth MA"}, true},
		{dataItem{Item: "John Daggett, 341 King Road, Plymouth TX"}, false},
		{dataItem{Item: ""}, false},
		{dataItem{Item: "this should be wrong"}, false},
		{dataItem{Item: "2211244"}, false},
		{dataItem{Item: "helloworld"}, false},
		{dataItem{Item: "го - хороший язык"}, false},
		{dataItem{Item: "Now, I'm not sure about this one, no name, no address, but state, MA"}, true},
		{dataItem{Item: "There is no obvious way to check if a string contains a valid name and address, IN"}, true},
		{dataItem{Item: "So, let it be like this, OK"}, true},
	}
	for _, v := range tests {
		var got bool
		d := []dataItem{v.r}
		_, err := parseMail(d)
		if err != nil {
			got = false
		} else {
			got = true
		}
		if v.exp != got {
			t.Errorf(`wrong result for string: %v`, v.r.Item)
		}
	}
}
