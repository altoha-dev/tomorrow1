package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestPrepareText(t *testing.T) {
	input := "Hello, world ... !?"
	// После правильной PrepareText знаки ДОЛЖНЫ быть отделены пробелами,
	// но группы должны быть целыми
	expected := "Hello , world ... !?"

	raw := PrepareText(input)
	result := strings.Join(strings.Fields(raw), " ")
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestFixPunctuations(t *testing.T) {
	words := []string{"Hello", ",", "world", "...", "!"}
	// Было: "world..."
	// Стало: "world...!" (потому что ! тоже должен приклеиться)
	expected := []string{"Hello,", "", "world...!", "", ""}
	FixPunctuations(words)
	if !reflect.DeepEqual(words, expected) {
		t.Errorf("Expected %v, got %v", expected, words)
	}
}

func TestHexBinCapLowUp(t *testing.T) {
	words := []string{"1E", "(hex)", "101", "(bin)", "ready", "(up)"}
	expected := []string{"30", "", "5", "", "READY", ""}
	HexBinCapLowUp(words)
	if !reflect.DeepEqual(words, expected) {
		t.Errorf("Expected %v, got %v", expected, words)
	}
}

func TestMultipleCapLowUp(t *testing.T) {
	words := []string{"this", "is", "so", "exciting", "(up,", "2)"}
	expected := []string{"this", "is", "SO", "EXCITING", "", ""}
	MultipleCapLowUp(words)
	if !reflect.DeepEqual(words, expected) {
		t.Errorf("Expected %v, got %v", expected, words)
	}
}

func TestHandleApostrophes(t *testing.T) {
	input := "As Elton John said: ' I am the most well-known homosexual in the world '"
	expected := "As Elton John said: 'I am the most well-known homosexual in the world'"
	result := HandleApostrophes(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestHandleArticles(t *testing.T) {
	words := []string{"A", "amazing", "rock"}
	expected := []string{"An", "amazing", "rock"}
	HandleArticles(words)
	if !reflect.DeepEqual(words, expected) {
		t.Errorf("Expected %v, got %v", expected, words)
	}
}
