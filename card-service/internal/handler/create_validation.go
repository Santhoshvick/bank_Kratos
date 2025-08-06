package handler

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func ValidateAccountNumber(accountNumber int64)(string,error){
     accountvalidate:=string(accountNumber)
	 if len(accountvalidate)!=10{
		return "Invalid Account Number",nil
	 }
	 return accountvalidate,nil
}

func RandomGenerator() string{
	rand.Seed(time.Now().UnixNano()) 
	refNum := rand.Intn(9000000000000000) + 1000000000000000 
	return fmt.Sprintf("%06d", refNum)
}

func CardType(cardType string) error{
	arr:=[]string{"Credit Card","Debit Card","Prepaid Card","Travel Card","Reward Card"}

	if cardType==""{
		return errors.New("card type field is empty Please enter the card type you wish to use We offer 1.Credit Card \n 2.Debit Card \n 3.Prepaid Card \n 4.Travel Card \n 5.Reward card")
	}
	for i:=0;i<len(arr);i++{
		val1:=arr[i]
		if cardType==val1{
			return nil
		}
	}
	return errors.New("card type should be only of 1.Credit Card  2.Debit Card \n 3.Prepaid Card \n 4.Travel Card \n 5.Reward card" )
}


func CardStatus(cardStatus string) error{
	   arr:=[]string{"Approved","Declined","Pending"}

	   if cardStatus==""{
		return errors.New("Card Status is Empty Please ensure field once")
	   }
	   
	   for i:=0;i<len(arr);i++{
		val:=arr[i]
		if val==cardStatus{
			return nil
		}
	   }
	   return errors.New("Please Ensure You Provide correct Card Status. Pending Approved Declined")
}

