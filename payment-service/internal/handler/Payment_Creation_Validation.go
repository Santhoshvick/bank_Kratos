package handler

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)


func PaymentMethod(paymentMethod string) error{
	method:=[]string{"Credit Card","Debit Card","UPI","NEFT"}

	for i:=0;i<len(method);i++{
		val:=method[i]
		if(val==paymentMethod){
			return nil
		}
	}
	return errors.New("payment method is not valid ,Please Provide a correct Payment Method")
}


func PaymentType(paymentType string) error{
	type1:=[]string{"Direct","Other"}

	for i:=0;i<len(type1);i++{
		val:=type1[i]
		if(val==paymentType){
			return nil
		}
	}
	return errors.New("invalid Payment Type Only Direct and Other types is acceptable")
}

func CurrencyValidation(currency string) error{
	arr:=[]string{"RUPEE","USD","EURO","DIRHAM"}

	for i:=0;i<len(arr);i++{
		val:=arr[i]
		if val==currency{
			return nil
		}
	}
	 return errors.New("please verify the currency Once.Allowed Curencies are USD  EURO  RUPEE   DIRHAM")
}

func AccountNumberValidation(accountNumber int64) error{
	val:=string(accountNumber)
	length:=len(val)
	if(length!=10){
		return errors.New("please enter a Valid Account Number the account Number digit should be of 10 digits")
	}

	return nil

}

func GenerateRefNumber() string {
	rand.Seed(time.Now().UnixNano()) 
	refNum := rand.Intn(900000) + 100000 
	return fmt.Sprintf("%06d", refNum)
}
// func AccountStatusValidation(){

// }


