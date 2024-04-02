package handlers

import (
	"app/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type AccountHandler struct {
	accountService AccountService
}

func NewAccountHandler(ac AccountService) *AccountHandler {
	return &AccountHandler{accountService: ac}
}

func (h *AccountHandler) Verify(c echo.Context) error {
	isValid := h.accountService.Verify(c)

	if isValid {
		return c.JSON(http.StatusOK, map[string]interface{}{"isAuthenticated": true})
	} else {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"isAuthenticated": false})
	}
}

func (h *AccountHandler) Register(c echo.Context) error {
	var user models.UserAuthRequest
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	// usernameの重複チェック
	usernames, err := h.accountService.GetAllUsernames()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	for _, u := range usernames {
		if user.Username == u {
			return echo.NewHTTPError(http.StatusConflict, "this username is already used")
		}
	}

	userID, err := h.accountService.CreateUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	jwt, err := h.accountService.CreateJWT(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	cookie := h.createCookie(jwt)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "register succeeded")
}

func (h *AccountHandler) HandleLogin(c echo.Context) error {
	var user models.UserAuthRequest
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization information")
	}

	userID, err := h.accountService.Login(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	jwt, err := h.accountService.CreateJWT(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	cookie := h.createCookie(jwt)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "login succeeded")
}

func (h *AccountHandler) HandleLogout(c echo.Context) error {
	cookie := h.accountService.CreateBeingDeletedCookie()
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "logout succeeded")
}

func (h *AccountHandler) HandleDelete(c echo.Context) error {
	userID := c.Get("userID").(int64)

	if err := h.accountService.DeleteUser(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := h.accountService.CreateBeingDeletedCookie()
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "account deleted")
}

func (h *AccountHandler) createCookie(jwt string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = jwt
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(time.Hour * 24 * 7)
	//cookie.SameSite = http.SameSiteNoneMode

	return cookie
}
