export type Role = "doctor" | "receptionist";

export interface User {
    id: number;
    username: string;
    role: Role;
}

export interface AuthContextType {
    user: User | null;
    token: string | null;
    login: (token: string, user: User) => void;
    logout: () => void;
    isAuthenticated: boolean;
    isLoading: boolean;
}
