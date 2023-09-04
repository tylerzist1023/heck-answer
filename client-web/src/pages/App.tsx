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

function App() {
    const [username, setUsername] = useState("");

    useEffect(() => {
        let text = "";
        let res = fetch("/api/session/", {
            method: "get"
        })
        .then((r) => r.text())
        .then((text) => {
            setUsername(text);
        });
    }, []);

    return (
        <div className="App">
            <h1>Hello{username == "" ? "" : ","} {username}</h1>

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
                  </ul>
                </nav>

                <Routes>
                  <Route path="/" element={<Home />} />
                  <Route path="/login" element={<Login />} />
                  <Route path="/register" element={<Register />} />
                </Routes>
              </div>
            </Router>
        </div>
    );
}

export default App;
