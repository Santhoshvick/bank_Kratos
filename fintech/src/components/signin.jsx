import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom';
// import { useState } from 'react';
import img1 from '../assets/stitch.svg'
import axios from 'axios';
import Home from './home';
import signup from './signup';
import { data, Link } from 'react-router-dom';



const signin = () => {

  const [signIn,SetSign]=useState({
    firstName:'',
    lastName:'',
    dateOfBirth:'',
    nationality:'',
    email:'',
    phone:'',
    address1:'',
    address2:'',
    city:'',
    state:'',
    postalCode:'',
    country:'',
    userName:'',
    password:''
  })
  const navigate = useNavigate();
  const nav = async (e) => {
    console.log('Form data before submit:', signIn);
    e.preventDefault();
    try {
       
      const response = await axios.post('http://localhost:8008/v1/user/create', signIn);
      console.log('User saved:', response.data);
      alert('Form submitted successfully!');
      navigate('/signup');
      alert("Registeration is done successfully")
    } catch (error) {
      console.error('Submission error:', error);
      alert('Please Fill the details!');
    }
  };


  const handleChange = (e) => {
  const { name, value } = e.target;
  let alphabtes=value
  if(name==='firstName'||name==='lastName'||name==='nationality'||name==='city'||name==='state'||name==='country'){
  alphabtes= value.replace(/[^A-Za-z]/g, '');
  }
  SetSign((prevData) => ({
    ...prevData,
    [name]: alphabtes
  }));
};

    
  // const handleLogin = () => {
   
  // };
  return (
    <>
    <div className='form-details'>
        <img src={img1} alt="val" style={{width:'200px',display:'block',marginBottom:'50px'}} />
        <form method='post' onSubmit={nav}>
            <label>Firstname</label> <br /> 
            <input type='text' placeholder='firstname' name='firstName' pattern='[A-Za-z]+' title='Only letters are allowed' value={signIn.firstName} onChange={handleChange} maxLength={30} required/> <br />  <br />
            <label>Lastname</label> <br /> 
            <input type='text' placeholder='lastname' name='lastName' value={signIn.lastName} onChange={handleChange} maxLength={30} required /> <br />  <br />
            <label>DateofBirth</label> <br />
            <input type='date' placeholder='dateofbirth'name='dateOfBirth' value={signIn.dateOfBirth} onChange={handleChange} maxLength={10} required /> <br /> <br />
            <label>Nationality</label> <br />
            <input type='text' placeholder='nationality' name='nationality' value={signIn.nationality} onChange={handleChange} maxLength={20} required/> <br /> <br />
            <label>Email</label> <br />
            <input type='email' placeholder='email' name='email' value={signIn.email} onChange={handleChange} maxLength={30} required/> <br /> <br />
            <label>Phone</label> <br /> 
            <input type="number" placeholder='phone' name='phone'value={signIn.phone} onChange={handleChange}maxLength={30} required/> <br /> <br />
            <label>Address1</label> <br />
            <input type='text' placeholder='address1' name='address1' value={signIn.address1} onChange={handleChange} maxLength={50} required/> <br /> <br />
        </form>
        <form action="" >
            <label>Address2</label> <br /> 
            <input type='text' placeholder='address2' name='address2' value={signIn.address2} onChange={handleChange} maxLength={50} required/> <br />  <br />
            <label>City</label> <br /> 
            <input type='text' placeholder='City' name='city' value={signIn.city} onChange={handleChange}maxLength={20} required/> <br />  <br />
            <label>State</label> <br />
            <input type='text' placeholder='state' name='state' value={signIn.state} onChange={handleChange}maxLength={20} required/> <br /> <br />
            <label>Postal code</label> <br />
            <input type='number' placeholder='postal code' name='postalCode' value={signIn.postalCode} onChange={handleChange} maxLength={20} required/> <br /> <br />
            <label>Country</label> <br />
            <input type='text' placeholder='country' name='country' value={signIn.country} onChange={handleChange} maxLength={20} required/> <br /> <br />
            <label>Username</label> <br /> 
            <input type="text" placeholder='username' value={signIn.userName} name='userName'required onChange={handleChange} maxLength={20}/> <br /> <br />
            <label>Password</label> <br />
            <input type='password' placeholder='password' name='password' value={signIn.password} required onChange={handleChange} maxLength={20} /> <br /> <br />
            
        </form>
        
    </div>
    <button onClick={nav}>Register</button>
    <Link className='link' to="/signup">Signin</Link>
    </>
  )
}

export default signin
