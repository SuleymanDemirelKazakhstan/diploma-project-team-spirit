import React from 'react';
import { Route, Switch } from "react-router-dom"
import { authRoutes } from "./route"
// import AuthService from "./Services/auth/auth.service"

const AppRouter = () => {
    // const user = AuthService.getCurrentUser()

    return (
        <Switch>
            {authRoutes.map(({ path, Component, state }) =>
                <Route key={path} path={path} component={Component} state={state} exact />
            )}

        </Switch>
    )
}

export default AppRouter;