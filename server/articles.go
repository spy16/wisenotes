package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/spy16/wisenotes/storage/db"
)

type Article struct {
	ID       int64     `json:"id"`
	Kind     string    `json:"kind"`
	Spec     string    `json:"spec"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Version  int64     `json:"version"`
	Profile  int64     `json:"profile"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
}

func listArticles(qu *db.Queries) http.HandlerFunc {
	return httpHandler(func(w http.ResponseWriter, r *http.Request) error {
		idStr := chi.URLParam(r, "profileID")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return err
		}

		rows, err := qu.GetArticles(r.Context(), id)
		if err != nil {
			return err
		}

		var articles []Article
		for _, row := range rows {
			articles = append(articles, Article{
				ID:       row.ID,
				Kind:     row.Kind,
				Spec:     row.Spec,
				Title:    row.Title,
				Content:  row.Content,
				Version:  row.Version,
				CreateAt: row.CreatedAt,
				UpdateAt: row.UpdatedAt,
			})
		}

		sendJSON(w, http.StatusOK, articles)
		return nil
	})
}

func createArticle(qu *db.Queries) http.HandlerFunc {
	return httpHandler(func(w http.ResponseWriter, r *http.Request) error {
		var article Article
		if err := readJSON(r, &article); err != nil {
			return err
		}

		return nil
	})
}

func readJSON(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return fmt.Errorf("%w: %s", errBadRequest, err.Error())
	}
	return nil
}
