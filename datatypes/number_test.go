package datatypes

import (
    "testing"
)

var (
    exampleFloat = NewFloatFromFloat64(0)
    exampleInt = NewIntFromInt64(0)
    exampleRational = NewRationalFromInt(exampleInt)
    exampleByte = NewByte(12)
)

func TestNumbersAreLispValues(t *testing.T) {
    dummyLispValueReceiver(exampleFloat)
    dummyLispValueReceiver(exampleInt)
    dummyLispValueReceiver(exampleRational)
    dummyLispValueReceiver(exampleByte)
}

func TestNumbersAreLispNumbers(t *testing.T) {
    dummyLispNumberReceiver(exampleFloat)
    dummyLispNumberReceiver(exampleInt)
    dummyLispNumberReceiver(exampleRational)
    dummyLispNumberReceiver(exampleByte)
}


func dummyLispValueReceiver(lv LispValue) {
}

func dummyLispNumberReceiver(ln LispNumber) {
}
