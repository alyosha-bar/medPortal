import { useEffect, useState } from "react";
import { useAuth } from "../hooks/useAuth";
import NotesModal from "./NotesModal";
import { API_BASE } from "../api/config";

interface Patient {
    ID: number,
    Firstname: string,
    Lastname: string,
    Age: number,
    Gender: string,
    MedicalNotes: string,
}

const DoctorHome = () => {

    // import token for api request
    const { token } = useAuth();

    const [patients, setPatients] = useState<Patient[]>([])
    const [showNotesModal, setShowNotesModal] = useState<boolean>(false)
    const [selectedPatient, setSelectedPatient] = useState<Patient | null>(null);

    const handleOpenNotesModal = (patient: Patient) => {
        setSelectedPatient(patient);
        setShowNotesModal(true);
    };



    const getAllPatients = async () => {
        const response = await fetch(`${API_BASE}/doctor/patients`, {
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

    // update medical notes fields
    const handleSaveNotes = async (updatedNotes: string, patientID: number) => {
        try {
            const response = await fetch(`${API_BASE}/doctor/patient/${patientID}`, {
                method: "PATCH",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${token}`
                },
                body: JSON.stringify({
                    medicalNotes: updatedNotes
                })
            })

            const data = await response.json()
            console.log(data)

        } catch ( error ) {
            console.error(error)
        }
    }
    return ( 
        <>  
            <h2 className="patient-list-heading">Your Patients</h2>
            <ul className="patient-list">
            {patients.map((patient) => (
                <li key={patient.ID} className="patient-item">
                <div className="diagnosis">
                    <span>
                    {patient.Firstname} {patient.Lastname} - {patient.Age} years old ({patient.Gender})
                    </span>
                    <span>
                    <strong>Notes: </strong> {patient.MedicalNotes}
                    </span>
                </div>

                <button
                    onClick={() => handleOpenNotesModal(patient)}
                    className="notes-btn"
                >
                    Medical Notes
                </button>
                </li>
            ))}
            </ul>


            {selectedPatient && (
                <NotesModal
                    isOpen={showNotesModal}
                    onClose={() => setShowNotesModal(false)}
                    onSubmit={handleSaveNotes} 
                    initialNotes={selectedPatient.MedicalNotes}
                    patientID={selectedPatient.ID}
                />
            )}
        </>
    );
}
 
export default DoctorHome;