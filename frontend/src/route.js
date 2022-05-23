import HomePage from "./components/home/HomePage";
import LoginRegistrationMain from "./components/LoginRegistrationMain";
import DetailPage from "./components/mainCatalog/DetailPage";
import MainCatalog from "./components/mainCatalog/MainCatalog";
import DetailOplata from "./components/mainCatalog/DetailOplata";

export const authRoutes = [
    {
        path: "/",
        Component: HomePage
    },
    {
        path: "/main-catalog",
        Component: MainCatalog
    },
    {
        path: "/detail-page",
        Component: DetailPage
    },
    {
        path: "/detail-payments",
        Component: DetailOplata
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