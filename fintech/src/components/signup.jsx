import React, { useState } from 'react'
import img1 from '../assets/stitch.svg'
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

const signup = () => {

    const [val1,setVal]=useState({
        email:'',
        password:'',
    })

    const [val2,setVal2]=useState({
        email:'',
        password:'',
        otp:''
    })

    const [otp1,setOTP1]=useState({
        otp:''
    })
     const [otp2,setOTP2]=useState({
        otp:'',
        email:''
    })

    console.log("OTP2 Email",otp2.email)

    const navigate = useNavigate();

    const fetch1=async(e)=>{
        e.preventDefault();
    try{
        console.log(val1)
        console.log(val2)
       const response = await axios.get(`http://localhost:8008/v1/user/find/${val1.email}`,val2);
       const responseData=response.data
        setVal2(responseData)
        if(val1.email===val2.email&& otp1.otp==otp2.otp && otp2!=" "){
            navigate('/home');
            alert("Login ssuccessfully")
        }
        else{
            alert('Invalid Email name or password')
        }
    }
    catch(error){
        console.log(error);
        // alert('please verify the api once')
    }
    }

    console.log("OTP2 Email",otp2.email)
    console.log("Val1 Email",val1.email)
    console.log("Val2 Email",val2.email)

    

      const oneTimePassword=async(e)=>{
        e.preventDefault();
        try{
        const response = await axios.post(`http://localhost:8020/v1/alert/otp`, otp2 )
        const response1 = await axios.get(`http://localhost:8008/v1/user/find/${val1.email}`,val2);
        const responseData=response1.data
        otp2.email=responseData.email

        const responseOtp=response.data
        if(val1.email!="" && val1.password!="" && val1.password!=val2.password && val1.email!=val2.email)
        {
          setOTP1(responseOtp)
          alert("otp has been sent successfull")
        }
        else if(val1.email==""){
            alert("Please Enter the Email First to get the otp")
        }
        else if(val1.password==""){
            alert("Please Enter the correct Password")
        }
        }
        catch{
            alert("Invalid OTP Please enter a Correct OTP")
        }
    }
       console.log(otp2.email)
       console.log(otp1)
       console.log(otp2)
//   const handleLogin = () => {
//     navigate('/home');
//     alert("Login ssuccessfully")
//   };

// console.log(val1)
// console.log(val2)


  const handleChange = (e) => {
  const { name, value } = e.target;
  setVal((prevData) => ({
    ...prevData,
    [name]: value
  }));
};

const handleChange1 = (e) => {
  const { name, value } = e.target;
  setOTP2((prevData) => ({
    ...prevData,
    [name]: value
  }));
};

  return (
    <>
    <div className="sign-up" style={{width:'100px',marginLeft:'35%',marginTop:'10%'}}>
        <img src={img1} alt="val" style={{width:'200px',display:'block',marginBottom:'50px'}} />
        <form action="" style={{fontSize:'30px'}} onSubmit={fetch1}>
        <label htmlFor="">Email</label> <br />
        <input type='email' name='email' placeholder='Email Id' onChange={handleChange} value={val1.email} maxLength={40} required/> <br />
        <label>Password</label> <br />
        <input type='password' name='password' placeholder='password'value={val1.password}  onChange={handleChange} maxLength={20} required/> <br />
        <input type='number' name='otp' placeholder='OTP' style={{display:'inline-block',width:'110px'}} onChange={handleChange1}  value={otp2.otp} />
        <button style={{width:'150px',marginLeft:'40px',marginTop:'-800px',display:'inline-block'}} onClick={oneTimePassword}>Send OTP</button>
        <button style={{marginLeft:'120px'}}>Login</button>
    </form>
    </div>
    </>
  )
}


export default signup

