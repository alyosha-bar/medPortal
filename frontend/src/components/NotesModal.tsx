import React, { useState } from "react";

interface NotesModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSubmit: (notes: string, patientID: number) => void;
  initialNotes: string;
  patientID: number;
}

const NotesModal: React.FC<NotesModalProps> = ({ isOpen, onClose, onSubmit, initialNotes, patientID }) => {
  const [notes, setNotes] = useState(initialNotes || "");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit(notes, patientID);
    onClose();
  };

  if (!isOpen) return null;

  return (
    <div className="modal-backdrop">
      <div className="modal">
        <h3>Update Medical Notes</h3>
        <form onSubmit={handleSubmit}>
          <textarea
            value={notes}
            onChange={(e) => setNotes(e.target.value)}
            rows={6}
            required
          />
          <div className="modal-buttons">
            <button className="save-btn" type="submit">Save</button>
            <button className="cancel-btn" type="button" onClick={onClose}>Cancel</button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default NotesModal;
