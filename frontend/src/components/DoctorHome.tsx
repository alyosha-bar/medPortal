import { useEffect, useState } from "react";
import { useAuth } from "../hooks/useAuth";
import { Link } from "react-router-dom";

interface Patient {
    ID: number,
    Firstname: string,
    Lastname: string,
    Age: number,
    Gender: string
}

const DoctorHome = () => {

     // import token for api request
    const { token } = useAuth();

    const [patients, setPatients] = useState<Patient[]>([])
    // view patients
    // view all patients
    const getAllPatients = async () => {
        const response = await fetch("/api/doctor/patients", {
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
            setPatients(data)
        }

        fetchPatients()
    }, [])

    // update data

    // export patients list into its own component and use here
    // see details, based on role should or should not render delete & assign buttons

    return ( 
        <>  
            <h2 className="patient-list-heading">Your Patients</h2>
            <ul className="patient-list">
            {patients.map((patient) => (
                <li key={patient.ID} className="patient-item">
                <span>
                    {patient.Firstname} {patient.Lastname} — {patient.Age} years old ({patient.Gender})
                </span>
                <Link to={`/patient/${patient.ID}`} className="details-link">Details</Link>
                </li>
            ))}
            </ul>
        </>
    );
}
 
export default DoctorHome;