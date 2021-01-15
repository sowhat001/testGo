package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

// User contains user information
type User struct {
	Age       uint8         `validate:"gte=0,lte=130"`
	Email     string        `validate:"required,email"`
	Addresses []interface{} `validate:"dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	Planet string `validate:"required"`
}

func ValidateStruct() {
	var validate *validator.Validate
	validate = validator.New()
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
	}
	_ = address
	user := &User{
		Age:       15,
		Email:     "Badger.Smith@gmail.com",
		Addresses: []interface{}{1},
	}
	user = nil
	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(user)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println("here1, ", err)
			return
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
		// from here you can create your own error messages in whatever language you wish
		return
	}
	// save user to database
}

func ValidateVariable() {
	var validate *validator.Validate
	validate = validator.New()
	myEmail := "joeybloggs.gmail.com"
	errs := validate.Var(myEmail, "required,email")
	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}
	// email ok, move on
}
