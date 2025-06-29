import { useEffect, useState } from "react";
import { useAuth } from "../hooks/useAuth";
import { Link } from "react-router-dom";
import RegisterModal from "./RegisterModal";
import { API_BASE } from "../api/config";

interface Patient {
    ID: number,
    Firstname: string,
    Lastname: string,
    Age: number,
    Gender: string
}


const ReceptionHome = () => {
    // import token for api request
    const { token } = useAuth();

    const [patients, setPatients] = useState<Patient[]>([])
    const [showRegisterModal, setShowRegisterModal] = useState<boolean>(false)


    // view all patients
    const getAllPatients = async () => {
        const response = await fetch(`${API_BASE}/receptionist/patients`, {
            method: "GET",
            headers: {
                "Content-type": "application/json",
                "Authorization": `Bearer ${token}`
            }
        })

        const data = await response.json()
        return data
    }

    useEffect(() => {
        const fetchPatients = async () => {
            const data = await getAllPatients()
            setPatients(data.data)
        }

        fetchPatients()
    }, [])

    
    // register new patient
    const registerPatient = async (newPatient: Omit<Patient, "ID">) => {
        const response = await fetch(`${API_BASE}/receptionist/register`, {
            method: "POST",
            headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(newPatient)
        });

        const data = await response.json();
        return data; 
    };



    const handleSubmit = async (newPatient: Omit<Patient, "ID">) => {
        try {
            const data = await registerPatient(newPatient);
            console.log(data);
        } catch (error) {
            console.error("Registration failed:", error);
        }
    };


    return ( 
        <>  
        <div>
            <div className="patient-list-top">
                <h2 className="patient-list-heading">Patient List</h2>
                <button 
                    onClick={() => setShowRegisterModal(!showRegisterModal)}
                    className="register-btn">Register New Patient
                </button>
                {showRegisterModal && (
                    <RegisterModal
                        onClose={() => setShowRegisterModal(false)}
                        onSubmit={handleSubmit}
                    />
                    )}
            </div>
            <ul className="patient-list">
            {patients && patients.map((patient) => (
                <li key={patient.ID} className="patient-item">
                <span>
                    {patient.Firstname} {patient.Lastname} - {patient.Age} years old ({patient.Gender})
                </span>
                <Link to={`/patient/${patient.ID}`} className="details-link">Details</Link>
                </li>
            ))}
            </ul>
        </div>
        </>
    );
}
 
export default ReceptionHome;