import { useState } from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Signin from './components/signin'
import Home from './components/home.jsx'
import Account from './components/Account.jsx'
import Transaction from './components/Transaction.jsx'
import Payment from './components/Payment.jsx'
import Card from './components/Card.jsx'
import './App.css'
import Signup from './components/signup.jsx'

function App() {
 

  return (
    <>
      <Router>
      <Routes>
        <Route path="/" element={<Signin />} />
        <Route path="/Signin" element={<Signin />} />
        <Route path="/Home" element={<Home />} />
        <Route path="/Account" element={<Account/>} />
        <Route path="/Transaction" element={<Transaction/>} />
        <Route path="/Payment" element={<Payment/>} />
         <Route path="/Card" element={<Card/>} />
         <Route path="/Signup" element={<Signup />} />
      </Routes>
    </Router>
    </>
  )
}

export default App
