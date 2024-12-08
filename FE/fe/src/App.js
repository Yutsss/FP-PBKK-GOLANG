import React, { useState } from "react";
import UserRegister from "./Regist";
import ServiceDesk from "./serviceDesk";
import UserCheckTickets from "./userLogin";
import UserCheckTicketsLoggedIn from "./checkTicket";
import TechHome from "./tech-home";
import TechnicianTickets from "./tech-tickets";
import AdminHome from "./admin-home";
import AdminTicket from "./admin-ticket";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Link,
  useNavigate,
} from "react-router-dom";

function App() {
  // Example state for user login status and role
  const [user, setUser] = useState({ loggedIn: false, role: "guest" }); // 'guest' is the default
  const navigate = useNavigate(); // Use React Router's navigate function

  // Function to handle "Check Tickets" logic
  const handleCheckTickets = () => {
    if (user.loggedIn) {
      navigate("/check-tickets"); // Navigate to check tickets if logged in
    } else {
      navigate("/login"); // Navigate to login page if not logged in
    }
  };

  return (
    <div style={{ backgroundColor: "#0d0911" }}>
      <header
        style={{
          backgroundColor: "#0d0911",
          color: "white",
          padding: "10px",
          border: "1px solid",
          borderColor: "#b892d8",
          marginTop: "-1px",
          marginLeft: "-1px",
          marginRight: "-1px",
        }}
      >
        <div
          style={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
          }}
        >
          <div style={{ display: "flex", alignItems: "center" }}>
            <img src="logo192.png" alt="Logo" style={{ height: "50px" }} />
            <h1>Tech Support</h1>
          </div>
          <nav>
            {/* Button for Check Tickets, redirecting to the appropriate page based on login status */}
            <Link
              to={user.loggedIn ? "/check-tickets" : "/login"} // Conditionally redirect based on login status
              style={{
                color: "white",
                padding: "5px 10px",
                borderRadius: "5px",
                textDecoration: "none",
              }}
            >
              Check Tickets
            </Link>

            <Link
              to="/"
              style={{
                color: "white",
                padding: "5px 10px",
                borderRadius: "5px",
                textDecoration: "none",
              }}
            >
              Create Ticket
            </Link>
            <Link
              to="/tech-home"
              style={{
                color: "white",
                padding: "5px 10px",
                borderRadius: "5px",
                textDecoration: "none",
              }}
            >
              Technician Home
            </Link>
            <Link
              to="/tech-tickets"
              style={{
                color: "white",
                padding: "5px 10px",
                borderRadius: "5px",
                textDecoration: "none",
              }}
            >
              Technician Tickets
            </Link>

            <Link
              to="/admin-home"
              style={{
                color: "white",
                padding: "5px 10px",
                borderRadius: "5px",
                textDecoration: "none",
              }}
            >
              Admin Home
            </Link>
            <Link
              to="/admin-tickets"
              style={{
                color: "white",
                padding: "5px 10px",
                borderRadius: "5px",
                textDecoration: "none",
              }}
            >
              Admin Tickets
            </Link>

            {/* Conditional rendering based on login state */}
            {!user.loggedIn ? (
              <Link
                to="/login"
                style={{
                  color: "white",
                  padding: "5px 10px",
                  borderRadius: "5px",
                  textDecoration: "none",
                }}
              >
                Login
              </Link>
            ) : (
              <Link
                to={`/${user.role}`}
                style={{
                  color: "white",
                  padding: "5px 10px",
                  borderRadius: "5px",
                  textDecoration: "none",
                }}
              >
                {user.role === "admin"
                  ? "Admin"
                  : user.role === "technician"
                  ? "Technician"
                  : "User"}
              </Link>
            )}
          </nav>
        </div>
      </header>

      {/* Added routes component */}
      <Routes>
        <Route path="/" element={<ServiceDesk />} />
        <Route path="/register" element={<UserRegister />} />
        <Route path="/login" element={<UserCheckTickets />} />
        {/* Dynamic routes for different user roles */}
        <Route path="/user" element={<div>User Dashboard</div>} />
        <Route path="/admin" element={<div>Admin Dashboard</div>} />
        <Route path="/technician" element={<div>Technician Dashboard</div>} />
        <Route path="/check-tickets" element={<UserCheckTicketsLoggedIn />} />
        <Route path="/tech-home" element={<TechHome />} />
        <Route path="/tech-tickets" element={<TechnicianTickets />} />
        <Route path="/admin-home" element={<AdminHome />} />
        <Route path="/admin-tickets" element={<AdminTicket />} />
      </Routes>
    </div>
  );
}

export default function AppWrapper() {
  return (
    <Router>
      {" "}
      {/* Wrap App component in Router */}
      <App />
    </Router>
  );
}
