import React from "react";
import "./admin-technician.css";

export const AdminTechnician = () => {
    return (
        <div className="admin-technician">
            <div className="group-wrapper">
                <div className="group">
                    <div className="overlap-group">
                        <p className="based-on-the">
                            Based on the category A,
                            <br />
                            here’s the list of available technicians
                        </p>

                        <div className="overlap">
                            <p className="technician-name">
                                Technician Name: <br />
                                Specialization: <br />
                                Telephone Number:
                                <br />
                                Email:
                            </p>
                        </div>

                        <div className="technician-name-wrapper">
                            <p className="text-wrapper">
                                Technician Name: <br />
                                Specialization: <br />
                                Telephone Number:
                                <br />
                                Email:
                            </p>
                        </div>

                        <div className="div-wrapper">
                            <p className="technician-name">
                                Technician Name: <br />
                                Specialization: <br />
                                Telephone Number:
                                <br />
                                Email:
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default AdminTechnician;