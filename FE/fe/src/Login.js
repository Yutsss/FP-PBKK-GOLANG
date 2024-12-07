import React, { useState } from "react";
import "./Login.css";

export const Login = () => {
  // Define state for email, password, login status, and user role
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loggedIn, setLoggedIn] = useState(false);
  const [role, setRole] = useState('guest'); // Default role

  // Handle login logic
  const handleLogin = () => {
    // Simulating login (replace with real logic)
    if (email === 'admin@example.com' && password === 'admin') {
      // Set user as logged in and assign admin role
      setLoggedIn(true);
      setRole('admin');
    } else if (email === 'user@example.com' && password === 'user') {
      setLoggedIn(true);
      setRole('user');
    } else {
      alert('Invalid login credentials');
    }
  };

  return (
    <div className="login">
      <div className="form-log-in">
        {/* Email input */}
        <div className="input-field">
          <div className="label">Email</div>
          <input
            type="email"
            name="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            className="value"
          />
        </div>

        {/* Password input */}
        <div className="input-field">
          <div className="label">Password</div>
          <input
            type="password"
            name="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            className="value"
          />
        </div>

        {/* Sign In button */}
        <div className="button-group">
          <button className="button" onClick={handleLogin}>
            Sign In
          </button>
        </div>
      </div>

      {/* Login header text */}
      <div className="text-wrapper-2">Login</div>
    </div>
  );
};
