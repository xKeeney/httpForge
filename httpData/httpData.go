package httpData

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func ReadBody(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func BytesToJson(data []byte, v any) error {
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}

// ReadRequestBody читает тело запроса и восстанавливает r.Body.
// Возвращает []byte для универсальной обработки.
// Если передан dest != nil — пытается распарсить JSON в dest.
func ReadRequestBody(r *http.Request, dest interface{}) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	// Восстанавливаем Body, чтобы downstream-обработчики могли читать его снова
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	// Если передана структура — пробуем распарсить JSON
	if dest != nil {
		if err := json.Unmarshal(body, dest); err != nil {
			return body, err
		}
	}

	return body, nil
}

func ResponseJSON(w http.ResponseWriter, data any, statusCode int) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonData)
	return nil
}

func ResponseJSONwithHeaders(w http.ResponseWriter, data any, statusCode int, headers map[string]string) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(statusCode)
	_, err = w.Write(jsonData)
	return err
}
