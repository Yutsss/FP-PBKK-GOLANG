// ServiceDesk.js
import React, { useState } from "react";
import styles from "./serviceDesk.css";

function ServiceDesk() {
  const [ticketSubject, setTicketSubject] = useState("");
  const [ticketType, setTicketType] = useState("Keluhan");
  const [masalah, setMasalah] = useState("");
  const [ticketBody, setTicketBody] = useState("");

  const handleTicketSubmit = (e) => {
    e.preventDefault();
    console.log("Ticket submitted:", {
      ticketSubject,
      ticketType,
      ticketBody,
    });
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
                className={styles.value}
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
                className={styles.value}
              >
                <option value="">Select...</option>
                <option value="software">Software</option>
                <option value="hardware">Hardware</option>
                <option value="electrical">Electrical</option>
                <option value="network">Network</option>
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
                className={styles.value}
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