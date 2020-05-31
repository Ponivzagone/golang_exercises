package htree

import "testing"



func TestGetHeight1(t *testing.T) {
	HT := MakeHeightTree(7)
	HT.ReadFrom("4 -1 4 1 1 2 5")
	HT.ComputJump()
	if HT.GetHeight() != 4 {
		t.Error("Expected Height for tree: ", HT.tree, " and jump map: ", HT.jump)
	}
}

func TestGetHeight2(t *testing.T) {
	HT := MakeHeightTree(5)
	HT.ReadFrom("4 -1 4 1 1")
	HT.ComputJump()
	if HT.GetHeight() != 2 {
		t.Error("Expected Height for tree: ", HT.tree, " and jump map: ", HT.jump)
	}
}

func TestGetHeight3(t *testing.T) {
	HT := MakeHeightTree(1)
	HT.ReadFrom("-1")
	HT.ComputJump()
	if HT.GetHeight() != 0 {
		t.Error("Expected Height for tree: ", HT.tree, " and jump map: ", HT.jump)
	}
}

func TestGetHeight4(t *testing.T) {
	HT := MakeHeightTree(5)
	HT.ReadFrom("-1 0 1 2 3")
	HT.ComputJump()
	if HT.GetHeight() != 4 {
		t.Error("Expected Height for tree: ", HT.tree, " and jump map: ", HT.jump)
	}
}

func TestGetHeight5(t *testing.T) {
	HT := MakeHeightTree(8)
	HT.ReadFrom("7 0 3 1 6 6 7 -1")
	HT.ComputJump()
	if HT.GetHeight() != 4 {
		t.Error("Expected Height for tree: ", HT.tree, " and jump map: ", HT.jump)
	}
}
