import React from 'react';
import {useState, useEffect} from 'react';
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Link,
  useNavigate
} from "react-router-dom";
import logo from './logo.svg';
import './App.css';
import Home from './Home';
import Login from './Login';
import Register from './Register';
import Submit from './Submit';
import Post from './Post';
import User from './User';

function App() {
    const [user, setUser] = useState();

    const refetchUser = () => {
        let res = fetch("/api/user", {
            method: "get"
        })
        .then((r) => r.json())
        .then((json) => {
            setUser(json);
        });
    };

    useEffect(refetchUser, []);

    return (
        <div className="App">
            <h1>Hello{user == undefined ? "" : ", "+user["username"]}</h1>

            <Router>
              <div>
                <nav>
                  <ul>
                    <li>
                      <Link to="/">Home</Link>
                    </li>
                    <li>
                      <Link to="/login">Login</Link>
                    </li>
                    <li>
                      <Link to="/register">Register</Link>
                    </li>
                    <li>
                      <Link to="/submit">New</Link>
                    </li>
                  </ul>
                </nav>

                <Routes>
                  <Route path="/" element={<Home />} />
                  <Route path="/login" element={<Login refetchUser={refetchUser} />} />
                  <Route path="/register" element={<Register refetchUser={refetchUser} />} />
                  <Route path="/submit" element={<Submit />} />
                  <Route path="/post" element={<Post />} />
                  <Route path="/user" element={<User />} />
                </Routes>
              </div>
            </Router>
        </div>
    );
}

export default App;
