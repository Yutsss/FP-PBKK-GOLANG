import React from 'react';
import RegistrationForm from './Regist';
import ServiceDesk from './serviceDesk';
import {BrowserRouter as Router, Routes, Route, Link} from 'react-router-dom'


function App() {
  return (
    <Router> {/*added a router to use the Link component*/}
      <div>
        <header style={{ backgroundColor: '#002d72', color: 'white', padding: '10px' }}>
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <div style={{ display: 'flex', alignItems: 'center' }}> <img src="logo192.png" alt="Logo" style={{ height: '50px' }} /> <h1>ServiceDesk</h1></div>
            <nav>
              <Link to="/">Home</Link> {/* Use Link component */}
              <Link to="/check-tickets">Check Tickets</Link> {/* Use Link component */}
              <Link to="/knowledge-base">Knowledge Base</Link> {/* Use Link component */}
              <Link to="/how-to-create-ticket">How to Create Ticket</Link> {/* Use Link component */}
              <Link to="/register" style={{ backgroundColor: 'red', color: 'white', padding: '5px 10px', borderRadius: '5px' }}>Register an Account</Link> {/* Use Link component */}
            </nav>
          </div>
        </header>


          <Routes> {/* Added routes component*/}
            <Route path="/" element={<ServiceDesk />} />
            <Route path="/register" element={<RegistrationForm />} />
          </Routes>

      </div>
    </Router>
  );
}


export default App;