package user_service_test

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	userService "../user.service"
	m "../../models"
	"time"
)

var userId string

func TestCreate(t *testing.T) {

	oid := primitive.NewObjectID()
	userId = oid.Hex()

	user := m.User{
		ID: oid,
		Name: "Ygor",
		Email: "ygordesa2010@hotmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := userService.Create(user)

	if err != nil {
		t.Error("error")
		t.Fail()
	} else {
		t.Log("Sucess status 200")
	}
}

func TestRead(t *testing.T) {

	users, err := userService.Read()

	if err != nil {
		t.Error("error to seach user")
		t.Fail()
	}

	if len(users) == 0 {
		t.Error("return data")
		t.Fail()
	} else {
		t.Log("sucess to finesh the test")
	}

}

func TestUpdate(t *testing.T) {

	user := m.User{
		Name: "Ygor",
		Email: "ygordesa2010@hotmail.com",
	}

	err := userService.Update(user, userId)

	if err != nil {
		t.Error("Error trying to update user")
		t.Fail()
	} else {
		t.Log("The update test ended successfully!")
	}
}

func TestDelete(t *testing.T) {

	err := userService.Delete(userId)

	if err != nil {
		t.Error("Error trying to delete user")
		t.Fail()
	} else {
		t.Log("The elimination test ended successfully!")
	}
}
