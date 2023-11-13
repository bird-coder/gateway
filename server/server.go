package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Start() {
	g := gin.New()
	s := &http.Server{
		Addr:    "",
		Handler: g,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			fmt.Println("")
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		fmt.Println("")
	}
}

func Stop() {

}
