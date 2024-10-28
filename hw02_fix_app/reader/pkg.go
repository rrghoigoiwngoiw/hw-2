package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/fixme_my_friend/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	// Открытие файла
	f, err := os.Open(filePath)
	if err != nil {
		// Возвращаем ошибку, если файл не удалось открыть
		return nil, fmt.Errorf("ошибка при открытии файла: %w", err)
	}
	// Гарантируем закрытие файла после выполнения функции
	defer f.Close()
	// Чтение содержимого файла
	bytes, err := io.ReadAll(f)
	if err != nil {
		// Возвращаем ошибку, если чтение не удалось

		return nil, fmt.Errorf("ошибка при чтении файла: %w", err)
	}

	// Распаковка JSON в слайс сотрудников
	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		// Возвращаем ошибку, если Unmarshal не удался
		return nil, fmt.Errorf("ошибка при распаковке JSON: %w", err)
	}

	return data, nil
}
