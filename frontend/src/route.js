import HomePage from "./components/home/HomePage";
import LoginRegistrationMain from "./components/LoginRegistrationMain";

export const authRoutes = [
    {
        path: "/home",
        Component: HomePage
    },
    {
        path: "/login",
        Component: LoginRegistrationMain
    }
]

export const publicRoute = [
    {
        path: "/login",
        Component: LoginRegistrationMain
    }
]