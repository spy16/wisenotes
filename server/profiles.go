package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/spy16/wisenotes/storage/db"
)

type Profile struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Emoji     string    `json:"emoji"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func getProfiles(qu *db.Queries) http.HandlerFunc {
	return httpHandler(func(w http.ResponseWriter, r *http.Request) error {
		profiles, err := qu.GetProfiles(r.Context())
		if err != nil {
			return err
		}

		var ps []Profile
		for _, p := range profiles {
			ps = append(ps, Profile{
				ID:        p.ID,
				Name:      p.Name,
				Emoji:     p.Emoji,
				CreatedAt: p.CreatedAt.Time,
				UpdatedAt: p.UpdatedAt.Time,
			})
		}

		sendJSON(w, http.StatusOK, ps)
		return nil
	})
}

func createProfile(qu *db.Queries) http.HandlerFunc {
	return httpHandler(func(w http.ResponseWriter, r *http.Request) error {
		var p Profile
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			return err
		}

		if err := qu.CreateProfile(r.Context(), db.CreateProfileParams{
			Name:  p.Name,
			Emoji: p.Emoji,
		}); err != nil {
			return err
		}

		pNew, err := qu.GetProfileByName(r.Context(), p.Name)
		if err != nil {
			return err
		}

		sendJSON(w, http.StatusCreated, Profile{
			ID:        p.ID,
			Name:      pNew.Name,
			Emoji:     pNew.Emoji,
			CreatedAt: pNew.CreatedAt.Time,
			UpdatedAt: pNew.UpdatedAt.Time,
		})
		return nil
	})
}
