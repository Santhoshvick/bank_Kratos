import React from 'react'
import { data, Link } from 'react-router-dom';
import img1 from '../assets/stitch.svg'
import img2 from '../assets/img2.png'
import img3 from '../assets/img3.png'
import img4 from '../assets/img4.avif'
import img5 from '../assets/img5.png'
import v1 from '../assets/video1.mp4'
import '../styles/Home.css'
import { useState } from 'react';
import axios from 'axios';
import img6 from '../assets/img6.png'
import { jwtDecode } from 'jwt-decode'; // 


// you will need to install via 'npm install jsonwebtoken' or in your package.json

// const jwt = require("jsonwebtoken");

const METABASE_SITE_URL = "http://localhost:3000";
const METABASE_SECRET_KEY = "31d4adb3959de05f97f4b23a4ac391edca9494153adb90f5915719ec844cb677";

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...";


const iframeUrl = METABASE_SITE_URL + "/embed/dashboard/" + token +
  "#bordered=true&titled=true";


const transactions = [
  {
    time: '12:45:32',
    type: 'Card Payment',
    amount: -45.99,
    status: 'POSTED'
  },
  {
    time: '12:44:15',
    type: 'ATM Withdrawal',
    amount: -100.0,
    status: 'PENDING'
  },
  {
    time: '12:43:01',
    type: 'Online Transfer',
    amount: 500.0,
    status: 'POSTED'
  },
  {
    time: '12:42:33',
    type: 'Fee Charge',
    amount: -2.5,
    status: 'POSTED'
  }
];







const Home = () => {
    const [dataVal, setFormData1] = useState({
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


    const handleFind= async(e)=>{
        e.preventDefault();
        try{
            const respone= await axios.get('http://localhost:8002/v1/account/find/21',dataVal)
            console.log("Account Details:",respone.data)
            setFormData1(respone.data)
            alert('Account is created successfully')
           let acc=document.getElementById('account-det1')
           acc.style.display = "inline";

        }
        catch(error){
            console.log("There is an issue with submitting the form data")
            alert('Please verify the api once')
        }
    }
    const [formData, setFormData] = useState({
     customerNumber: '',
  firstName: '',
  lastName: '',
  email: '',
  phone: '',
  dateOfBirth: '',
  status: ''
  });

 const handleChange = (e) => {
  const { name, value } = e.target;
  setFormData((prevData) => ({
    ...prevData,
    [name]: value
  }));
};


const handleSubmit = async (e) => {
    console.log('Form data before submit:', formData);
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8000/v1/create', formData);
      console.log('User saved:', response.data);
      alert('Form submitted successfully!');
    } catch (error) {
      console.error('Submission error:', error);
      alert('Something went wrong while submitting!');
    }
  };

  

  return (
    <>
    <nav className="navbar">
  <img src={img1} alt="Stitch Logo" className="logo" />
  <div className="nav-links">
    <Link className='link' to="/Home" onClick={handleFind}>Home</Link>
    <Link className='link' to="/account">Account</Link>
    <Link className='link' to="/transaction">Transaction</Link>
    <Link className='link' to="/payment">Payment</Link>
    <Link className='link' to="/card">Card</Link>
    <Link className='link' to="/signin">Signin</Link>
  </div>
</nav>
<video src={v1}  autoPlay loop  style={{width:'100%'}}/>
<iframe src="http://localhost:3000/public/question/bb63c1b8-7c56-41f2-a38c-4870cf39b55b" width="100%" height="600px" frameborder="0"></iframe>
<img src={img2} style={{width:'400px',height:'400px',display:'inline-block',marginLeft:'20px'}} />
<img src={img4} style={{width:'500px',height:'400px',display:'inline-block',marginLeft:'20px'}} />
<img src={img5} style={{width:'500px',height:'400px',display:'inline-block',marginLeft:'10px'}} />



<div className='dashboard'>
    <h1>Total Balance</h1>
    <h2>${dataVal.availableBalance}</h2>
    <h2>+2.3% from last</h2>
</div>

<div className='dashboard'>
    <h1>Active Cards</h1>
    <h2>12 Cards</h2>
    <h2>2 blocked</h2>
</div>
<div className='dashboard1'>
    <h1>Pending Transaction</h1>
    <h2>5 Transactions</h2>
    <h2>${dataVal.pendingBalance}</h2>
</div>

<div className="real-transactions">
     <h1>Real-Time Transaction Stream</h1>
     <div className="live-transaction">

     </div>


</div>

<div className="transaction-container">
      <div className="transaction-header">
        ⚡ Live Transactions
        <span className="auto-refresh">[Auto-refresh]</span>
      </div>
      <ul className="transaction-list">
        {transactions.map((txn, index) => (
          <li key={index} className={`transaction-item ${txn.status.toLowerCase()}`}>
            <span className="txn-time">{txn.time}</span>
            <span className="txn-type">{txn.type}</span>
            <span className={`txn-amount ${txn.amount < 0 ? 'negative' : 'positive'}`}>{txn.amount < 0 ? `-$${Math.abs(txn.amount).toFixed(2)}` : `+$${txn.amount.toFixed(2)}`}</span>
            <span className="txn-status">[{txn.status}]</span>
          </li>
        ))}
      </ul>
    </div>

    <div className="payments">
       <Link
  to="/payment"
  style={{
    display: 'inline-block',
    width: '150px',
    color: 'white',
    marginLeft: '200px',
    textDecoration: 'none',
    padding: '6px 12px',   
    backgroundColor: 'black', 
    borderRadius: '4px',
    textAlign: 'center',
    fontSize: '16px',
    fontWeight: 'bold',
    border: 'none',
    cursor: 'pointer',
  }}
>
  Payment
</Link>
<Link
  to="/card"
  style={{
    display: 'inline-block',
    width: '150px',
    color: 'white',
    marginLeft: '200px',
    textDecoration: 'none',
    padding: '6px 12px',   
    backgroundColor: 'black', 
    borderRadius: '4px',
    textAlign: 'center',
    fontSize: '16px',
    fontWeight: 'bold',
    border: 'none',
    cursor: 'pointer',
  }}
>
  Add Cards
</Link>
        <button className="btn" style={{width:'150px',color:'darkWhite',marginLeft:'200px'}}>Transfer</button>
        <button className="btn" style={{width:'150px',color:'darkWhite',marginLeft:'200px'}}>View Reports</button>
    </div>


    <div style={{ display: 'flex', alignItems: 'flex-start', marginTop: '20px', marginLeft: '5%' }}>

  <div className="create-customer" style={{ marginRight: '50px' }}>
    <h1>Customer Info</h1>
    <form method="post" onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', gap: '10px' }} >
      {/* <label>Customer ID</label>
      <input type="text" name="customerId"  onChange={handleChange}   required /> */}

      <label>Customer Number</label>
      <input type="text" name="customerNumber" value={formData.customerNumber} onChange={handleChange}  />

      <label>First Name</label>
      <input type="text" name="firstName" value={formData.firstName}  onChange={handleChange} />

      <label>Last Name</label>
      <input type="text" name="lastName" value={formData.lastName} onChange={handleChange}  />

      <label>Email</label>
      <input type="email" name="email" value={formData.email} onChange={handleChange}  />

      <label>Phone</label>
      <input type="number" name="phone" value={formData.phone} onChange={handleChange} />

      <label>Date Of Birth</label>
      <input type="date" name="dateOfBirth" value={formData.dateOfBirth} onChange={handleChange} />

      <label> Status</label>
      <select name="" id=""  value={formData.status} onChange={handleChange}>
        <option value="">Active</option>
        <option value="">Suspended</option>
        <option>Closed</option>
      </select>
      <button type="submit" className='customer'>Submit</button>
    </form>
  </div>


  <img src={img6} alt="Customer Visual" style={{ width: '900px', height: '800px', objectFit: 'cover', borderRadius: '8px',marginLeft:'100px',  marginTop:'30px'}}/>
</div>
   
   <footer className="footer">
      <div className="footer-container">


        <div className="footer-section">
          <h3>FinTrust Bank</h3>
          <p>Secure. Smart. Seamless Banking for the Digital Era.</p>
        </div>

     
        <div className="footer-section">
          <h4>Quick Links</h4>
          <ul>
            <li><a href="/about">About Us</a></li>
            <li><a href="/services">Services</a></li>
            <li><a href="/support">Customer Support</a></li>
            <li><a href="/privacy-policy">Privacy Policy</a></li>
            <li><a href="/terms">Terms of Use</a></li>
          </ul>
        </div>

       
        <div className="footer-section">
          <h4>Contact Us</h4>
          <p>📞 +1 (800) 555-0199</p>
          <p>✉️ support@fintrustbank.com</p>
          <p>🏢 123 Fintech Avenue, San Francisco, CA</p>
        </div>

        <div className="footer-section">
          <h4>Follow Us</h4>
          <div className="social-links">
            <a href="#">LinkedIn</a>
            <a href="#">Twitter</a>
            <a href="#">Facebook</a>
          </div>
        </div>

      </div>

      <div className="footer-bottom">
        &copy; 2025 FinTrust Bank. All rights reserved. FinTrust is a registered trademark. Member FDIC.
      </div>
    </footer>

    </>
  )
}

export default Home
