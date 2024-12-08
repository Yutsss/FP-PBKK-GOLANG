import React from "react";
import "./tech-tickets.css";

export const TechnicianTickets = () => {
    return (
        <div className="technician-tickets">
            <div className="div">
                <div className="overlap">
                    <div className="ticket-subject">
                        Ticket Subject:
                        <br />
                        Category:
                        <br />
                        Body:
                        <br />
                        Technician: A
                    </div>

                    <div className="button-group">
                        <button className="button">
                            <button className="text-wrapper">Unresolved</button>
                        </button>
                    </div>

                    <button className="button-wrapper">
                        <button className="button-2">Contact</button>
                    </button>
                </div>

                <div className="overlap-group">
                    <div className="ticket-subject-2">
                        Ticket Subject:
                        <br />
                        Category:
                        <br />
                        Body:
                        <br />
                        Technician:
                    </div>

                    <button className="div-wrapper">
                        <button className="button-3">Resolved</button>
                    </button>
                </div>

                <div className="overlap-2">
                    <div className="ticket-subject-2">
                        Ticket Subject:
                        <br />
                        Category:
                        <br />
                        Body:
                        <br />
                        Technician: A
                    </div>

                    <div className="button-group">
                        <button className="button-4">
                            <button className="button-3">Resolved</button>
                        </button>
                    </div>

                    <button className="button-5">
                        <button className="button-2">Contact</button>
                    </button>
                </div>

                <div className="overlap-3">
                    <div className="ticket-subject-2">
                        Ticket Subject:
                        <br />
                        Category:
                        <br />
                        Body:
                        <br />
                        Technician: A
                    </div>

                    <div className="button-group-2">
                        <button className="button-4">
                            <button className="text-wrapper">Unresolved</button>
                        </button>
                    </div>

                    <button className="button-6">
                        <button className="button-2">Contact</button>
                    </button>
                </div>
            </div>
        </div>
    );
};

export default TechnicianTickets;