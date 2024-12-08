import React, { useState } from "react";
import "./admin-ticket.css";
import AdminTechnician from "./admin-technician"; // Change: Imported AdminTechnician component

export const AdminTicket = () => {
    const [isPopupVisible, setIsPopupVisible] = useState(false); // Change: Added state for popup visibility

    const handleUnresolvedClick = () => {
        setIsPopupVisible(true); // Change: Show popup
    };

    const handleClosePopup = () => {
        setIsPopupVisible(false); // Change: Hide popup
    };

    return (
        <div className="admin-ticket">
            <div className="div">
                <div className="group">
                    <div className="overlap-group">
                        <div className="ticket-subject">
                            Ticket Subject:
                            <br />
                            Category: B<br />
                            Body:
                            <br />
                            Technician:
                        </div>

                        <button className="button" onClick={handleUnresolvedClick}> {/* Change: Added onClick */}
                            Unresolved
                        </button>
                    </div>
                </div>

                {/* Other tickets */}
                <div className="overlap-wrapper">
                    <div className="overlap">
                        <div className="ticket-subject">
                            Ticket Subject:
                            <br />
                            Category: A<br />
                            Body:
                            <br />
                            Technician: B
                        </div>

                        <div className="button-group">
                            <button className="button-wrapper">
                                <button className="button-2">Resolved</button>
                            </button>
                        </div>
                    </div>
                </div>

                {/* Popup */}
                {isPopupVisible && ( // Change: Popup added here
                    <div className="popup-overlay">
                        <div className="popup-content">
                            <button className="close-button" onClick={handleClosePopup}> {/* Change: Close button */}
                                X
                            </button>
                            <AdminTechnician />
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
};

export default AdminTicket;
