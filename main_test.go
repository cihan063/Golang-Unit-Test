package main

import "testing"

func TestCompare(t *testing.T) {
	bdtxt := ReadBodyTxt()
	exampletxt := ReadExampleHtml()
	if bdtxt != exampletxt {
		t.Errorf("Files are not same")
	}
}

func TestSubtract(t *testing.T) {
	sub_res_HTML := GetSubResult()
	if sub_res_HTML != "80" {
		t.Errorf("Files are not same")
	}
}

func TestAdd(t *testing.T) {
	add_res_HTML := GetAddResult()
	if add_res_HTML != "100" {
		t.Errorf("Files are not same")
	}
}
