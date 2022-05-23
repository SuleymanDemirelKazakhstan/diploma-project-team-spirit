import React from 'react';
// import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import './index.css'
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import './App.css';
import AppRouter from "./AppRouter";
function Navbar() {
    return (
        <Router>
            <div>
                <nav className="navbar navbar-expand-lg navbar-light fixed-top">
                    <div className="container">
                        <Link className="navbar-brand" to={"/"}>2ndChance.kz</Link>
                        <div className="collapse navbar-collapse" id="navbarTogglerDemo02" style={{textAlign: 'end'}}>
                            <ul className="navbar-nav ml-auto">
                                <li className="nav-item">
                                    <Link className="nav-link" to={"/login"}>Войти</Link>
                                </li>
                            </ul>
                        </div>
                    </div>
                </nav>
                <div className={'content'}>
                    <div className={"auth-wrapper"} >
                        <AppRouter />
                    </div>
                </div>
                <footer className="footer">
                    <div className="flex-wrapper-one">
                        <p className="num-2nd-chance-kz">2ndChance.kz</p>
                        <p className="footer-title">
                            Покупай винтажные вещи с любимых <br/>
                            магазинов прямо на сайте
                        </p>
                        <p
                            className="footer-rights"
                        >
                            ©2022 2ndChance.kz. Все права защищены.
                        </p>
                    </div>
                    <div className="footer-categories">
                        <a className="footer-titles">Категории</a>
                        <div className="flex-wrapper-two">
                            <div className="footer-categories-list">
                                <a className="footer-lists">Мужчинам</a>
                                <a className="footer-lists">Женщинам</a>
                                <a className="footer-lists">Аксессуары</a>
                            </div>
                            <div className="footer-categories-list">
                                <a className="footer-lists">Скидки</a>
                                <a className="footer-lists">Аукцион</a>
                            </div>
                        </div>
                    </div>
                    <div className="menu">
                        <a className="footer-titles">О платформе</a>
                        <a className="footer-lists">Свдения о платформе</a>
                        <a className="footer-lists">Контакты</a>
                        <a className="footer-lists">Магазины</a>
                    </div>
                    <div className="menu-two">
                        <a className="footer-titles">Авторизация</a>
                        <a className="footer-lists">Регистрация</a>
                        <a className="footer-lists">Войти</a>
                    </div>
                </footer>
            </div>
        </Router>
    );
}
export default Navbar;