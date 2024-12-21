package service

import (
	"fmt"
	"log/slog"
	"pscan/web/router"
)

func RunService(logger *slog.Logger) {
	r := router.InitRouters()
	fmt.Println("[+] Please visit http://localhost:8989/")
	err := r.Run("127.0.0.1:8989")
	if err != nil {
		logger.Error(err.Error())
		return
	}
}
