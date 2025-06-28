import { useState } from "react";
import type { ChangeEvent, FormEvent } from "react";
import { useNavigate } from "react-router-dom";
import type { Role } from "../types/auth";
import { useAuth } from "../hooks/useAuth";

interface LoginFormData {
    username: string;
    password: string;
}

const Login = () => {
    const [formData, setFormData] = useState<LoginFormData>({
        username: '',
        password: '',
    });

    const { login } = useAuth()
    const navigate = useNavigate(); 

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData(prev => ({
            ...prev,
            [name]: value,
        }));
    };

    const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        console.log("Login submitted:", formData);
        
        try {

            // if login fails --> catch block is triggered
            const userData = await sendLogin(formData)

            // handle the token and user data
            console.log(userData)
            const token = userData.token
            const user = userData.user as { id: number, username: string; role: Role };

            // pass into userContext and local storage
            login(user, token)

            if (user.role === "doctor") {
                // redirect to doctors home page
                navigate("/doctors")
            } else if (user.role === "receptionist") {
                // redirect to receptionist home page
                navigate("/reception")
            } else {
                navigate("/")
            }


        } catch (error) {
            console.error(error)
        }

        
    };

    const sendLogin = async (formData : LoginFormData) => {
        const response = await fetch(`/api/auth/login`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(formData)
        })

        const data = await response.json()
        return data
    }

    const { isAuthenticated } = useAuth()

    return (
        <>
            <h2 className="login-title">Login Form</h2>
            <form onSubmit={handleSubmit} className="login-form">
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
                <button type="submit" className="submit-btn">Login</button>
            </form>
        </>
    );
};

export default Login;
