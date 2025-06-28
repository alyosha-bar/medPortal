import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import type { Role } from "../types/auth";

interface ProtectedRouteProps {
  requiredRole?: Role;
  redirectTo?: string;
}

export const ProtectedRoute = ({ requiredRole, redirectTo = "/login" }: ProtectedRouteProps) => {
  const { isAuthenticated, user, isLoading } = useAuth();

  if (isLoading) {
    return <div>Loading...</div>; // or null/spinner
  }

  if (!isAuthenticated) {
    return <Navigate to={redirectTo} replace />;
  }

  if (requiredRole && user?.role !== requiredRole) {
    return <Navigate to="/" replace />;
  }

  return <Outlet />;
};
