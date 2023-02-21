package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jollej/item-catalog/pgm/models"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, _ := json.Marshal(models.GetItems())

	w.WriteHeader(http.StatusOK)
	w.Write(items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	items, _ := json.Marshal(models.GetItem(id))
	w.WriteHeader(http.StatusOK)
	w.Write(items)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	models.DeleteItem(id)
	w.WriteHeader(http.StatusOK)

}

func AddIitem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	models.AddItem(item)
}
