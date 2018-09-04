package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Azure AD Pod Identity + Keyvault Sample")
}

func getKeyvaultSecret(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("Keyvault Secret value is: %s", "TODO"))

	keyvaultName := os.Getenv("AZURE_KEYVAULT_NAME")
	keyvaultSecretName := os.Getenv("AZURE_KEYVAULT_SECRET_NAME")
	keyvaultSecretVersion := os.Getenv("AZURE_KEYVAULT_SECRET_VERSION")

	keyClient := keyvault.New()
	authorizer, err := auth.NewAuthorizerFromEnvironment()

	if err == nil {
		keyClient.Authorizer = authorizer
	}

	secret, err := keyClient.GetSecret(context.Background(), fmt.Sprintf("https://%s.vault.azure.com", keyvaultName), keyvaultSecretName, keyvaultSecretVersion)
	if err != nil {
		log.Printf("failed to retrieve the Keyvault secret: %v", err)
		return
	}

	io.WriteString(w, fmt.Sprintf("The value of the Keyvault secret is: %v", secret.Value))
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/keyvault", getKeyvaultSecret)
	log.Println("http server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
