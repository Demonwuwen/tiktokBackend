package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestBcryptingIsEasy(t *testing.T) {
	pass := []byte("mypassword")
	hp, err := bcrypt.GenerateFromPassword(pass, 0)
	fmt.Println("mypassword hp = ", string(hp))
	if err != nil {
		t.Fatalf("GenerateFromPassword error: %s", err)
	}

	if bcrypt.CompareHashAndPassword(hp, pass) != nil {
		t.Errorf("%v should hash %s correctly", hp, pass)
	}

	notPass := "notthepass"
	err = bcrypt.CompareHashAndPassword(hp, []byte(notPass))
	if err != bcrypt.ErrMismatchedHashAndPassword {
		t.Errorf("%v and %s should be mismatched", hp, notPass)
	}
}

func TestPasswordHash(t *testing.T) {
	ps1, _ := GenPasswordHash("ilooveyou")
	ps2, _ := GenPasswordHash("ilooveyou")
	ps3, _ := GenPasswordHash("ilooveyou")

	fmt.Println(ps1)
	fmt.Println(ps2)
	fmt.Println(ps3)
	//	$2a$10$zQ5tvER8QewyWriaUzUVUuBksyTgLYaO0Tb1oDuXduPreF/2iu.6e
	//$2a$10$icdiMSDGWaPm7XVkkxffce3P7rHIUcPP7YMQDzqrUHwWYBvIasoSy
	//$2a$10$B5EoXD2OWa58gbilJuClYObM2VufMgEUddzcWH0ugI0qCo1gyiqV6
}

func TestPasswordVerify(t *testing.T) {
	r := PasswordVerify("ilooveyou", "$2a$10$zQ5tvER8QewyWriaUzUVUuBksyTgLYaO0Tb1oDuXduPreF/2iu.6e")
	fmt.Println(r)
	r1 := PasswordVerify("iloveyou", "$2a$10$icdiMSDGWaPm7XVkkxffce3P7rHIUcPP7YMQDzqrUHwWYBvIasoSy")
	fmt.Println(r1)
	r2 := PasswordVerify("ilooveyou", "$2a$10$B5EoXD2OWa58gbilJuClYObM2VufMgEUddzcWH0ugI0qCo1gyiqV6")
	fmt.Println(r2)
}

func TestRandNameString(t *testing.T) {
	str := RandNameString(12)
	fmt.Println(str)
}
