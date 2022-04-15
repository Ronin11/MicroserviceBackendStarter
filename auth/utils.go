package auth

import (
	"fmt"
	"os"
	"context"
	
	"github.com/Nerzal/gocloak/v11"
	"github.com/golang-jwt/jwt/v4"

	"nateashby.com/gofun/logging"
)

type AuthClaims struct {
	Id string `json:"sub"`
}

type AdminCreds struct {
	user string
	password string
	token *gocloak.JWT
	realm string
	clientId string
	secretId string
}

type AuthHandler struct {
	keycloakUrl string
	adminCreds *AdminCreds
	realm string
	clientId string
	secretId string
	client gocloak.GoCloak
}

var authHandler *AuthHandler

func initialize(
	adminCreds *AdminCreds,
	keycloakUrl string,
	realm string,
	clientId string,
	secretId string,
) (*AuthHandler) {

	authHandler = &AuthHandler{
		adminCreds: adminCreds,
		keycloakUrl: keycloakUrl,
		realm: realm,
		clientId: clientId,
		secretId: secretId,
	}
	client := gocloak.NewClient(keycloakUrl)
	authHandler.client = client
	ctx := context.Background()
	token, err := client.Login(ctx, authHandler.adminCreds.clientId, authHandler.adminCreds.secretId, authHandler.adminCreds.realm, authHandler.adminCreds.user, authHandler.adminCreds.password)
	authHandler.adminCreds.token = token
	if err != nil {
		logging.Log("Keycloak connection error: ", err)
	}else{
		logging.Log("Keycloak connection established")
	}
	return authHandler
}

func GetAuthHandlerInstance() (*AuthHandler) {
	if authHandler != nil {
		return authHandler
	}

	fmt.Println("getHandlerInstance")

	adminCreds := &AdminCreds{
		user: os.Getenv("KEYCLOAK_ADMIN_USER"),
		password: os.Getenv("KEYCLOAK_ADMIN_PASSWORD"),
		realm: os.Getenv("KEYCLOAK_ADMIN_REALM_NAME"),
		clientId: os.Getenv("KEYCLOAK_ADMIN_CLIENT_ID"),
		secretId: os.Getenv("KEYCLOAK_ADMIN_SECRET_ID"),
	}

	return initialize(
		adminCreds,
		os.Getenv("KEYCLOAK_URL"),
		os.Getenv("KEYCLOAK_REALM_NAME"),
		os.Getenv("KEYCLOAK_CLIENT_ID"),
		os.Getenv("KEYCLOAK_SECRET_ID"),
	)
}

func (ah *AuthHandler) GetUserFromToken(tokenString string) *User{
	ctx := context.Background()
	rptResult, err := ah.client.RetrospectToken(ctx, tokenString, ah.clientId, ah.secretId, ah.realm)
	if err != nil {
		logging.Log("Inspection failed:"+ err.Error())
	 	return nil
	}
   
	if !*rptResult.Active {
		logging.Log("Token is not active")
		return nil
	}

	token, _, err := ah.client.DecodeAccessToken(ctx, tokenString, ah.realm)
	if err != nil {
		logging.Log("Failed to decode token")
		return nil
	}

	fmt.Println("WUT: ", token, token.Claims)
	fmt.Println(token.Claims.(*jwt.MapClaims))
	// if claims, ok := *token.Claims.(jwt.MapClaims); ok {
	// 	fmt.Println("CLAIMS: ", claims)
	// 	// return claims, nil
	// }
	// asdf := claims.sub
	// claims2 := claims
	// fmt.Println(make(map[string]claims))
	// fmt.Println(&token.Claims["id"])
	// fmt.Println(claims)
	// id := token.Claims.
	// fmt.Println("ID: ", id)
	// claimsMap := *claims
	// // fmt.Println("STUFF: ", claimsMap["id"].(string))
	// fmt.Println("RESULT: ", claimsMap["sub"].(string))
	return &User{Id: 0}
}

func (ah *AuthHandler) login(user string, passhash string) (string, error) {
	ctx := context.Background()
	token, err := ah.client.Login(ctx, ah.clientId, ah.secretId, ah.realm, user, passhash)
	if err != nil {
		logging.Log("Login failed:"+ err.Error())
		return "", err
	}
	return token.AccessToken, nil
}

func (ah *AuthHandler) createUser(username string, passhash string) (string, error) {
	user := gocloak.User{
		// FirstName: gocloak.StringP("Bob"),
		// LastName:  gocloak.StringP("Uncle"),
		// Email:     gocloak.StringP("something@really.wrong"),
		Enabled:   gocloak.BoolP(true),
		Username:  gocloak.StringP(username),
	}
	  
	ctx := context.Background()
	userId, err := ah.client.CreateUser(ctx, ah.adminCreds.token.AccessToken, ah.realm, user)
	if err != nil {
		logging.Log("Create User failed:"+ err.Error())
		return "", err
	}
	err = ah.client.SetPassword(ctx, ah.adminCreds.token.AccessToken, userId, ah.realm, passhash, false)
	if err != nil {
		logging.Log("Password Set failed:"+ err.Error())
		return "", err
	}
	return ah.login(username, passhash)
}


