package cmd

import (
	"fmt"

	"github.com/sanjaysingh/aad-jwt-gen/services"
)

func RunHeadless(clientId, tenantId, clientSecret, scope string) {
	token, err := services.GenerateToken(clientId, tenantId, clientSecret, scope)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("Generated JWT Token:", token)
}
