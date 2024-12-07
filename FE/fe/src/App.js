import React, { useState } from 'react';
import UserRegister from './Regist';
import ServiceDesk from './serviceDesk';
import UserCheckTicketsLoggedIn from './checkTicket';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';

function App() {
  // Example state for user login status and role
  const [user, setUser] = useState({ loggedIn: false, role: 'guest' }); // 'guest' is the default

  return (
    <Router>
      {/* Added a router to use the Link component */}
      <div style={{ backgroundColor: '#0d0911' }}>
        <header style={{ backgroundColor: '#0d0911', color: 'white', padding: '10px', border: '1px solid', borderColor: '#b892d8', marginTop:'-1px', marginLeft:'-1px', marginRight:'-1px'}}>
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <div style={{ display: 'flex', alignItems: 'center' }}>
              <img src="logo192.png" alt="Logo" style={{ height: '50px' }} />
              <h1>Tech Support</h1>
            </div>
            <nav>
              <Link
                to="/"
                style={{
                  color: 'white',
                  padding: '5px 10px',
                  borderRadius: '5px',
                  textDecoration: 'none',
                  font: "Noto Sans KR-Regular",
                }}
              >
                Home
              </Link>
              <Link
                to="/check-tickets"
                style={{
                  color: 'white',
                  padding: '5px 10px',
                  borderRadius: '5px',
                  textDecoration: 'none',
                  font: "Noto Sans KR-Regular",
                }}
              >
                Check Tickets
              </Link>
              <Link
                to="/create-ticket"
                style={{
                  color: 'white',
                  padding: '5px 10px',
                  borderRadius: '5px',
                  textDecoration: 'none',
                }}
              >
                Create Ticket
              </Link>

              {/* Conditional rendering based on login state */}
              {!user.loggedIn ? (
                <Link
                  to="/register"
                  style={{
                    color: 'white',
                    padding: '5px 10px',
                    borderRadius: '5px',
                    textDecoration: 'none',
                  }}
                >
                  Register an Account
                </Link>
              ) : (
                <Link
                  to={`/${user.role}`}
                  style={{
                    color: 'white',
                    padding: '5px 10px',
                    borderRadius: '5px',
                    textDecoration: 'none',
                  }}
                >
                  {user.role === 'admin' ? 'Admin' : user.role === 'technician' ? 'Technician' : 'User'}
                </Link>
              )}
            </nav>
          </div>
        </header>

        {/* Added routes component */}
        <Routes>
          <Route path="/" element={<ServiceDesk />} />
          <Route path="/register" element={<UserRegister />} />
          {/* Dynamic routes for different user roles */}
          <Route path="/user" element={<div>User Dashboard</div>} />
          <Route path="/admin" element={<div>Admin Dashboard</div>} />
          <Route path="/technician" element={<div>Technician Dashboard</div>} />
          <Route path="/tickets" element={<div>UserCheckTicketsLoggedIn</div>} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;