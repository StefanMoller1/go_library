package router

import (
	"context"
	"net/http"
	"strconv"

	"github.com/StefanMoller1/go_library/models"
)

type ContextKey string

const paginationContext ContextKey = "pagination"

func (m *Manager) Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			page = 0
		}

		size, err := strconv.Atoi(r.URL.Query().Get("size"))
		if err != nil || size == 0 {
			size = 10
		}

		ctx := context.WithValue(r.Context(), paginationContext, &models.Pagination{Page: int32(page), Size: int32(size)})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
