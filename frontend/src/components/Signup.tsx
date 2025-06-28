import { useState } from "react";
import type { ChangeEvent, FormEvent } from "react";
import { redirect } from "react-router-dom";

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
            redirect("/login")
            
        } catch (error) {
            console.error(error)
        }


        

    };

    const sendSignup = async (formData : SignupFormData) => {
        const response = await fetch(`/api/auth/signup`, {
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
            <h2>Signup Form</h2>
            <form onSubmit={handleSubmit}>
                <div>
                    <label htmlFor="username">Username: </label>
                    <input
                        type="text"
                        id="username"
                        name="username"
                        value={formData.username}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <label htmlFor="password">Password: </label>
                    <input
                        type="password"
                        id="password"
                        name="password"
                        value={formData.password}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <label>Role: </label>
                    <label>
                        <input
                            type="radio"
                            name="role"
                            value="doctor"
                            checked={formData.role === "doctor"}
                            onChange={handleChange}
                        />
                        Doctor
                    </label>
                    <label>
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
                <button type="submit">Submit</button>
            </form>
        </>
    );
};

export default Signup;
