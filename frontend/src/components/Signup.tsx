import { useState } from "react";
import type { ChangeEvent, FormEvent } from "react";
import { useNavigate } from "react-router-dom";
import { API_BASE } from "../api/config";

type Role = "doctor" | "receptionist";

interface SignupFormData {
    username: string;
    password: string;
    role: Role;
}

const Signup = () => {
    const [formData, setFormData] = useState<SignupFormData>({
        username: '',
        password: '',
        role: 'doctor'
    });

    const navigate = useNavigate()


    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData(prev => ({
            ...prev,
            [name]: value
        }));
    };

    const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        console.log("Form submitted:", formData);
        
        // make api request
        try {
            sendSignup(formData)

            // redirect to login
            navigate("/login")
            
        } catch (error) {
            console.error(error)
        }


        

    };

    const sendSignup = async (formData : SignupFormData) => {
        const response = await fetch(`${API_BASE}/auth/signup`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(formData)
        })

        const data = await response.json()
        return data
    }


    return (
        <>
            <h2 className="signup-title">Signup Form</h2>
            <form onSubmit={handleSubmit} className="signup-form">
                <div className="form-group">
                    <label htmlFor="username">Username: </label>
                    <input
                    type="text"
                    id="username"
                    name="username"
                    value={formData.username}
                    onChange={handleChange}
                    required
                    className="form-input"
                    />
                </div>
                <div className="form-group">
                    <label htmlFor="password">Password: </label>
                    <input
                    type="password"
                    id="password"
                    name="password"
                    value={formData.password}
                    onChange={handleChange}
                    required
                    className="form-input"
                    />
                </div>
                <div className="form-group role-group">
                    <label>Role: </label>
                    <label className="radio-label">
                    <input
                        type="radio"
                        name="role"
                        value="doctor"
                        checked={formData.role === "doctor"}
                        onChange={handleChange}
                    />
                    Doctor
                    </label>
                    <label className="radio-label">
                    <input
                        type="radio"
                        name="role"
                        value="receptionist"
                        checked={formData.role === "receptionist"}
                        onChange={handleChange}
                    />
                    Receptionist
                    </label>
                </div>
                <button type="submit" className="submit-btn">Submit</button>
            </form>

        </>
    );
};

export default Signup;
