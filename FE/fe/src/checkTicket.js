import React, { useState, useEffect } from "react";
import "./checkTicket.css";

const TicketCard = ({ ticket, onClick }) => {
  return (
    <div className="ticket-card" onClick={() => onClick(ticket.id)}>
      <div className="ticket-info">
        <p><strong>Ticket Subject:</strong> {ticket.title}</p>
        <p><strong>Category:</strong> {ticket.category}</p>
        <p><strong>Description:</strong> {ticket.description}</p>
        <p><strong>Technician:</strong> {ticket.technician_id || "Not Assigned"}</p>
      </div>
      <div className="ticket-actions">
        <button
          className={`button ${ticket.status === "closed" ? "button-closed" : "button-open"}`}
        >
          {ticket.status === "closed" ? "Resolved" : "Unresolved"}
        </button>
        <button className="button contact-button">Contact</button>
      </div>
    </div>
  );
};

export const TechnicianTickets = () => {
  const [tickets, setTickets] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [selectedTicket, setSelectedTicket] = useState(null); // State untuk tiket yang dipilih
  const [modalLoading, setModalLoading] = useState(false); // State untuk loading saat modal dibuka
  const [modalError, setModalError] = useState(null); // State untuk error di modal

  useEffect(() => {
    const fetchTickets = async () => {
      setLoading(true);
      setError(null);

      try {
        const response = await fetch("https://api.beteam1genics.my.id/api/ticket/user", {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
        });

        if (!response.ok) {
          throw new Error(response.status === 401 ? "Unauthorized access" : "Failed to fetch tickets");
        }

        const data = await response.json();
        if (data.status && data.code === 200) {
          setTickets(data.data.tickets);
        } else {
          throw new Error(data.message || "Unexpected error");
        }
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchTickets();
  }, []);

  const handleCardClick = async (ticketId) => {
    setModalLoading(true);
    setModalError(null);

    try {
      const response = await fetch(`https://api.beteam1genics.my.id/api/ticket/${ticketId}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
      });

      if (!response.ok) {
        throw new Error("Failed to fetch ticket details");
      }

      const data = await response.json();
      if (data.status && data.code === 200) {
        setSelectedTicket(data.data.ticket);
      } else {
        throw new Error(data.message || "Unexpected error");
      }
    } catch (err) {
      setModalError(err.message);
    } finally {
      setModalLoading(false);
    }
  };

  const closeModal = () => {
    setSelectedTicket(null);
  };

  if (loading) return <div className="loading-spinner">Loading...</div>;
  if (error) return <div className="error-message">Error: {error}</div>;

  return (
    <div className="technician-tickets">
      {tickets.length === 0 ? (
        <p>No tickets available</p>
      ) : (
        tickets.map((ticket) => (
          <TicketCard key={ticket.id} ticket={ticket} onClick={handleCardClick} />
        ))
      )}

      {/* Modal */}
      {selectedTicket && (
        <div className="modal">
          <div className="modal-content">
            {modalLoading ? (
              <div className="loading-spinner">Loading ticket details...</div>
            ) : modalError ? (
              <div className="error-message">Error: {modalError}</div>
            ) : (
              <div>
                <h3>Ticket Details</h3>
                <p><strong>ID:</strong> {selectedTicket.id}</p>
                <p><strong>Title:</strong> {selectedTicket.title}</p>
                <p><strong>Description:</strong> {selectedTicket.description}</p>
                <p><strong>Category:</strong> {selectedTicket.category}</p>
                <p><strong>Status:</strong> {selectedTicket.status}</p>
                <p><strong>Technician:</strong> {selectedTicket.technician_id || "Not Assigned"}</p>
                <h4>Log:</h4>
                <p><strong>Activity:</strong> {selectedTicket.log.activity}</p>
              </div>
            )}
            <button className="close-button" onClick={closeModal}>Close</button>
          </div>
        </div>
      )}
    </div>
  );
};

export default TechnicianTickets;
