package dali_test

import (
	"fmt"
	"testing"

	"github.com/matthewapeters/dali"
)

func TestDiv(t *testing.T) {

	P := dali.NewDiv("TestDivOne", "TestDivOne")
	html := fmt.Sprintf("%s", P)
	expected := `<div name="TestDivOne" id="TestDivOne"></div>`
	if html != expected {
		t.Errorf(`expected "%s" but got "%s"`, expected, html)
	}

	S := dali.NewSpanElement("TestSpanOne", "", "Test Text")
	P.Elements.AddElement(S)
	html = fmt.Sprintf("%s", P)
	expected = `<div name="TestDivOne" id="TestDivOne"><span name="TestSpanOne">Test Text</span></div>`
	if html != expected {
		t.Errorf(`expected "%s" but got "%s"`, expected, html)
	}
}
