package infrastructure

import (
	"log"
	"strings"
	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func FirebaseAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		creds, err := google.FindDefaultCredentials(ctx)
		if err != nil {
			log.Fatalf("error")
		}
		opt := option.WithCredentials(creds)
		app, err := firebase.NewApp(ctx, nil, opt)
		auth, err := app.Auth(ctx)
		authHeader := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := auth.VerifyIDToken(ctx, idToken)
		log.Printf("token: %v\n", token)
	
		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}