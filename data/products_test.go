package data

import "testing"

func TestValidate(t *testing.T) {
	p := &Product{
		Name:  "Nics",
		Price: 2.00,
		SKU:   "asdf-fas-asdf",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
