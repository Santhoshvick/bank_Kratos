package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	// "strconv"

	"math/rand"
	pb "message-service/api/helloworld/v1"
	"message-service/internal/biz"

	"github.com/go-redis/redis"
	"github.com/nats-io/nats.go"
	gomail "gopkg.in/mail.v2"
)

// GreeterService is a greeter service.
type NotificationService struct {
	pb.UnimplementedNotificationServer
	uc *biz.NotificationUsecase
}

type Message1 struct {
	CustomerNumber string `json:"CustomerNumber"`
	FirstName      string `json:"FirstName"`
	Email          string `json:"Email"`
	Status         string `json:"Status"`
}

type PaymentStatus struct{
	FromEmail string `json:"FromEmail"`
	ToEmail string `json:"ToEmail"`
	FromBalance string `json:"FromBalance"`
	ToBalance string  `json:"ToBalance"`
	FromFirstName string `json:"FromFirstName"`
	ToFirstName string `json:"ToFirstName"`
	Amount string `json:"Amount"`
	FromAccountNumber int64  `json:"AccountNumber"`
	ToAccountNumber int64 `json:"To AccountNumber"`
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.NotificationUsecase) *NotificationService {
	return &NotificationService{uc: uc}
}
func connectToNATS() (*nats.Conn, error) {
	fmt.Println("Nats is connected successfully")
    return nats.Connect(nats.DefaultURL)
}
func (s *NotificationService) CreateNotification(ctx context.Context, req *pb.NotificationRequest) (*pb.NotificationResponse, error) {
	
	otp:=generate()
	val2:=strconv.Itoa(otp)

	// rdb:=redis.NewClient(&redis.Options{
	// 	Addr:"localhost:6379",
	// 	DB:0,
	// })

	// nc,_:=connectToNATS()
	// nc.Subscribe("Create",func(msg *nats.Msg) {
	// 	var data Message1
	// 	if err := json.Unmarshal(msg.Data, &data); err != nil {
	// 		fmt.Println("Failed to unmarshal NATS message:", err)
	// 		return
	// 	}
	
	// 	fmt.Println("Received customer data for:", data.Email)
	// 	var timeout=100*time.Second

	//    result,err:=rdb.Set("Message","Dear "+data.FirstName +"as per your Request We have created  a Sepearate  customer id for You.\n you can enjoy our bank service.",timeout).Result()

	// 	if err!=nil{
	// 		log.Fatalf("There is issue with creating the key to store the data in the redis %v",err)
	// 		return
	// 	}
	// 	fmt.Println(result)

		message := gomail.NewMessage()
		message.SetHeader("From", "ssanthoshvicky2003@gmail.com")
		message.SetHeader("To", "ssanthoshvicky2003@gmail.com")
		message.SetHeader("Subject", "Hello from the Stitch team")
		// message.SetBody("text/plain", "Your OTP is: "+val2)
		message.SetBody("text/plain","Dear "+"\nas per your Request We have created  a Sepearate  customer id for You.\n you can enjoy our bank service."+val2)
		dialer := gomail.NewDialer("smtp.gmail.com", 587, "ssanthoshvicky2003@gmail.com", "lopb awgc uncw mkwb")

		if err := dialer.DialAndSend(message); err != nil {
			fmt.Println("Failed to send email:", err)
			return nil,err
		}
		fmt.Println("Email sent to",req.Email)	
	// })
	g, err := s.uc.CreateNotification(ctx,&biz.Notification{
		Email: req.Email,
	})
	
	if err != nil {
		return nil, err
	}
	fmt.Println(g)
	return &pb.NotificationResponse{ 
		Otp:int64(otp),
		}, 
		nil
}


func generate()int{
	  return 100000 + rand.Intn(900000)
}

func (s *NotificationService) CreateTransactionNotification(ctx context.Context,req *pb.TransactionRequest) (*pb.TransactionResponse,error){
	g,err:=s.uc.CreateNotification(ctx,&biz.Notification{
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(g);

	rdb:=redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		DB:0,
	})

	nc,_:=connectToNATS()
	nc.Subscribe("Real",func(msg *nats.Msg) {
		var data1 PaymentStatus
		if err := json.Unmarshal(msg.Data, &data1); err != nil {
			fmt.Println("Failed to unmarshal NATS message:", err)
			return
		}
		fmt.Println(data1.FromBalance)
		fmt.Println(data1.ToBalance)
		message := gomail.NewMessage()
		message.SetHeader("From", "ssanthoshvicky2003@gmail.com")
		var timeout=200*time.Second
	   result,err:=rdb.Set("Message3","Dear "+data1.FromFirstName+"\n We’re writing to confirm that your recent transaction has been successfully processed.\n"+"Amount Transfered:"+data1.Amount+"\n To Account Number:If you have any questions, feel free to reply to this email or reach us at Support Contact Info stitchsupport@gmail.com."+"\n AccountBalance"+string(data1.FromBalance),timeout).Result()
	   if err!=nil{
		 log.Fatalf("There is an issue with creating the key %v",err)
	   }
	   fmt.Println(result)
    
		message.SetHeader("To", data1.FromEmail)
		message.SetHeader("Subject", "Transaction Confirmation"+data1.Amount+"Successfully Processed")
		// message.SetBody("text/plain", "Your OTP is: "+val2)
		value,err:=rdb.Get("Message3").Result()
		if err!=nil{
			log.Fatalf("Key not found")
		}
		message.SetBody("text/plain",string(value))

		dialer := gomail.NewDialer("smtp.gmail.com", 587, "ssanthoshvicky2003@gmail.com", "lopb awgc uncw mkwb")

		if err := dialer.DialAndSend(message); err != nil {
			fmt.Println("Failed to send email:", err)
			return
		}
		fmt.Println("Email sent to", data1.FromEmail)	
	})


		//To Account Notification Email
		nc.Subscribe("Real",func(msg *nats.Msg) {
			var data1 PaymentStatus
		if err := json.Unmarshal(msg.Data, &data1); err != nil {
			fmt.Println("Failed to unmarshal NATS message:", err)
			return
		}

		fmt.Println("ToBalance",data1.ToBalance)
		fmt.Println("To Account",data1.ToFirstName)
		message1 := gomail.NewMessage()
		message1.SetHeader("From", "ssanthoshvicky2003@gmail.com")
		var timeout1=200*time.Second
	   result1,err:=rdb.Set("Message4","Dear "+data1.ToFirstName+"Amount:"+data1.Amount+"\n has been credited successfully in your account \n If you have any questions, feel free to reply to this email or reach us at Support Contact Info stitchsupport@gmail.com."+"\n AccountBalance"+data1.ToBalance,timeout1).Result()
	   if err!=nil{
		 log.Fatalf("There is an issue with creating the key %v",err)
	   }
	   fmt.Println(result1)
    
		message1.SetHeader("To", data1.ToEmail)
		message1.SetHeader("Subject", "Transaction  Confirmation "+data1.Amount+"Credited Successfully in your Account")
		// message.SetBody("text/plain", "Your OTP is: "+val2)
		value2,err:=rdb.Get("Message4").Result()
		if err!=nil{
			log.Fatalf("Key not found")
		}
		message1.SetBody("text/plain",string(value2))

		dialer1:= gomail.NewDialer("smtp.gmail.com", 587, "ssanthoshvicky2003@gmail.com", "lopb awgc uncw mkwb")

		if err := dialer1.DialAndSend(message1); err != nil {
			fmt.Println("Failed to send email:", err)
			return
		}
		fmt.Println("Email sent to", data1.ToEmail)

		result2,err:=rdb.Get("Payment1").Result()
		if err!=nil{
			log.Fatalf("Payment key not found %v",err)
		}
	    fmt.Println(result2)

	})

	for i:=1;i<20;i++{
		result3,err:=rdb.Get(fmt.Sprintf("Customer%d",i)).Result()
		if err==redis.Nil{
			continue
		}

		var customer Message1

		err = json.Unmarshal([]byte(result3), &customer)
		if err!=nil{
			log.Fatalf("Cannot unmarshal the customer data %v",err)
		}
	    
		if(customer.CustomerNumber=="90212411"){
			fmt.Println(customer)
			break;
		}
	}
	fmt.Println()
	return &pb.TransactionResponse{ 
		Message1: "Transaction is done successfully,Please Verify your email once",
		}, 
		nil
}






