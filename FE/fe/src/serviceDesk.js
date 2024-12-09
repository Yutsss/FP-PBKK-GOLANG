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

  const handleTicketSubmit = async (e) => {
    e.preventDefault();

    // Membuat data ticket yang akan dikirim ke backend
    const ticketData = {
      Title: ticketSubject,          // Menggunakan nilai dari input Ticket Subject
      Description: ticketBody,       // Menggunakan nilai dari input Ticket Body
      Category: masalah,             // Menggunakan nilai dari select box kategori
    };

    try {
      // Mengirim request POST ke backend
      const response = await fetch('https://api.beteam1genics.my.id/api/ticket/create', {
        credentials: 'include',
        method: 'POST',  // Menggunakan method POST
        headers: {
          'Content-Type': 'application/json',  // Menyatakan format data yang dikirim adalah JSON
        },
        body: JSON.stringify(ticketData),  // Mengirimkan data dalam format JSON
      });

      // Memeriksa apakah permintaan berhasil
      if (!response.ok) {
        throw new Error('Failed to submit ticket');
      }

      const data = await response.json();  // Menunggu respons JSON dari server
      console.log('Ticket submitted successfully:', data);

      // Menampilkan informasi sukses atau error
      alert('Ticket has been submitted successfully!');
      
      // Reset form setelah submit (optional)
      setTicketSubject('');
      setTicketBody('');
      setMasalah('');
    } catch (error) {
      console.error('Error submitting ticket:', error);
      alert('Error submitting ticket. Please try again later.');
    }
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
                <option value="network">Network</option>
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

          <button className="button-group" type="submit" onClick={handleTicketSubmit}>
            <button className="button">Submit Ticket</button>
          </button>
        </div>
      </div>
    </div>
  );
}

export default ServiceDesk;
