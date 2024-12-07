import React from "react";
import "./Header.css";

export const Header = () => {
    return (
        <div className="header">
            <div className="overlap-group">
                <div className="register"> Register</div>

                <div className="text-wrapper">Check tickets</div>

                <div className="div">Create a Ticket</div>

                <div className="text-wrapper-2">Tech Support</div>
            </div>
        </div>
    );
};
