import { useContext, useEffect, useState } from 'react'
import { UserContext } from './context/UserContext'
import Home from './pages/Home'
import {Routes, Route, useNavigate} from "react-router-dom"
import Navbar from "../src/components/Navbar"
import CardJackets from './components/CardJackets'
import CardPants from './components/CardPants'
import CardAll from './components/CardAll'
import Login from './components/Login'
import DetailCard from './pages/DetailCard'
import Footer from './pages/Footer'
import FavoriteButton from './pages/FavoriteButton'

import { API, setAuthToken } from './config/Api'
import Register from './components/Register'


function App() {
  let navigate = useNavigate();
  const [state, dispatch] = useContext(UserContext);
  const [isLoading, setIsLoading] = useState(true)

  useEffect (() => {
    if (!isLoading) {
      if (state.isLogin === false) {
        navigate("/")
      }
    }
  }, [isLoading]);

  useEffect(() => {
    if (localStorage.token) {
      setAuthToken(localStorage.token)
      CheckAuth()
    } else {
      setIsLoading(false)
    }
  }, []);

  const CheckAuth = async () => {
    try {
      const response = await API.get('/check-auth')
      console.log('check user success : ', response)
      // get user data
      let payload = response.data.data
      // get token from localstorage
      payload.token = localStorage.token
      // send data to useContext
      dispatch ({
        type: 'USER_SUCCESS',
        payload,
      });
      setIsLoading(false)
    } catch(error) {
      console.log('check user failed : ', error)
      dispatch({
        type: 'AUTH_ERROR',
      })
      setIsLoading(false)
    }
  }

  return (
    <>
    {isLoading ? null : (
    <>
    <Navbar/>
    <Routes>
      <Route path='/' element={<Home />}/>
      <Route path='/jacket' element={<CardJackets/>}/>
      <Route path='/pants' element={<CardPants/>}/>
      <Route path='/all-product' element={<CardAll/>}/>
      <Route path='detail-card' element={<DetailCard/>}/>
    </Routes>
    <Footer/>
    </>
    )}
    </>
  )
}

export default App



