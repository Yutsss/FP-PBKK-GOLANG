import React from "react";
import "./tech-home.css";

export const techHome = () => {
    return (
        <div className="box">
            <div className="group">
                <div className="overlap">
                    <div className="text-wrapper">Hey Technician!</div>

                    <div className="div">Your current stats are</div>

                    <div className="assigned">
                        <div className="overlap-group">
                            <div className="overlap-group-wrapper">
                                <div className="div-wrapper">
                                    <div className="text-wrapper-2">Assigned Tickets</div>
                                </div>
                            </div>

                            <div className="text-wrapper-3">12</div>
                        </div>
                    </div>

                    <div className="unresolved">
                        <div className="overlap-group">
                            <div className="overlap-group-wrapper">
                                <div className="overlap-group-2">
                                    <div className="text-wrapper-4">Unresolved Tickets</div>
                                </div>
                            </div>

                            <div className="text-wrapper-3">2</div>
                        </div>
                    </div>

                    <div className="resolved">
                        <div className="overlap-2">
                            <div className="group-wrapper">
                                <div className="group-2">
                                    <div className="overlap-group-3">
                                        <div className="text-wrapper-4">Resolved Tickets</div>
                                    </div>
                                </div>
                            </div>

                            <div className="text-wrapper-5">10</div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default techHome;