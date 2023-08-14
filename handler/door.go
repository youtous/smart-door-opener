package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)
import "github.com/go-resty/resty/v2"

type Door struct {
	secretCode            string
	iftttServerName       string
	iftttServerKey        string
	iftttWebhookEventName string
	httpClient *resty.Client
}

func (h *Door) GetOpenDoor(c echo.Context) error {
	if c.Param("accessCode") != h.secretCode {
		return c.Render(http.StatusForbidden, "wrongCode", echo.Map{
			"title": "Bienvenue à Saint-Aubin ! Arrivée autonome",
		})
	}

	return c.Render(http.StatusOK, "index", echo.Map{
		"title": "Bienvenue à Saint-Aubin ! Arrivée autonome",
		"csrfToken": c.Get(middleware.DefaultCSRFConfig.ContextKey).(string),
	})
}

func (h *Door) PostOpenDoor(c echo.Context) error {
	if c.Param("accessCode") != h.secretCode {
		return c.Render(http.StatusForbidden, "wrongCode", echo.Map{
			"title": "Bienvenue à Saint-Aubin ! Arrivée autonome - Mauvais code",
		})
	}

	// open the door
	resp, respErr := h.httpClient.R().Post(fmt.Sprintf("/trigger/%s/with/key/%s", h.iftttWebhookEventName, h.iftttServerKey))

	successMessage := ""
	errorMessage := ""
	if respErr != nil {
		c.Logger().Error("Received an error while triggering IFTTT event: %v", respErr)
		errorMessage = "Une erreur est survenue lors de l'ouverture, veuillez nous contacter."
	} else {
		c.Logger().Info(fmt.Sprintf("Door opening trigger accepted: %v", resp.Body))
		successMessage = "La commande d'ouverture de porte a été exécutée, elle devrait s'effectuer dans quelques secondes, poussez la porte."
	}

	return c.Render(http.StatusOK, "index", echo.Map{
		"title": "Bienvenue à Saint-Aubin ! Arrivée autonome - Ouverture de porte",
		"csrfToken": c.Get(middleware.DefaultCSRFConfig.ContextKey).(string),
		"errorMessage": errorMessage,
		"successMessage": successMessage,
	})
}

func NewHandlerDoor(secretCode string, iftttServerName string, iftttServerKey string, iftttWebhookEventName string) *Door {
	doorHandler := &Door{
		secretCode:            secretCode,
		iftttServerName:       iftttServerName,
		iftttServerKey:        iftttServerKey,
		iftttWebhookEventName: iftttWebhookEventName,
		httpClient: resty.New(),
	}
	doorHandler.httpClient.SetBaseURL(fmt.Sprintf("https://%s", doorHandler.iftttServerName))
	doorHandler.httpClient.SetHeader("Content-Type", "application/json")
	doorHandler.httpClient.SetHeader("Accept", "application/json")

	return doorHandler
}
