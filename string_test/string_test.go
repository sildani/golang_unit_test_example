package string_test // "github.com/sildani/unit_test_example/string_test"

import (
  "testing"
  utestring "github.com/sildani/unit_test_example/string"
)

func TestReverse(t *testing.T) {
  input := "hello"
  output := utestring.Reverse(input)
  expected := "olleh"
  if output != expected {
    t.Errorf("Reverse(%q) == %q but expected %q", input, output, expected)
  }
}

func TestReverseUsingTable(t *testing.T) {
  var tests = []struct {
    input, expected string
  }{
    {"Backward", "drawkcaB"},
    {"Hello, 汉字", "字汉 ,olleH"},
    {"", ""},
  }

  for _, test := range tests {
    output := utestring.Reverse(test.input)
    if output != test.expected {
      t.Errorf("Reverse(%q) == %q but expected %q", test.input, output, test.expected)
    }
  }
}