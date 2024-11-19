package fix_app

import (
	"os"
	"testing"

	"github.com/rrghoigoiwngoiw/hw-2/hw02_fix_app/reader"
)

func Test_ReadJSON_Success(t *testing.T) {
	// Создаем временный JSON-файл
	file, err := os.CreateTemp("", "test_data_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	content := `[
		{"UserID": 1, "Name": "John Doe", "Age": 30, "DepartmentID": 101},
		{"UserID": 2, "Name": "Jane Doe", "Age": 25, "DepartmentID": 102}
	]`
	_, err = file.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	file.Close()

	// Тестируем чтение JSON
	employees, err := reader.ReadJSON(file.Name())
	if err != nil {
		t.Errorf("Unexpected error while reading JSON: %v", err)
	}

	if len(employees) != 2 {
		t.Errorf("Expected 2 employees, got %d", len(employees))
	}

	if employees[0].Name != "John Doe" {
		t.Errorf("Expected first employee to be 'John Doe', got '%s'", employees[0].Name)
	}
}

func Test_ReadJSON_FileNotFound(t *testing.T) {
	_, err := reader.ReadJSON("nonexistent_file.json")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func Test_ReadJSON_InvalidJSON(t *testing.T) {
	// Создаем временный файл с некорректным JSON
	file, err := os.CreateTemp("", "invalid_data_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	content := `invalid_json_content`
	_, err = file.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	file.Close()

	_, err = reader.ReadJSON(file.Name())
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}
