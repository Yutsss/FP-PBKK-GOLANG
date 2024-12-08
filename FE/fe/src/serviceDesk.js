// ServiceDesk.js
import React, { useState } from "react";
import "./serviceDesk.css";

function ServiceDesk() {
  const [ticketSubject, setTicketSubject] = useState("");
  const [email, setEmail] = useState("");
  const [priority, setPriority] = useState("Low");
  const [ticketType, setTicketType] = useState("Keluhan");
  const [destination, setDestination] = useState("");
  const [masalah, setMasalah] = useState("");
  const [ticketBody, setTicketBody] = useState("");
  const [telpArea, setTelpArea] = useState("");
  const [loginEmail, setLoginEmail] = useState("");
  const [loginPassword, setLoginPassword] = useState("");

  const handleTicketSubmit = (e) => {
    e.preventDefault();
    console.log("Ticket submitted:", {
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
    console.log("Login submitted", { loginEmail, loginPassword });
  };

  return (
    <div className="form">
      <div className="overlap">
        <div className="form-contact">
          <div className="text-wrapper-2">Need support? Create a ticket!</div>
          <div className="ticket-body">
            <div className="text-wrapper">Ticket Subject</div>
            <div className="text-area">
              <input
                type="text"
                id="ticketSubject"
                value={ticketSubject}
                onChange={(e) => setTicketSubject(e.target.value)}
                required
              />
            </div>
          </div>

          <div className="email">
            <div className="text-wrapper">Email</div>
            <div className="text-area">
              <input
                type="email"
                id="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
              />
            </div>
          </div>
          <div className="files">
            <div className="overlap-group">
              <div className="text-wrapper">Kategori</div>
              <select
                id="masalah"
                value={masalah}
                onChange={(e) => setMasalah(e.target.value)}
              >
                <option value="">Select...</option>
                <option value="software">Software</option>
                <option value="hardware">Hardware</option>
                <option value="unkown">Tidak diketahui</option>
              </select>
            </div>
          </div>

          <div className="ticket-body">
            <div className="text-wrapper">Ticket Body</div>
            <div className="text-area">
              <textarea
                id="ticketBody"
                value={ticketBody}
                onChange={(e) => setTicketBody(e.target.value)}
              />
            </div>
          </div>

          <div className="ticket-body">
            <div className="text-wrapper">Phone Number</div>
            <div className="text-area">
              <textarea
                id="telparea"
                value={telpArea}
                onChange={(e) => setTelpArea(e.target.value)}
              />
            </div>
          </div>

          <button className="button-group" type="submit">
            <button className="button">Submit Ticket</button>
          </button>
        </div>
      </div>
    </div>
  );
}

export default ServiceDesk;
