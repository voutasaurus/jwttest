package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	errUnexpectedSigningMethod = errors.New("unexpected signing method")
)

func main() {
	x := time.Now().UnixNano()
	fmt.Println("before:", x)
	preToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"x": x,
	})
	k := []byte("s")
	token, err := preToken.SignedString(k)
	if err != nil {
		log.Fatal(err)
	}
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errUnexpectedSigningMethod
		}
		return k, nil
	})
	if err != nil {
		log.Fatal(err)
	}
	claimsMap, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		log.Fatal("bad claims type")
	}
	xIf := claimsMap["x"]

	switch xx := xIf.(type) {
	case int64:
		fmt.Println("cool:", xx)
	case float64:
		fmt.Println("jwt no. no jwt. stop. STAHP:", xx)
	}
}
