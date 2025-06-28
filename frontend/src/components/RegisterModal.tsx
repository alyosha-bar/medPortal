import { useState } from "react";

interface Patient {
  ID: number;
  Firstname: string;
  Lastname: string;
  Age: number;
  Gender: string;
}

interface Props {
  onClose: () => void;
  onSubmit: (patient: Omit<Patient, "ID">) => void;
}

const RegisterModal = ({ onClose, onSubmit }: Props) => {
  const [formData, setFormData] = useState<Omit<Patient, "ID">>({
    Firstname: "",
    Lastname: "",
    Age: 0,
    Gender: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: name === "Age" ? Number(value) : value,
    }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit(formData);
    onClose();
  };

  return (
    <div className="modal-overlay">
      <div className="modal">
        <button className="close-btn" onClick={onClose}>×</button>
        <h2>Register New Patient</h2>
        <form onSubmit={handleSubmit} className="patient-form">
          <input name="Firstname" placeholder="First Name" value={formData.Firstname} onChange={handleChange} required />
          <input name="Lastname" placeholder="Last Name" value={formData.Lastname} onChange={handleChange} required />
          <input type="number" name="Age" placeholder="Age" value={formData.Age} onChange={handleChange} required />
          <select name="Gender" value={formData.Gender} onChange={handleChange} required>
            <option value="">Select Gender</option>
            <option value="M">Male</option>
            <option value="F">Female</option>
          </select>
          <button type="submit">Register</button>
        </form>
      </div>
    </div>
  );
};

export default RegisterModal;