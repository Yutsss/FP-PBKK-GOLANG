import React, { useState, useEffect } from "react";
import "./admin-ticket.css";
import AdminTechnician from "./admin-technician";

export const AdminTicket = () => {
  const [tickets, setTickets] = useState([]); // State untuk tiket
  const [loading, setLoading] = useState(true); // State untuk loading
  const [error, setError] = useState(null); // State untuk error
  const [isPopupVisible, setIsPopupVisible] = useState(false); // State untuk popup tiket
  const [selectedTicket, setSelectedTicket] = useState(null); // State untuk tiket yang dipilih
  const [assignFormVisible, setAssignFormVisible] = useState(false); // State untuk form assign
  const [technicianID, setTechnicianID] = useState(""); // State untuk input technicianID

  // Fetch tiket saat komponen pertama kali dimuat
  useEffect(() => {
    const fetchTickets = async () => {
      setLoading(true);
      setError(null);

      try {
        const response = await fetch("https://api.beteam1genics.my.id/api/ticket/all", {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
        });

        if (!response.ok) {
          throw new Error("Failed to fetch tickets");
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

  // Fungsi untuk menampilkan popup dengan detail tiket
  const handleUnresolvedClick = (ticket) => {
    setSelectedTicket(ticket);
    setIsPopupVisible(true);
  };

  // Fungsi untuk menutup popup
  const handleClosePopup = () => {
    setIsPopupVisible(false);
    setSelectedTicket(null);
  };

  // Fungsi untuk membuka form assign
  const handleAssignClick = (ticket) => {
    setSelectedTicket(ticket);
    setAssignFormVisible(true);
  };

  // Fungsi untuk menutup form assign
  const handleCloseAssignForm = () => {
    setAssignFormVisible(false);
    setTechnicianID("");
  };

  // Fungsi untuk submit assign
  const handleAssignSubmit = async (e) => {
    e.preventDefault();

    if (!technicianID) {
      alert("Please enter a technician ID.");
      return;
    }

    try {
      const response = await fetch(
        `https://api.beteam1genics.my.id/api/ticket/${selectedTicket.id}/assign/${technicianID}`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
        }
      );

      if (!response.ok) {
        throw new Error("Failed to assign technician");
      }

      const data = await response.json();
      if (data.status && data.code === 200) {
        alert("Technician assigned successfully!");
        // Update state to reflect the change (refresh ticket list or update locally)
        setTickets((prevTickets) =>
          prevTickets.map((ticket) =>
            ticket.id === selectedTicket.id
              ? { ...ticket, technician_id: technicianID }
              : ticket
          )
        );
        setAssignFormVisible(false);
        setTechnicianID("");
      } else {
        throw new Error(data.message || "Unexpected error");
      }
    } catch (err) {
      alert(`Error: ${err.message}`);
    }
  };

  if (loading) return <div className="loading-spinner">Loading...</div>;
  if (error) return <div className="error-message">Error: {error}</div>;

  return (
    <div className="admin-ticket">
      <div className="ticket-list">
        {tickets.map((ticket) => (
          <div key={ticket.id} className="ticket-card">
            <div className="ticket-info">
              <p><strong>Ticket Subject:</strong> {ticket.title}</p>
              <p><strong>Category:</strong> {ticket.category}</p>
              <p><strong>Description:</strong> {ticket.description}</p>
              <p><strong>Technician:</strong> {ticket.technician_id || "Not Assigned"}</p>
            </div>
            <div className="ticket-actions">
              {ticket.status === "open" ? (
                <button
                  className="button unresolved-button"
                  onClick={() => handleUnresolvedClick(ticket)}
                >
                  Unresolved
                </button>
              ) : (
                <button className="button resolved-button">Resolved</button>
              )}
              <button
                className="button assign-button"
                onClick={() => handleAssignClick(ticket)}
              >
                Assign
              </button>
            </div>
          </div>
        ))}
      </div>

      {/* Popup untuk tiket */}
      {isPopupVisible && selectedTicket && (
        <div className="popup-overlay">
          <div className="popup-content">
            <button className="close-button" onClick={handleClosePopup}>
              X
            </button>
            <h3>Ticket Details</h3>
            <p><strong>ID:</strong> {selectedTicket.id}</p>
            <p><strong>Title:</strong> {selectedTicket.title}</p>
            <p><strong>Description:</strong> {selectedTicket.description}</p>
            <p><strong>Category:</strong> {selectedTicket.category}</p>
            <p><strong>Status:</strong> {selectedTicket.status}</p>
            <p><strong>Technician:</strong> {selectedTicket.technician_id || "Not Assigned"}</p>
            <AdminTechnician />
          </div>
        </div>
      )}

      {/* Form untuk assign technician */}
      {assignFormVisible && selectedTicket && (
        <div className="popup-overlay">
          <div className="popup-content">
            <button className="close-button" onClick={handleCloseAssignForm}>
              X
            </button>
            <h3>Assign Technician</h3>
            <form onSubmit={handleAssignSubmit}>
              <p><strong>Ticket:</strong> {selectedTicket.title}</p>
              <input
                type="text"
                placeholder="Enter Technician ID"
                value={technicianID}
                onChange={(e) => setTechnicianID(e.target.value)}
              />
              <button type="submit" className="button submit-button">Submit</button>
            </form>
          </div>
        </div>
      )}
    </div>
  );
};

export default AdminTicket;
