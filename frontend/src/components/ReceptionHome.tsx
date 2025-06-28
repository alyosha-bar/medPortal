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


const ReceptionHome = () => {
    // import user info
    const { user, token } = useAuth();

    const [patients, setPatients] = useState<Patient[]>([])

    // view all patients
    const getAllPatients = async () => {
        const response = await fetch("/api/receptionist/patients", {
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


    // specific patient
    // register new patient
    
    
    // delete patient

    return ( 
        <>
            <div> Receptionist's Home Page </div>
            <div> {user?.username} </div>
            <div> {user?.role} </div>
            <h2>Patient List</h2>
            <ul>
                {patients.map((patient) => (
                <li key={patient.ID}>
                    {patient.Firstname} {patient.Lastname} — {patient.Age} years old ({patient.Gender})
                    <Link to={`/patient/${patient.ID}`}> Details </Link>
                </li>
                ))}
            </ul>
        </>
    );
}
 
export default ReceptionHome;