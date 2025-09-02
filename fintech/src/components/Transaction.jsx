import React from 'react'
import img1 from '../assets/stitch.svg'
import '../styles/Account.css'
import '../styles/Home.css'
import '../styles/Transaction.css'
import { Link } from 'react-router-dom';
import Account from './Account';
import img7 from '../assets/img7.jpg'

const tablevalue=[
    {   id:1,
        Status:'🟡',
        Time:'12:25:45',
        Account:'CHK-001234',
        Type:'Card Pay',
        Amount:'-$67.89 '
    },
    {
        id:2,
        Status:'🟡',
        Time:'11:12:45',
        Account:'CHK-903241',
        Type:'UPI Pay',
        Amount:'+$97.89 '
    },
    {
        id:3,
        Status:'🔴',
        Time:'09:45:45',
        Account:'CHK-002421',
        Type:'Transfer In',
        Amount:'+$20.89 '
    },
    {   
        id:4,
        Status:'🟢 ',
        Time:'08:25:15',
        Account:'CHK-992123',
        Type:'Online Pay',
        Amount:'-$7.89 '
    },
    {
        id:5,
        Status:'🔴 ',
        Time:'12:44:58',
        Account:'SAV-778899 ',
        Type:'Fee Charge ',
        Amount:'-$5.00  '
    }
]



const Transaction = () => {
    const close=async(e)=>{
        let val=document.getElementById('payment');
        val.style.display='none'
    }
    const action=async(e)=>{
        let val=document.getElementById('payment');
        val.style.display='inline-block'
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
<img src={img7} alt="" srcset="" style={{height:'300px',width:'70%',marginLeft:'15%',marginTop:'20px'}}/>
<iframe src="http://localhost:3000/public/question/d4fabb90-60f9-4ab9-816d-99d447f299a9" width="100%" height="600px"></iframe>
<h1>Transactions</h1>
<iframe src="http://localhost:3000/public/question/14c3fc81-9d6c-4c6b-9859-231d0c93f488" style={{width:'100%',height:'1000px',border:'2px solid red'}} frameborder="0"></iframe>
<div className="transaction-details">
 <table>
    <thead>
      <tr>
        <th>Status</th>
        <th>Date</th>
        <th>Amount</th>
        <th>Account</th>
        <th>Type</th>
      </tr>
    </thead>
    <tbody>
      {tablevalue.map((val, index) => (
        <tr key={index}>
          <td>{val.Status}</td>
          <td>{val.Time}</td>
          <td>{val.Amount}</td> 
          <td>{val.Account}</td>
          <td>{val.Type}</td>
        </tr>
))} </tbody></table>
</div>
<h2 style={{display:'block'}}> Legend: 🟢 Posted  🟡 Pending  🔴 Failed </h2>
<h2 style={{marginTop:'0px',fontSize:'30px'}}>Transaction Details</h2>
<button onClick={close} className='close-btn'>Close</button>
<button onClick={action} className='action-btn'>Actions</button>
<div className="payment" id='payment'>
  {tablevalue.map((val, index) => {
    if (val.id === 1 || val.id==2) {
      return (
        <div key={index}>
          <h3>Status:{val.Status}</h3>
          <h3>Created Time:{val.Time}</h3>
          <h3>Amount{val.Amount}</h3> 
          <h3>Account:{val.Account}</h3>
          <h3>Type:{val.Type}</h3>
        </div>
      );
    } else {
      return null;
    }
  })}
</div>

    </>
  )
}

export default Transaction
