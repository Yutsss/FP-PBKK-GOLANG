import React, { useState } from "react";
import "./Login.css";

export const Login = () => {
  // Define state for email, password, login status, and user role
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loggedIn, setLoggedIn] = useState(false);
  const [role, setRole] = useState('guest'); // Default role

  // Handle login logic
  const handleLogin = async () => {
    try {
      const response = await fetch('https://api.beteam1genics.my.id/api/user/login', {
        method: 'POST',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }), // Mengirimkan email dan password
      });
  
      // Jika response tidak OK, lemparkan error
      if (!response.ok) {
        throw new Error('Invalid login credentials');
      }
  
      // Ambil data dari response (misalnya role dan token)
      const data = await response.json();
  
      // Jika login sukses, set user status
      setLoggedIn(true);
      setRole(data.role); // Misalnya, role datang dari response
  
      // Menyimpan token atau data lainnya ke localStorage/sessionStorage jika diperlukan
      localStorage.setItem('authToken', data.token); // Contoh token untuk autentikasi
  
    } catch (error) {
      // Tampilkan error jika ada
      alert(error.message);
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