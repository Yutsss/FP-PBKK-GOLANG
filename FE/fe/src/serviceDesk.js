// ServiceDesk.js
import React, { useState } from 'react';

function ServiceDesk() {
    const [ticketSubject, setTicketSubject] = useState('');
    const [email, setEmail] = useState('');
    const [priority, setPriority] = useState('Low');
    const [ticketType, setTicketType] = useState('Keluhan');
    const [destination, setDestination] = useState('');
    const [masalah, setMasalah] = useState('');
    const [ticketBody, setTicketBody] = useState('');
    const [telpArea, setTelpArea] = useState('');
    const [loginEmail, setLoginEmail] = useState('');
    const [loginPassword, setLoginPassword] = useState('');


    const handleTicketSubmit = (e) => {
        e.preventDefault();
        console.log('Ticket submitted:', {
            ticketSubject,
            email,
            priority,
            ticketType,
            destination,
            ticketBody,
        });
    };


    const handleLoginSubmit = (e) => {
        e.preventDefault();
        console.log('Login submitted', { loginEmail, loginPassword });

    }

    return (
        <div>
          <header style={{ backgroundColor: '#002d72', color: 'white', padding: '10px' }}>
          </header>
    
          <div style={{ padding: "20px", maxWidth: "800px", margin: "0 auto" }}>
    
            {/* <div style={{ backgroundColor: '#ffe6e6', padding: '10px', marginBottom: '20px', border: '1px solid #ffcccc', borderRadius: '5px' }}>
              <span style={{ color: 'red', marginRight: '5px' }}>⚠</span> Pengumuman Penutupan Jam Layanan
              <button style={{ float: "right" }} type="button" className="close" aria-label="Close"> <span aria-hidden="true">×</span> </button>
    
            </div> */}
    
    
            <div style={{ display: 'flex', justifyContent: "space-between" }}>
              <div style={{ width: '65%' }}>
    
                <h2>Need Support? Create a Ticket.</h2>
                <form onSubmit={handleTicketSubmit}>
                  <div>
                    <label htmlFor="ticketSubject">Ticket Subject:</label>
                    <input
                      type="text"
                      id="ticketSubject"
                      value={ticketSubject}
                      onChange={(e) => setTicketSubject(e.target.value)}
                      required
                    />
                  </div>
                  <div>
                    <label htmlFor="email">Your Email:</label>
                    <input
                      type="email"
                      id="email"
                      value={email}
                      onChange={(e) => setEmail(e.target.value)}
                      required
                    />
                  </div>
    
                  <div>
                    <label htmlFor="priority">Priority:</label>
                    <select id="priority" value={priority} onChange={(e) => setPriority(e.target.value)}>
                      <option value="Low">Low</option>
                      <option value="Medium">Medium</option>
                      <option value="High">High</option>
                    </select>
                  </div>
                  <div> <label>Ticket type:</label> <br />
                    <label>
                      <input type="radio" 
                      name="ticketType" 
                      value="Keluhan" 
                      checked={ticketType === 'Keluhan'} 
                      onChange={(e) => setTicketType(e.target.value)} /> 
                      Keluhan
                    </label>
    
                    <label>
                      <input
                        type="radio"
                        name="ticketType"
                        value="Permintaan/Pertanyaan"
                        checked={ticketType === 'Permintaan/Pertanyaan'}
                        onChange={(e) => setTicketType(e.target.value)}
                      /> Permintaan/Pertanyaan
                    </label>
                  </div>
    
                  <div>
                    <label htmlFor="destination">Unit of Destination:</label>
                    <select id="destination" value={destination} onChange={(e) => setDestination(e.target.value)}>
                      <option value="">Select...</option>
                      {/* Add options for units of destination */}
                      <option value="IT">IT</option>
                      <option value="HR">HR</option>
    
                    </select>
                  </div>

                  <div>
                    <label htmlFor="masalah">Tipe Masalah:</label>
                    <select id="masalah" value={masalah} onChange={(e) => setMasalah(e.target.value)}>
                      <option value="">Select...</option>
                      {/* Add options for units of destination */}
                      <option value="software">Software</option>
                      <option value="hardware">Hardware</option>
                      <option value="unkown">Tidak diketahui</option>
    
                    </select>
                  </div>
    
    
    
                  <div>
                    <label htmlFor="ticketBody">Ticket Body:</label>
                    <textarea id="ticketBody" value={ticketBody} onChange={(e) => setTicketBody(e.target.value)} />
                  </div>

                  <div>
                    <label htmlFor="notelp">No Telp:</label>
                    <textarea id="telparea" value={telpArea} onChange={(e) => setTelpArea(e.target.value)}></textarea>
                  </div>
    
    
    
                  <button type="submit">Submit Ticket</button>
                </form>
    
              </div>
    
              <div style={{ width: '30%' }}>
                <h2>Member Login</h2>
                <form onSubmit={handleLoginSubmit}>
    
                  <div>
                    <label htmlFor="loginEmail">Email or Username:</label>
                    <input type="text" id="loginEmail" value={loginEmail} onChange={(e) => setLoginEmail(e.target.value)} required />
                  </div>
                  <div>
                    <label htmlFor="loginPassword"> Password:</label>
                    <input type="password" id="loginPassword" value={loginPassword} onChange={(e) => setLoginPassword(e.target.value)} required />
                  </div>
                  <button type="submit">Login</button>
                </form>
                <a href="/forgotten-password">Forgotten Password?</a>
    
              </div>
            </div>
          </div>
    
        </div>
      );
    }
    
    export default ServiceDesk;