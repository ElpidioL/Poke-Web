import React, { useRef, useState, useEffect } from "react";
import {Link, useNavigate} from 'react-router-dom'
import { LoginPerson } from "../Defaults/classes"
import { Connect, Send } from "../Api/socketConnection";
import { Verify } from "../scripts/passwordVerify";
import classes from "./LoginPage.module.css";
import { GetColour, GetEmail } from "../scripts/getCookies";

function LoginPage(){
  const navigate = useNavigate(); 
  const [error, setError] = useState("");
  let inputEl = new LoginPerson()
  inputEl.email  = useRef(null);
  inputEl.password = useRef(null);

  function Nav() {
    navigate("../", { replace: true });
  }
  
  useEffect(() => {
    if(GetColour() && GetEmail()){
      Nav()
     }
  })

  function send(){
    let login = new LoginPerson(inputEl.email.current.value,  inputEl.password.current.value, "login" );
    if(Verify(login.password, login.email)){
      Connect();
      Send(login);
      setTimeout(Nav, 1100); 
    }else{
      setError(<p className={classes.error}>Weak Password or Login, Password should be at least 8 characters long</p>)
    }
  }

    return (
      <div>
        <div className={classes.background}>
          <div className={classes.shape}></div>
          <div className={classes.shape}></div>
        </div>
        <form className={classes.loginForm}>
          <h3>Login Here</h3>
          <label htmlFor ="Email">Email</label>
          <input type="email" placeholder="Email" id="Email"  ref={inputEl.email}/>

          <label htmlFor ="password" >Password</label>
          <input type="password" placeholder="Password" id="password" ref={inputEl.password}/>

          <Link to="/Register" className={classes.link}>Register</Link>

          <button onClick={send} type="button" className={classes.buttonForm}> Log In</button>
          <div>{error}</div>
        </form> 
      </div>
    );

}

export default LoginPage;
