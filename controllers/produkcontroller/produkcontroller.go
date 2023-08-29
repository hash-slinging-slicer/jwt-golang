package produkcontroller

import (
	"jwt/helper"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id":    1,
			"nama":  "Lenovo",
			"stock": 20,
		},
		{
			"id":    2,
			"nama":  "Asus",
			"stock": 10,
		},
		{
			"id":    3,
			"nama":  "HP",
			"stock": 18,
		},
	}

	// Return JSON
	helper.ResponeJSON(w, http.StatusOK, data)
}
