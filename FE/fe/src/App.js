import React from "react";
import RegistrationForm from "./Regist";
import ServiceDesk from "./serviceDesk";
import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import "./App.css";

function App() {
  // Example state for user login status and role
  const [user, setUser] = useState({ loggedIn: false, role: 'guest' }); // 'guest' is the default

  return (
    <Router>
      {" "}
      {/*added a router to use the Link component*/}
      <div>
        <header>
          <div className="header">
            <div className="overlap-group">
              <div className="register"> Register</div>
              <div className="text-wrapper">Check tickets</div>
              <div className="div">Create a Ticket</div>
              <div className="text-wrapper-2">Tech Support</div>
            </div>
          </div>
        </header>

        <Routes>
          {" "}
          {/* Added routes component*/}
          <Route path="/" element={<ServiceDesk />} />
          <Route path="/register" element={<RegistrationForm />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;

