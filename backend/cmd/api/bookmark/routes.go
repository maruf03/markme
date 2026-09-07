package api

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/maruf03/markme/backend/internal/bookmark"
)

type BookmarkController struct {
	bookmarkService *bookmark.BookmarkService
	routeGroup      *echo.Group
}

func NewBookmarkController(
	bookmarkService *bookmark.BookmarkService,
	routeGroup *echo.Group,
) *BookmarkController {
	return &BookmarkController{
		bookmarkService: bookmarkService,
		routeGroup:      routeGroup,
	}
}

func (b *BookmarkController) RegisterRoutes() {
	b.routeGroup.GET("", b.GetAllBookmarks)
	b.routeGroup.POST("", b.CreateBookmark)
	b.routeGroup.GET("/:bookmarkId", b.GetBookmark)
	b.routeGroup.PATCH("/:bookmarkId", b.UpdateBookmark)
	b.routeGroup.DELETE("/:bookmarkId", b.DeleteBookmark)
}

func (b *BookmarkController) GetAllBookmarks(e *echo.Context) error {
	return e.JSON(http.StatusOK, nil)
}

func (b *BookmarkController) CreateBookmark(e *echo.Context) error {
	return e.JSON(http.StatusOK, nil)
}

func (b *BookmarkController) GetBookmark(e *echo.Context) error {
	return e.JSON(http.StatusOK, nil)
}

func (b *BookmarkController) UpdateBookmark(e *echo.Context) error {
	return e.JSON(http.StatusOK, nil)
}

func (b *BookmarkController) DeleteBookmark(e *echo.Context) error {
	return e.JSON(http.StatusOK, nil)
}
