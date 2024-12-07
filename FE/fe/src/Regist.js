import React, { useState } from 'react';
import { useCallback } from 'react';
import { useNavigate } from "react-router-dom";
import styles from './Regist.module.css';


const UserRegister = () => {
  const navigate = useNavigate();

  const onRegisterTextClick = useCallback(() => {
    navigate("/");
  }, [navigate]);


  const onCheckTicketsTextClick = useCallback(() => {
    // Add your code here
  }, []);

  // State hooks to store user input
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    password: '',
    phoneNumber: '',
    address: '',
  });

  // Handle input change
  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  // Handle form submission
  const handleSubmit = (e) => {
    e.preventDefault();
    // Here you can process the data, e.g., send it to an API or update the state.
    console.log(formData);
  };

  return (
    <div className={styles.userRegister}>
      <div className={styles.register1}>Register</div>
      <form onSubmit={handleSubmit} className={styles.formRegister}>
        <div className={styles.inputField}>
          <div className={styles.label}>Name</div>
            <input
              type="text"
              name="name"
              value={formData.name}
              onChange={handleInputChange}
              className={styles.value}
              required
            />
        </div>

        <div className={styles.inputField}>
          <div className={styles.label}>Email</div>
            <input
              type="email"
              name="email"
              value={formData.email}
              onChange={handleInputChange}
              className={styles.value}
              required
            />
        </div>

        <div className={styles.inputField}>
          <div className={styles.label}>Password</div>
            <input
              type="password"
              name="password"
              value={formData.password}
              onChange={handleInputChange}
              className={styles.value}
              required
            />
        </div>

        <div className={styles.inputField}>
          <div className={styles.label}>Phone Number</div>
            <input
              type="tel"
              name="phoneNumber"
              value={formData.phoneNumber}
              onChange={handleInputChange}
              className={styles.value}
              required
            />
        </div>

        <div className={styles.inputField}>
          <div className={styles.label}>Address</div>
            <textarea
              name="address"
              value={formData.address}
              onChange={handleInputChange}
              className={styles.value}
              required
            />
        </div>
        <button className={styles.button}>Register</button>
      </form>
    </div>
  );
};

export default UserRegister;
