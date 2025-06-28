import React, { useState } from "react";

interface UpdateFieldModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSubmit: (field: string, value: string) => void;
}

const fields = [
  { value: "firstname", label: "First Name" },
  { value: "lastname", label: "Last Name" },
  { value: "age", label: "Age" },
  { value: "gender", label: "Gender" },
  { value: "medical_notes", label: "Medical Notes" },
];

const UpdateFieldModal: React.FC<UpdateFieldModalProps> = ({ isOpen, onClose, onSubmit }) => {
  const [selectedField, setSelectedField] = useState(fields[0].value);
  const [value, setValue] = useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit(selectedField, value);
    setValue("");
    setSelectedField(fields[0].value);
    onClose();
  };

  if (!isOpen) return null;

  return (
    <div className="modal-backdrop">
      <div className="modal">
        <h3>Update Patient Field</h3>
        <form onSubmit={handleSubmit}>
          <label>
            Field:
            <select value={selectedField} onChange={(e) => setSelectedField(e.target.value)}>
              {fields.map((field) => (
                <option key={field.value} value={field.value}>
                  {field.label}
                </option>
              ))}
            </select>
          </label>
          <label>
            New Value:
            <input
              type="text"
              value={value}
              onChange={(e) => setValue(e.target.value)}
              required
            />
          </label>
          <div className="modal-buttons">
            <button type="submit">Submit</button>
            <button type="button" onClick={onClose}>Cancel</button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default UpdateFieldModal;
