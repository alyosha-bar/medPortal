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

    // delete patient profile
    const handleDelete = async (ID: number) => {
        try {
            const response = await fetch(`/api/receptionist/patient/${ID}`, {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${token}`
                }
            })

            const data = await response.json()
            console.log(data)

            navigate("/reception")

        } catch ( error ) {
            console.error(error)
        }
    }

    return (
        <>
            <div>
                <button className="back-btn" onClick={() => navigate(-1)}>Back</button>
                <div className="patient-details">
                    <h2>Patient Details</h2>
                    <p><strong>ID:</strong> {patient?.ID}</p>
                    <p><strong>Name:</strong> {patient?.Firstname} {patient?.Lastname}</p>
                    <p><strong>Age:</strong> {patient?.Age}</p>
                    <p><strong>Gender:</strong> {patient?.Gender}</p>

                    {patient && 
                    <div className="btn-group">
                        <button className="update-btn">Update</button>
                        <button className="assign-btn">Assign</button>
                        <button onClick={() => handleDelete(patient.ID)} className="delete-btn">Delete</button>
                    </div>}
                </div>
            </div>
        </>
    );
}
 
export default PatientDetails;