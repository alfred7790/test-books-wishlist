package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"test-books-wishlist/internal/common"
	"test-books-wishlist/internal/entity"
	"test-books-wishlist/internal/handler/mocks"
	"test-books-wishlist/internal/service"
	"testing"
	"time"
)

func TestAPI_GetWishList(t *testing.T) {
	tests := []struct {
		expectedBody interface{}
		expectedCode int
		mockFunc     func(wishlistID uint) (*entity.WishList, string, error)
	}{
		{
			expectedBody: nil,
			expectedCode: http.StatusOK,
			mockFunc: func(wishlistID uint) (*entity.WishList, string, error) {
				return &entity.WishList{
					ID:        1,
					UserID:    1,
					Name:      "First",
					CreatedAt: time.Now(),
					Items:     make([]*entity.ItemWishList, 0),
				}, "", nil
			},
		},
	}

	for _, test := range tests {
		app := service.NewService()
		mockRepo := &mocks.RepositoryMock{}
		app.Repo = mockRepo
		mockRepo.GetBooksByWishListMockFn = test.mockFunc

		api := API{
			App: app,
		}

		rootApp := gin.Default()
		r := rootApp.Group("v1")
		wishlist := r.Group("wishlist/:id")

		token, err := common.GetToken(1, "1")
		if err != nil {
			t.Fatalf("unexpected error.  Expected %v, got: %d", nil, err)
		}
		wishlist.GET("", api.GetWishList)
		req, _ := http.NewRequest(http.MethodGet, "/v1/wishlist/1", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		recorder := httptest.NewRecorder()
		rootApp.ServeHTTP(recorder, req)

		if test.expectedCode != recorder.Code {
			t.Fatalf("unexpected HTTP response code.  Expected %d, got: %d", test.expectedCode, recorder.Code)
		}
	}
}
