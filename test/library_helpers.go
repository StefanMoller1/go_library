package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/StefanMoller1/go_library/models"
)

func GetBooks() ([]models.Book, error) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/library", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response models.Response

	json.Unmarshal(resBody, &response)

	jsonString, _ := json.Marshal(response.Data)

	newBooks := []models.Book{}
	json.Unmarshal(jsonString, &newBooks)

	return newBooks, nil
}

func CreateBook(book models.Book) (string, error) {
	jsonBody, err := json.Marshal(book)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/library", bytes.NewReader(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var response models.Response

	json.Unmarshal(resBody, &response)

	jsonString, _ := json.Marshal(response.Data)

	newBook := models.Book{}
	json.Unmarshal(jsonString, &newBook)

	return newBook.ID, nil
}

func GetBook(id string) (*models.Book, error) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/library/"+id, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response models.Response

	json.Unmarshal(resBody, &response)

	jsonString, _ := json.Marshal(response.Data)

	newBook := models.Book{}
	json.Unmarshal(jsonString, &newBook)

	return &newBook, nil
}

func UpdateBook(book models.Book) error {
	jsonBody, err := json.Marshal(book)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/library/"+book.ID, bytes.NewReader(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var response models.Response

	json.Unmarshal(resBody, &response)

	return nil
}

func DeleteBook(id string) error {
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8000/api/v1/library/"+id, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var response models.Response

	json.Unmarshal(resBody, &response)

	return nil
}
