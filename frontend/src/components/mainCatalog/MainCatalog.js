import axios from 'axios';
import React, { useEffect } from 'react';
import { Row, Col, Container, Button } from 'reactstrap';
import { Link } from "react-router-dom";
export default function MainCatalog() {

    const list = ["Платья", "Верхняя одежда", "Блузы и топы"]
    const list1 = ["Верхняя одежда", "Пиджаки", "Свитера и толстовки", "Обувь", "Другое"]
    const listPr = ["Платья", "Верхняя одежда", "Блузы и топы"]
    // function getInitial() {
    //     axios.get("http://192.168.146.233:3000/g/allproduct").then((response) => {
    //         console.log(response)
    //     })
    // }
    // useEffect(() => {
    //     getInitial()
    // }, [])
    const listItems = list.map((item) =>
        <li className="list-item">
            <div className="list-label-inner">{item}</div></li>
    );
    const listItems1 = list1.map((item) =>
        <li className="list-item">
            <div className="list-label-inner">{item}</div></li>
    );
    const listProduct = listPr.map((item) =>
        <Link className="nav-link" to={"/detail-page"}>
            <div className="product-card-main-catalog">
                <img
                    alt=""
                    className="rectangle-3"
                    src="https://i.imgur.com/EYaFLBC.png"
                />
                <div className="description-product-card">
                    <div className="descriptionSale">
                        <p className="titleSale">Пиджак</p>
                        <p className="maleFemale">Женская, S-M</p>
                    </div>
                    <div className="frame-3183486">
                        <p className="firstNumber">9 990 ₸</p>
                        <p className="secondNumber">12 990 ₸</p>
                    </div>
                </div>

                <Button className='buy-btn'>Купить сейчас</Button>
            </div>
        </Link>

    );
    return (
        <>
            <div className="containerMainCatalog">
                <nav className="navigation">
                    <ul className="list">
                        <li className="list-item">
                            <div className="list-label">Женщинам</div>
                            <ul className="list">
                                {listItems}
                            </ul>
                        </li>
                    </ul>
                    <ul className="list">
                        <li className="list-item">
                            <div className="list-label">Мужчинам</div>
                            <ul className="list">
                                {listItems1}
                            </ul>
                        </li>
                    </ul>
                </nav>
                <div className='mainCatalog'>
                    <div className='produc-main-catalog-title'>Платья</div>
                    <div className='listProducts'>
                        {listProduct}
                        {listProduct}
                    </div>
                </div>
            </div>

        </>
    )
}
