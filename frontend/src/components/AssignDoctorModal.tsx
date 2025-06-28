import React, { useState } from "react";

interface Doctor {
  ID: number;
  Username: string;
}

interface AssignDoctorModalProps {
  isOpen: boolean;
  onClose: () => void;
  doctors: Doctor[];
  onAssign: (doctorId: number) => void;
}

const AssignDoctorModal: React.FC<AssignDoctorModalProps> = ({
  isOpen,
  onClose,
  doctors,
  onAssign,
}) => {
  const [selectedDoctorId, setSelectedDoctorId] = useState<number | null>(null);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (selectedDoctorId !== null) {
      onAssign(selectedDoctorId);
      onClose(); // close the modal after assignment
    }
  };

  if (!isOpen) return null;

  return (
    <div className="modal-overlay">
      <div className="modal">
        <h2>Assign Doctor</h2>
        <form onSubmit={handleSubmit}>
          <label htmlFor="doctor">Select a Doctor:</label>
          <select
            id="doctor"
            value={selectedDoctorId ?? ""}
            onChange={(e) => setSelectedDoctorId(Number(e.target.value))}
            required
          >
            <option value="" disabled>Select a doctor</option>
            {doctors.map((doc) => (
              <option key={doc.ID} value={doc.ID}>
                {doc.Username}
              </option>
            ))}
          </select>

          <div className="btn-group">
            <button type="submit" className="assign-btn">Assign</button>
            <button type="button" className="cancel-btn" onClick={onClose}>Cancel</button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default AssignDoctorModal;
