package romanconv

import (
	"github.com/AdamCBernstein/romanconv"
	"testing"
)

type testResults struct {
	decimal    int
	romanValue string
}

var testValues = []testResults{
	{1, "I"},
	{3, "III"},
	{4, "IV"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{11, "XI"},
	{13, "XIII"},
	{14, "XIV"},
	{15, "XV"},
	{16, "XVI"},
	{19, "XIX"},
	{20, "XX"},
	{21, "XXI"},
	{24, "XXIV"},
	{25, "XXV"},
	{26, "XXVI"},
	{29, "XXIX"},
	{30, "XXX"},
	{39, "XXXIX"},
	{40, "XL"},
	{41, "XLI"},
	{44, "XLIV"},
	{45, "XLV"},
	{46, "XLVI"},
	{49, "XLIX"},
	{50, "L"},
	{51, "LI"},
	{54, "LIV"},
	{55, "LV"},
	{56, "LVI"},
	{98, "XCVIII"},
	{99, "XCIX"},
	{100, "C"},
	{111, "CXI"},
	{222, "CCXXII"},
	{333, "CCCXXXIII"},
	{888, "DCCCLXXXVIII"},
	{1888, "MDCCCLXXXVIII"},
	{2888, "MMDCCCLXXXVIII"},
	{3888, "MMMDCCCLXXXVIII"}, // longest possible roman numeral
}

func TestIntToRoman(t *testing.T) {
	_, err := romanconv.IntToRoman(4000)
	if err == nil {
		t.Errorf("wanted %q", "Input out of range")
	}

	_, err = romanconv.IntToRoman(0)
	if err == nil {
		t.Errorf("wanted %q", "Input out of range")
	}

	_, err = romanconv.IntToRoman(-1)
	if err == nil {
		t.Errorf("wanted %q", "Input out of range")
	}

	for _, testVal := range testValues {
		romanConv, err := IntToRoman(testVal.decimal)
		if err != nil {
			t.Errorf("Something bad happened in roman conversion: %v", err)
		}
		if romanConv != testVal.romanValue {
			t.Errorf("Decimal: %d got %q, wanted %q", testVal.decimal, romanConv, testVal.romanValue)
		}
	}
}
