import React, { useState } from 'react'

import axios from 'axios';
import img1 from '../assets/stitch.svg'
import { Link, redirect } from 'react-router-dom';
import img10 from '../assets/img10.webp'
import img14 from '../assets/img14.jpg'
import img15 from '../assets/img15.png'
import img16 from '../assets/img16.webp'
import img17 from '../assets/img17.jpg'
import img19 from '../assets/img19.webp'
import img20 from '../assets/img20.jpeg'
import '../styles/Card.css'

const Card = () => {

  const [cardData,setCard]=useState({
    cardId:'',
    accountNumber:'',
    cardNumber:'',
    cardType:'',
    expiryDate:'',
    dailyLimit:'',
    monthlyLimit:'',
    pinAttempt:'',
  })
  console.log(cardData,"setset")

  const card1=async(e)=>{
     try{const response= await axios.get('http://localhost:8009/v1/find/4',cardData)
            setCard(response.data)
            console.log(cardData)
            alert('Card is applied Successfully we will verify your details soon and Will get back to you')
        }
        catch(error){
            console.log(error)
            alert("Please Verify the port")
        }
  }

  const debit=async(e)=>{
    let val1=document.getElementById('debitcard')
    val1.style.display="inline-block"
    let val2=document.getElementById('creditcard')
    val2.style.display="none"
    let val3=document.getElementById('forexcard')
    val3.style.display="none"
    console.log("Hello World");
  }
  const credit=async(e)=>{
    let val1=document.getElementById('debitcard')
    val1.style.display="none"
    let val2=document.getElementById('creditcard')
    val2.style.display="inline-block"
    let val3=document.getElementById('forexcard')
    val3.style.display="none"
    console.log("Hello World");
  }
  const forex=async(e)=>{
    let val1=document.getElementById('debitcard')
    val1.style.display="none"
    let val2=document.getElementById('creditcard')
    val2.style.display="none"
    let val3=document.getElementById('forexcard')
    val3.style.display="inline-block"
    console.log("Hello World");
  }
  return (
    <>
     <nav className="navbar">
  <img src={img1} alt="Stitch Logo" className="logo" />
  <div className="nav-links">
    <Link className='link' to="/Home">Home</Link>
    <Link className='link' to="/account">Account</Link>
    <Link className='link' to="/transaction">Transaction</Link>
    <Link className='link' to="/payment">Payment</Link>
    <Link className='link' to="/card">Card</Link>
    <Link className='link' to="/signin">Signin</Link>
  </div>
</nav>
<img src={img10} style={{width:'100%'}}  alt="" srcset="" />

<div className="card-detaits" >
    <p  className='card-explore' > 💳 Explore Our Cards</p>
    <p  className='card-explore' style={{display:'inline-block'}}>Choose the right card that fits your lifestyle </p>
    <button className='cards-type' onClick={debit}>Debit Card</button>
    <button className='cards-type' onClick={credit}>Credit Card</button>
    <button className='cards-type' onClick={forex}>Forex Card</button>
</div>

<div className="debitcard" id='debitcard' style={{display:'none'}}>
    <form action="" className="debit-card" style={{width:'400px',backgroundColor:'lightcyan'}}>
      <h2 style={{display:'inline-block',color:'royalblue'}}>Debit Card</h2>
        <label htmlFor="" style={{display:'block'}}>Full Name</label> <br />
         <input type="text" name="fullname" id="" style={{marginTop:'-30px'}} /> <br />
        <label htmlFor="">Date of Birth</label> <br />
        <input type="datetime-local" name="" id="" required/> <br />
        <label>Email</label> <br />
        <input type="email" name="" id="" required/> <br />
        <label>Mobile No</label><br />
        <input type="number" name="" id="" required/> <br />
        <label>Resential Address</label> <br />
        <input type="text" name="" id="" required/> <br />
        <label>PAN Number</label> <br />
        <input type="text" /> <br />
        <label htmlFor="">ID Proof</label> <br />
        <input type="file" name="" id="" required/> <br />
        <button style={{display:'inline-block',marginLeft:'90px',width:'200px'}} onClick={card1}>Apply Card</button>
    </form>
    <img src={img14} alt="" srcset=""  style={{display:'inline-block',width:'800px',marginLeft:'150px'}}/>
</div>

<div className="forexcard" id='creditcard' style={{display:'none'}} >
    <form action="" className="debit-card" style={{width:'400px',backgroundColor:'lightsteelblue'}}>
      <h2 style={{display:'inline-block',color:'royalblue'}}>Credit Card</h2>
        <label htmlFor="" style={{display:'block'}}>Full Name</label> <br />
         <input type="text" name="fullname" id="" style={{marginTop:'-30px'}} required/> <br />
        <label htmlFor="">Date of Birth</label> <br />
        <input type="datetime-local" name="" id="" required/> <br />
        <label>Email</label> <br />
        <input type="email" name="" id="" required/> <br />
        <label>Mobile No</label><br />
        <input type="number" name="" id="" required/> <br />
        <label>Resential Address</label> <br />
        <input type="text" name="" id="" required/> <br />
        <label>PAN Number</label> <br />
        <input type="text" required/> <br />
        <label htmlFor="">ID Proof</label> <br />
        <input type="file" name="" id="" required/> <br />
        <button style={{display:'inline-block',marginLeft:'90px',width:'200px'}} onClick={card1}>Apply Card</button>
    </form>
    <img src={img15} alt="" srcset=""  style={{display:'inline-block',width:'800px',marginLeft:'150px'}}/>
</div>
<div className="forexcard" id='forexcard'>
    <form action="" className="debit-card" style={{width:'400px',backgroundColor:'lightseagreen'}}>
      <h2 style={{display:'inline-block',color:'royalblue'}}>Forex Card</h2>
        <label htmlFor="" style={{display:'block'}}>Full Name</label> <br />
        <input type="text" name="fullname" id="" style={{marginTop:'-30px'}} required/> <br />
        <label htmlFor="">Date of Birth</label> <br />
        <input type="datetime-local" name="" id="" required/> <br />
        <label>Email</label> <br />
        <input type="email" name="" id="" required/> <br />
        <label>Mobile No</label><br />
        <input type="number" name="" id="" required/> <br />
        <label>Resential Address</label> <br />
        <input type="text" name="" id="" required/> <br />
        <label>PAN Number</label> <br />
        <input type="text" /> <br />
        <label htmlFor="">ID Proof</label> <br />
        <input type="file" name="" id="" required/> <br />
        <button style={{display:'inline-block',marginLeft:'90px',width:'200px'}} onClick={card1}>Apply Card</button>
    </form>
    <img src={img16} alt="" srcset=""  style={{display:'inline-block',width:'800px',marginLeft:'150px'}}/>
</div>

<div className="display-card" >
  <img src={img17} alt="" srcset="" style={{width:'50px',display:'block',marginLeft:'10px'}}/>
  <h5 style={{display:'inline-block',marginTop:'0px'}}>PREPAID TRAVEL CARD</h5>
  <img src={img20} alt="" srcset="" style={{width:'90px',display:'block'}}/>
  <h5 style={{marginLeft:'40px',fontSize:'20px'}}>{cardData.cardNumber}</h5>
  <h5 style={{marginLeft:'40px',display:'inline-block',marginTop:'-30px'}}>{cardData.expiryDate}</h5>
  <img src={img19} alt="" srcset="" style={{width:'100px',marginLeft:'390px',marginTop:'-70px',display:'inline-block'}} />
</div>
    </>
  )
}

export default Card
