import React from 'react'
import { useState } from 'react';
import img1 from '../assets/stitch.svg'
import { Link, redirect } from 'react-router-dom';
import '../styles/Payment.css'
import img9 from '../assets/img9.jpg'
import axios from 'axios';


const Payment = () => {
    const [newDebit,setDebitAmount]=useState(0);
    const [newAmount, setNewAmount] = useState(0);
    const [originalAmount,setOriginalAmount]=useState(0);

    
    const [accountData,setAccountData]=useState({
        accountId:'',
        accountNumber:'',
        accountType:'',
        currency:'',
        status:'',
        availableBalance:'',
    })

    const [accountData2,setAccountData1]=useState({
        accountId:'22',
        accountNumber:'',
        accountType:'',
        currency:'',
        status:'',
        availableBalance:'',
    })
    const [transferAmount,setTransferAmount]=useState({
        transactionId:'',
        fromAccountId:accountData.accountNumber,
        toAccountId:'',
        currency:'',
        paymentType:'Credit',
        status:'',
    }
  )
  const [amt, setAmount] = useState({
  amount: '',
});
     
    //from account details
    transferAmount.transactionId=2001
    transferAmount.fromAccountId=accountData.accountNumber
    transferAmount.amount=String(newDebit)
    transferAmount.currency=String(accountData.currency)
    transferAmount.status=String(accountData.status)

    console.log(transferAmount)
    const transaction=async(e)=>{
        try{
    accountData2.accountNumber=transferAmount.toAccountId
    accountData2.availableBalance=Number(accountData2.availableBalance)+Number(transferAmount.amount)
    let val1=String(accountData2.availableBalance)
    accountData2.availableBalance=val1
    console.log("The new2 avaliable balance",accountData2.availableBalance)
   console.log(accountData2)
    accountData.availableBalance=Number(accountData.availableBalance)-Number(transferAmount.amount)
    let val=String(accountData.availableBalance)
    accountData.availableBalance=val
    console.log("The new Balance is ",accountData.availableBalance)
        const response=await axios.post('http://localhost:8005/v1/transaction/create',transferAmount)
        const respose1=await axios.post('http://localhost:8004/v1/payment/create',transferAmount)      
        console.log("payment is updated successfully")
        setTransferAmount(respose1.data)
        alert('payment is trasnfered successfully')
        setTransferAmount(response.data)

        const response3= await axios.put('http://localhost:8002/v1/account/update',accountData)
            console.log(accountData)
            setAccountData(response3.data)

        const response4= await axios.put('http://localhost:8002/v1/account/update',accountData2)
            console.log(accountData2)
            setAccountData1(response4.data)
        }
        
        catch(error){
            console.log(error)
            alert("Please Verify the port")
        }
    }

    const handleFun= async(e)=>{
        try{const response3= await axios.get('http://localhost:8002/v1/account/find/21',accountData)
            console.log(accountData)
            setAccountData(response3.data)
            alert('retrived the transaction history')
        }
        catch(error){
            console.log(error)
            alert("Please Verify the port")
        }
    }

    const fee = async (e) => {
    let amt = parseFloat(document.getElementById('amount').value);
    setOriginalAmount(amt)
    let fees = 0;
    if (isNaN(amt)) {
      alert("Please enter a valid number");
      return;
    }
    if (amt < 10000) {
      fees = (amt * 0.5) / 100;
    } else if (amt > 10000 && amt < 50000) {
      fees = (amt * 0.9) / 100;
    }
    setNewAmount(fees);
    console.log("The Fee is calculated successfully");
    console.log(fees)
    let val=fees+amt
    setDebitAmount(val)
  };
  const handleChange = (e) => {
  const { name, value } = e.target;
  let alphabets=value
  if(name==='toAccountId'||name==='amount'){
     alphabets=value.replace(/[^0-9]/g,'')
  }
  setAmount((prevData)=>({
    ...prevData,
    [name]: alphabets
  }));
  setTransferAmount((prevData) => ({
    ...prevData,
    [name]: alphabets
  }));
};


  function show(){
    let doc=document.getElementById('payment-det')
    doc.style.display='inline-block'
    console.log("function is working properly")
  }
  return (
    <>
    <nav className="navbar">
    <img src={img1} alt="Stitch Logo" className="logo" />
    <input type='text' placeholder='Search'  className='search-field' />
    <div className="nav-links">
    <Link className='link' to="/Home">Home</Link>
    <Link className='link' to="/account">Account</Link>
    <Link className='link' to="/transaction">Transaction</Link>
    <Link className='link' to="/payment">Payment</Link>
    <Link className='link' to="/card">Card</Link>
    <Link className='link' to="/signin">Signin</Link>
  </div>
</nav>
<img src={img9} style={{height:'500px',width:'100%',position:'absolute'}} alt="" srcset="" />

<div className="payment-body">
    <h3 style={{display:'inline-block',fontSize:'30px',marginLeft:'30px', position:'relative'}}>Payments</h3>
    <button id='make-payment' style={{position:'relative'}} onClick={show}>Make Payment</button>
</div>

<div className="payment-det" id='payment-det' style={{position:'relative',marginTop:'450px'}}>

    <h3 style={{color:'red'}}>Payment Wizards</h3>
    <div className="from-account" style={{marginLeft:'10px'}}>
        <button onClick={handleFun}>generate</button>
        <h2 style={{marginTop:'-20px'}}>From-Account</h2>
        {/* <button onClick={updateBalance}>Refresh</button> */}
        <div className="user-account">
        <input type="radio" name="choice" id="" value={'CHK-001234 - Checking'} style={{display:'inline'}} />
        <h4 style={{display:'inline-block'}}>{accountData.accountNumber} - Checking</h4>
        <h4 style={{display:'inline-block'}}>Avaliable Balance:${accountData.availableBalance} </h4>
        </div>
        <div className="user-account">
        <input type="radio" name="choice" id="" value={'SAV-005678 - Savings '} style={{display:'inline'}} />
        <h4 style={{display:'inline-block'}}>SAV-005678 - Savings </h4>
        <h4 style={{display:'inline-block'}}>Avaliable Balance:$12,450.75 </h4>
        </div>
        <div className="user-account">
        <input type="radio" name="choice" id="" value={'CRD-112233 - Credit Card'} style={{display:'inline'}} />
        <h4 style={{display:'inline-block'}}>CRD-112233 - Credit Card</h4>
        <h4 style={{display:'inline-block'}}>Avaliable Balance:$12,450.75 </h4>
        </div>
    </div>

        <div className="to-account" style={{marginTop:'20px',height:'100px',marginLeft:'10px'}}>
            <h2 style={{display:'block'}} >To Account/Beneficiary:</h2>
            <label>Account Number:</label>
            <input type="text" name="toAccountId" id="" value={transferAmount.toAccountId} onChange={handleChange}   inputmode="numeric"
  pattern="\d{16}"
  maxlength="16"
  placeholder="Enter 16-digit account number"/>
            <input type="radio" className="radio" name='choice' /> Internal Account
            <input type="radio" name="choice" id="" /> External Account
            <label style={{marginLeft:'10px'}}>Beneficiary Name:</label>
            <input type="text" name="" id="" />
            {/* <label htmlFor="">Bank Name:</label>
            <input type="text" name="" id="" /> */}
        </div>
        
        <br />

                  <span style={{marginLeft:'200px'}}>Amount:$</span><input type="text" name="amount" id="amount" onChange={handleChange} value={amt.amount}  style={{borderTop:'none',borderLeft:'none',borderRight:'none'}} inputmode="numeric"
  pattern="\d{10}"
  maxlength="10"/>USD
        <button onClick={transaction} style={{display:'inline-block',marginTop:'0px'}}>transfer</button>
        <button className="fee-calculation" style={{width:'200px',marginTop:'-90px',display:'inline-block'}} onClick={fee}>Calculate Fees</button>
        <div className="final-amount" style={{marginLeft:'20px'}}>
            <h3 > Transfer Amount:$      {originalAmount}</h3>
            <h3>Processing Fee:$     {newAmount}</h3>
            <h3>Total Debit:$        {newDebit}</h3>
        </div>
</div>
<div>
</div>
    </>
  )
}

export default Payment
