import { useEffect, useState } from "react";
import { useAuth } from "../hooks/useAuth";
import { useNavigate, useParams } from "react-router-dom";
import AssignDoctorModal from "./AssignDoctorModal";
import UpdateFieldModal from "./UpdateFieldModal";

interface Patient {
    ID: number,
    Firstname: string,
    Lastname: string,
    Age: number,
    Gender: string
}

interface Doctor {
    ID: number,
    Username: string
}

const PatientDetails = () => {

    const navigate = useNavigate()
    const { token } = useAuth();

    let { id } = useParams()

    const [patient, setPatient] = useState<Patient>()
    const [doctors, setDoctors] = useState<Doctor[]>([])
    const [assignModal, setAssignModal] = useState<boolean>(false);
    const [updateModal, setUpdateModal] = useState<boolean>(false);

    // get specific patient details
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
    
    const getDoctorNames = async () => {
        const response = await fetch(`/api/receptionist/doctors`, {
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

    useEffect(() => {
        const fetchDoctors = async () => {
            const data = await getDoctorNames()
            setDoctors(data)
        }

        fetchDoctors()
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

    // assign doctor
    const handleAssign = async (doctorID : number) => {
        try {
            const response = await fetch(`/api/receptionist/patient/assign/${patient?.ID}`, {
                method: "PATCH",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${token}`
                },
                body: JSON.stringify({
                    doctorID: doctorID
                })
            })

            const data = await response.json()
            console.log(data)

            navigate("/reception")

        } catch ( error ) {
            console.error(error)
        }
    }


    // update simple fields
    const handleUpdate = async (field: string, value: string) => {
        try {
            const response = await fetch(`/api/receptionist//details/update/${patient?.ID}`, {
                method: "PATCH",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${token}`
                },
                body: JSON.stringify({
                    field: field,
                    value: value
                })
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
                        <button onClick={() => {
                            setUpdateModal(!updateModal)
                        }}className="update-btn">Update</button>
                        <button onClick={() => {
                            setAssignModal(!assignModal)
                        }} className="assign-btn">Assign</button>
                        <button onClick={() => handleDelete(patient.ID)} className="delete-btn">Delete</button>
                    </div>}

                    <AssignDoctorModal
                        isOpen={assignModal}
                        onClose={() => setAssignModal(false)}
                        doctors={doctors}
                        onAssign={(doctorId) => {
                            handleAssign(doctorId)  
                        }}
                    />

                    <UpdateFieldModal
                        isOpen={updateModal}
                        onClose={() => setUpdateModal(false)}
                        onSubmit={handleUpdate}
                    />
                </div>
            </div>
        </>
    );
}
 
export default PatientDetails;