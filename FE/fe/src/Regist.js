import React, { useState } from 'react';
import './Regist.css';

function RegistrationForm() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [formData, setFormData] = useState({
    email: '',
    username: '',
    password: '',
    confirmPassword: '',
    firstName: '',
    lastName: '',
    status: '', // Default value
    unit: '',
    noHp: '',
    captcha: '',
  });

    const [usernameAvailable, setUsernameAvailable] = useState(null);

  const handleInputChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
        });
        if (e.target.name === "username"){
            setUsernameAvailable(null);
        }
  };


    const checkUsername = async () => {
        try{
            const response = await fetch(`/api/check-username?username=${formData.username}`);
            const data = await response.json();
            setUsernameAvailable(data.available);
        } catch (error){
            console.error("Error checking username:", error)
            setUsernameAvailable(null);
        }
    }

  const handleSubmit = async (e) => {
    e.preventDefault();

        if (formData.password !== formData.confirmPassword){
            alert("Passwords do not match")
            return;
        }

      //captcha validation would go here

    try {
        const response = await fetch('/api/register', {
            method: 'POST',
            headers:{
                'Content-Type' : 'application/json'
            },
            body: JSON.stringify(formData),

        });

        if(response.ok){
            //registration successful
            alert("registration Successful");
            closeModal()
            //reset form
            setFormData({
                email: '',
                username: '',
                password: '',
                confirmPassword: '',
                firstName: '',
                lastName: '',
                status: '', 
                unit: '',
                noHp: '',
                captcha: '',
              });

        } else{
            const errorData = await response.json()
            console.error("Registration failed:", errorData);
            alert("Registration Failed" + errorData.message); 
        }

    } catch(error){
        console.error("Registration error", error);
        alert("An error occurred during registration.")
    }
  };

  const openModal = () => {
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
  };

  return (
      <div>
          <button onClick={openModal} style = {{backgroundColor: "green", color: "white", padding: "10px 20px", border: "none", borderRadius: "5px", cursor: "pointer"}}>Register</button>


      {isModalOpen && (
    <div className="modal-overlay" onClick={closeModal}>       
    <div className="modal-content" onClick={e => e.stopPropagation()}>
        <span className="close" onClick={closeModal}>×</span> {/* Close button */}
          <h2>Register an account at Service Desk | Layanan Aspirasi ITS</h2>
      <form onSubmit={handleSubmit}>
        <div>
            <label htmlFor="email">Email:</label>
            <input type="email" id="email" name="email" value={formData.email} onChange={handleInputChange} required />
        </div>
          <div>
              <label htmlFor = "username">Username:</label>
              <input type = "text" id="username" name="username" value = {formData.username} onChange = {handleInputChange} required onBlur = {checkUsername} />

              {usernameAvailable === true && <p style={{color: "green"}}>Username Available</p>}
              {usernameAvailable === false && <p style = {{color: "red"}}>Username is taken</p>}
              {usernameAvailable === null && formData.username !== ""&& <p style={{color:"grey"}}>Checking Username...</p> }
          </div>
        <div>
            <label htmlFor="password">Password:</label>
            <input type="password" id="password" name="password" value={formData.password} onChange={handleInputChange} required />
        </div>
        <div>
            <label htmlFor="confirmPassword">Confirm Password:</label>
            <input type="password" id="confirmPassword" name="confirmPassword" value={formData.confirmPassword} onChange={handleInputChange} required />
        </div>
        <div>
            <label htmlFor="firstName">First Name:</label>
            <input type="text" id="firstName" name="firstName" value={formData.firstName} onChange={handleInputChange} required />
          </div>
          <div> <label htmlFor="lastName">Last Name:</label>
          <input
              type="text"
              id="lastName"
              name="lastName"
              value={formData.lastName}
                onChange={handleInputChange}
                required
                />
          </div>
        <div>
            <label htmlFor="status">Status *:</label>
            <select id="status" name="status" value={formData.status} onChange={handleInputChange} required>
            <option value="Dosen">Dosen</option>
            <option value="Mahasiswa">Mahasiswa</option>
            {/* Add other status options as needed */}
            </select>
        </div>
        <div>
          <label htmlFor="unit">Unit *:</label>
            <input
                type="text"
                id="unit"
                name="unit"
              value={formData.unit}
              onChange={handleInputChange}
                required
            />
        </div>
        <div>
            <label htmlFor="noHp">No Hp/Telp *:</label>
            <input type="tel" id="noHp" name="noHp" value={formData.noHp} onChange={handleInputChange} required />
        </div>

          <div>
            <label>Captcha</label>
            <img src = "captcha.png" alt = "captcha"/>
              <input type = "text" name = "captcha" value={formData.captcha} onChange={handleInputChange}/>

          </div>

        <p>* = Required.</p>


        <button type="submit" style={{ backgroundColor: '#007bff', color: 'white', padding: '10px 20px', border: 'none', borderRadius: '5px', cursor: 'pointer' }}>Register</button>
      </form>

          <p style = {{marginTop: "20px"}}>Already have an account?  <button onClick = {closeModal} style = {{backgroundColor: "green", color: "white", padding: "5px 10px", border: "none", borderRadius: "5px", cursor: "pointer"}}>Login</button></p>
    </div>
    </div>
      )}
    </div>
  );
}

export default RegistrationForm;