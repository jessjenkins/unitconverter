package main

import "testing"

func TestConverter(t *testing.T) {

	weightConverter := converter{
		"weight",
		map[string]float64{
			"kg": 1000,
			"g":  1,
		},
	}
	if name := weightConverter.Name; name != "weight" {
		t.Errorf("Name: expected \"weight\" got \"%v\"", name)
	}
	if kg := weightConverter.Unit["kg"]; kg != 1000 {
		t.Errorf("KG: expected 1000 got %v", kg)
	}

	tests := []struct {
		from, to      string
		val, expected float64
	}{
		{"kg", "g", 1, 1000},
		{"g", "kg", 1, 0.001},
		{"kg", "g", 1000, 1000000},
		{"g", "kg", 1000, 1},
	}

	for _, c := range tests {
		g := weightConverter.convert(c.from, c.to, c.val)
		if g != c.expected {
			t.Errorf("Convert %v%s to %s: expected %v got %v", c.val, c.from, c.to, c.expected, g)
		}
	}

}
