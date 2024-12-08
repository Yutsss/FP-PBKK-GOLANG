import React from "react";
import { Header } from "./Header";
import { Login } from "./Login";
import "./loginStyle.css";

export const UserCheckTickets = () => {
    return (
        <div className="user-check-tickets">
            <div className="div-2">
                <Login className="login-instance" />
                <p className="p">
                    Login with your existing account to check your existing tickets
                </p>
            </div>
        </div>
    );
};
// function UserLogin() {
//     return (
//         <div className="user-login">
//             <UserCheckTickets />
//         </div>
//     );
// }

export default UserCheckTickets;