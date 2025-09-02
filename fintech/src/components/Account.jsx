import React from 'react'
import img1 from '../assets/stitch.svg'
import '../styles/Account.css'
import '../styles/Home.css'
import { useRef } from 'react';
import { data, Link } from 'react-router-dom';
import { useState } from 'react';
import axios from 'axios';
import Home from '../components/home.jsx'

const transactionshistory = [
  {
    date: '12:45:32',
    type: 'Card Payment',
    amount: "-$45.99",
    status: 'POSTED',
    Reference:'TXN-4521456'
  },
  {
     date: '12:45:32',
     type: 'ATM Withdrawal',
    amount: "$-100.0",
    status: 'PENDING',
    Reference:'TXN-4521456'
  },
  {
    date: '12:43:01',
    type: 'Online Transfer',
    amount: "$500.0",
    status: 'POSTED',
    Reference:'TXN-4521456'
  },
  {
    date: '12:42:33',
    type: 'Fee Charge',
    amount: "$2.5",
    status: 'POSTED',
    Reference:'TXN-4521456'
  }
];

let val=1;

const Account = () => {
    const [dataVal, setFormData] = useState({
    // accountId:'',
    customerId:'',
    accountNumber: '',
    accountType: '',
    currency: '',
    availableBalance: '',
    pendingBalance: '',
    creditLimit: '',
    status: ''
      });
      const [dataVal1,setCustomerData]=useState({
        customerId:'',
        customerName:'',
      })

const count = useRef(0);
    const handleFind= async(e)=>{
        e.preventDefault();
        try{
          console.log(count)
          if(count.current===0){
            const response= await axios.get('http://localhost:8002/v1/account/find/8',dataVal)
            const response2= await axios.get('http://localhost:8000/v1/find/18',dataVal1)
            
            console.log("Account Details:",response.data)
            const accountdetails=response.data
            const customerDetails=response2.data
            setFormData(accountdetails)
            let val1=document.getElementById('pendingBalance')
            let val2=document.getElementById('availableBalance')
            let val3=document.getElementById('creditLimit')
            console.log(customerDetails.customerId)
            if (String(accountdetails.customerId) === String(customerDetails.customerId)){
              count.current++
            alert('Account is created successfully')
           let acc=document.getElementById('account-det1')
           let val4=document.getElementsByClassName('account-det')
             acc.style.display = "inline";
             val1.style.display="inline"
             val2.style.display="inline"
             val3.style.display="inline"
            for (let el of val4) 
              {
                   el.style.display = "inline";
              }
            }
            else (
              alert('account does not exist')
            )
        }
          else if(count.current>0){
            alert('account is already created')
          }
        }
        catch(error){
            console.log("There is an issue with submitting the form data")
            alert('Please verify the api once',error)
        }
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
    <Link className='link' to="/Signin">Signin</Link>

  </div>
</nav>
   <div className="nav">
       <h1 style={{fontSize:'25px',marginTop:'-10px'}} className='head1'>Account Details</h1>
      <button className='create-button' onClick={handleFind}>+Create Account</button>
       </div>

    <div className="main">
        <h2 className='details'  style={{display:'none'}}>Account Number:{dataVal.accountNumber}</h2>
        <h2 className='details' style={{display:'none'}}>status:{dataVal.status}</h2>
        <h2 className='details' style={{display:'none'}}>Customer Id:{dataVal.customerId}</h2>
        {/* <h2 className='details'>Type:{dataval.T}</h2> */}
    </div>

    <iframe src="http://localhost:3000/public/question/5ccd4126-67f4-40d2-aeb5-9fce8a3ae57d" width="100%" height="400px" frameborder="0"></iframe>
    
    <div className='dashboard' style={{marginLeft:'9px',width:'500px'}}>
    <h1  >Avaliable Balance</h1>
    <h2 id='availableBalance' style={{display:'none'}}>${dataVal.availableBalance}</h2>
    <h2>+2.3% from last</h2>
</div>

<div className='dashboard'>
    <h1>Pending Balance</h1>
    <h2 id='pendingBalance' style={{display:'none'}}>${dataVal.pendingBalance}</h2>
    <h2>2 Pending</h2>
</div>
<div className='dashboard1' style={{width:'450px'}}>
    <h1>Credit Limit</h1>
   <h2 id='creditLimit' style={{display:'none'}}>${dataVal.creditLimit}</h2>
    <h2>75% Utilized</h2>
</div>

<div className="transaction-history">
    <h1>Transaction History</h1>
    <table >
        <thead>
            <th>Date</th>
            <th>Type  </th>
            <th>Amount</th>
            <th>Status</th>
            <th>Reference</th>
        </thead>
        <tbody>
            {transactionshistory.map((txn, index) => (
             <tr>
                <td>{txn.date}</td>
                <td>{txn.type}</td>
                <td>{txn.amount}</td>
                <td>{txn.status}</td>
                <td>{txn.Reference}</td>
             </tr>
            ))}
        </tbody>
    </table>
</div>

<div className="account-action">
    <button  style={{width:'250px',marginLeft:'100px'}}>Generate Statment </button>
    <button  style={{width:'250px',marginLeft:'100px'}}>Close Statement</button>
    <button style={{width:'150px',marginLeft:'100px'}}>Block Account</button>
    <form action="" style={{marginTop:'20px',marginLeft:'22px'}}>
        <label htmlFor="" style={{border:'1px solid black',backgroundColor:'black',color:'white',padding:'5px',marginLeft:'50px'}}>Adjust Limits</label>
        <input  type="range" name="" id="" min='0' max='1000' step="10" />
    </form>
</div> 
<div className="account-det" id='account-det1'>
<h1>Account Details </h1>
<h1>Account Number:{dataVal.accountNumber}</h1>
<h1>Account Type:{dataVal.accountType}</h1>
{/* <h1>Account Id:{dataVal.accountId}</h1> */}
<h1>Currency :{dataVal.currency}</h1>
<h1>Status:{dataVal.status}</h1>
<h1>Available Balance:{dataVal.availableBalance}</h1>
<h1>Pending Balance:{dataVal.pendingBalance}</h1>
<h1>Credit Limit:{dataVal.creditLimit}</h1>
</div>
   </>
  )
}

export default Account
