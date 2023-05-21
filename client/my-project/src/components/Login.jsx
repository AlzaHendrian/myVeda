import React, { useContext, useState } from "react";
import { ModalContext } from "../context/ModalContext";
import Register from "./Register";
import { UserContext } from "../context/UserContext";
import { useMutation } from "react-query";
import { API, setAuthToken } from "../config/Api";
import {useNavigate} from "react-router-dom"

const Login = () => {
    let navigate = useNavigate()
    const [message, setMessage] = useState(null)
    const [modalState, modalDispatch] = useContext(ModalContext);
    const [_, dispatch] = useContext(UserContext)
    const [form, setForm] = useState({
        email: "",
        password: "",
    })

    const {email, password} = form

    const handleOnChange = (e) => {
        setForm ({
            ...form,
            [e.target.name]: e.target.value,
        })
    }
    
    const handleSubmit = useMutation(async (e) => {
        try {
            e.preventDefault();
            const response = await API.post("/login", form)
            console.log("login success :", response)
            dispatch({
                type: 'LOGIN_SUCCESS',
                payload: response.data.data
            })
            setAuthToken(localStorage.token)
            if (response.data.data.role === "admin") {
                navigate("/detail-card")
            }else {
                navigate("/")
            }

            alert("successfully login!")
        }catch (err) {
            alert("Login Failed!")
        }
    })

  return (
    <>
      <div className="absolute w-[100%] mt-20 flex justify-center rounded-md">
        <form onSubmit={(e) => handleSubmit.mutate(e)} className="flex flex-col w-80 p-6 bg-gray-100 rounded-lg shadow-md fixed top-36">
          <div className="flex justify-between items-center">
            <h2 className="text-2xl font-bold">Login</h2>
            <h2
              className="font-bold text-2xl cursor-pointer"
              onClick={() => modalDispatch({ type: "CLOSE_MODAL" })}
            >
              X
            </h2>
          </div>
          <div className="mb-4 mt-4">
            <label htmlFor="email" className="block text-gray-700 font-semibold mb-2">
              Email
            </label>
            <input
              value={email}
              name="email"
              onChange={handleOnChange}
              type="email"
              id="email"
              className="border-gray-300 focus:border-indigo-500 focus:ring focus:ring-indigo-200 rounded-md p-2 w-full"
              required
            />
          </div>
          <div className="mb-6">
            <label htmlFor="password" className="block text-gray-700 font-semibold mb-2">
              Password
            </label>
            <input
              value={password}
              name="password"
              onChange={handleOnChange}
              type="password"
              id="password"
              className="border-gray-300 focus:border-indigo-500 focus:ring focus:ring-indigo-200 rounded-md p-2 w-full"
              required
            />
          </div>
          <button
            className="bg-indigo-500 text-white rounded-md py-2 px-4 hover:bg-indigo-600 transition duration-200"
          >
            Login
          </button>
          <div className="flex justify-center mt-5">
            <h2>
              don't have an account ?{" "}
              <button onClick={() => modalDispatch({ type: 'REGISTER_MODAL' })}>Sign up</button>
            </h2>
          </div>
        </form>
      </div>
      {modalState.isModalRegist && (
      <Register />
      )}
    </>
  );
};

export default Login;