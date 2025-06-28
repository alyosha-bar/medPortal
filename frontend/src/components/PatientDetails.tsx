import { useEffect, useState } from "react";
import { useAuth } from "../hooks/useAuth";
import { useNavigate, useParams } from "react-router-dom";

interface Patient {
    ID: number,
    Firstname: string,
    Lastname: string,
    Age: number,
    Gender: string
}

const PatientDetails = () => {

    const navigate = useNavigate()
    const { token } = useAuth();

    let { id } = useParams()

    const [patient, setPatient] = useState<Patient>()

    // view all patients
    const getPatientDetails = async () => {
        const response = await fetch(`/api/receptionist/patient/${id}`, {
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
        const fetchPatient = async () => {
            const data = await getPatientDetails()
            setPatient(data)
        }

        fetchPatient()
    }, [])

    return (
        <>
            <div>
                <button onClick={() => navigate(-1)}>Back</button>
                <h2>Patient Details</h2>
                <p><strong>ID:</strong> {patient?.ID}</p>
                <p><strong>Name:</strong> {patient?.Firstname} {patient?.Lastname}</p>
                <p><strong>Age:</strong> {patient?.Age}</p>
                <p><strong>Gender:</strong> {patient?.Gender}</p>
            </div>
        </>
    );
}
 
export default PatientDetails;