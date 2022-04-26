import React from 'react';
import Background1 from "./img/bg-img/bg1.png"
import AuctionMain from "./img/bg-img/auctionMain.png"
import { Row, Col, Container } from 'reactstrap';
export default function HomePage() {

    return (
        <>
            <div>
                <section class="top_catagory_area d-md-flex clearfix">
                    <div class="single_catagory_area d-flex align-items-center bg-img" style={{ backgroundImage: `url(${Background1})` }} >

                    </div>

                    <div class="single_catagory_area d-flex align-items-center bg-img" style={{ backgroundColor: '#7f1aff' }}>
                        <div class="catagory-content">
                            <h2 style={{ color: 'white', fontWeight: 'bold' }}>Уникальная одежда <br />для уникальных людей</h2>
                            <h6 style={{ color: 'white', marginTop: '12px' }}>Покупай винтажные вещи с любимых <br /> магазинов прямо на сайте</h6>
                            <a href="#" class="btn karl-btn">Начать покупки</a>
                        </div>
                    </div>
                </section>
                <Container className='saleContainer'>
                    <div className='sale'>Категории</div>
                    <Row xs="auto">
                        <Col className='saleCol'>
                            <div class="product-card">
                                <img
                                    alt=""
                                    class="rectangle-2"
                                    src="https://via.placeholder.com/302x300"
                                />
                                <div class="description-product-card">
                                    <div class="descriptionSale">
                                        <p class="titleSale">Пиджак</p>
                                    </div>
                                </div>
                            </div>
                        </Col>
                        <Col>
                            <div class="product-card">
                                <img
                                    alt=""
                                    class="rectangle-2"
                                    src="https://via.placeholder.com/302x300"
                                />
                                <div class="description-product-card">
                                    <div class="descriptionSale">
                                        <p class="titleSale">Пиджак</p>
                                    </div>
                                </div>
                            </div>
                        </Col>
                        <Col>
                            <div class="product-card">
                                <img
                                    alt=""
                                    class="rectangle-2"
                                    src="https://via.placeholder.com/302x300"
                                />
                                <div class="description-product-card">
                                    <div class="descriptionSale">
                                        <p class="titleSale">Пиджак</p>
                                    </div>
                                </div>
                            </div>
                        </Col>

                    </Row>
                </Container>
                <Container className='saleContainer'>
                    <div className='sale'>Скидки дня</div>
                    <Row xs="auto">
                        <Col className='saleCol'>
                            <div class="product-card">
                                <img
                                    alt=""
                                    class="rectangle-3"
                                    src="https://via.placeholder.com/302x300"
                                />
                                <div class="description-product-card">
                                    <div class="descriptionSale">
                                        <p class="titleSale">Пиджак</p>
                                        <p class="maleFemale">Женская, S-M</p>
                                    </div>
                                    <div class="frame-3183486">
                                        <p class="firstNumber">9 990 ₸</p>
                                        <p class="secondNumber">12 990 ₸</p>
                                    </div>
                                </div>
                            </div>
                        </Col>
                        <Col>
                            <div class="product-card">
                                <img
                                    alt=""
                                    class="rectangle-3"
                                    src="https://via.placeholder.com/302x300"
                                />
                                <div class="description-product-card">
                                    <div class="descriptionSale">
                                        <p class="titleSale">Пиджак</p>
                                        <p class="maleFemale">Женская, S-M</p>
                                    </div>
                                    <div class="frame-3183486">
                                        <p class="firstNumber">9 990 ₸</p>
                                        <p class="secondNumber">12 990 ₸</p>
                                    </div>
                                </div>
                            </div>
                        </Col>
                        <Col>
                            <div class="product-card">
                                <img
                                    alt=""
                                    class="rectangle-3"
                                    src="https://via.placeholder.com/302x300"
                                />
                                <div class="description-product-card">
                                    <div class="descriptionSale">
                                        <p class="titleSale">Пиджак</p>
                                        <p class="maleFemale">Женская, S-M</p>
                                    </div>
                                    <div class="frame-3183486">
                                        <p class="firstNumber">9 990 ₸</p>
                                        <p class="secondNumber">12 990 ₸</p>
                                    </div>
                                </div>
                            </div>
                        </Col>
                        <Col>
                            <div class="product-card">
                                <img
                                    alt=""
                                    class="rectangle-3"
                                    src="https://via.placeholder.com/302x300"
                                />
                                <div class="description-product-card">
                                    <div class="descriptionSale">
                                        <p class="titleSale">Пиджак</p>
                                        <p class="maleFemale">Женская, S-M</p>
                                    </div>
                                    <div class="frame-3183486">
                                        <p class="firstNumber">9 990 ₸</p>
                                        <p class="secondNumber">12 990 ₸</p>
                                    </div>
                                </div>
                            </div>
                        </Col>
                    </Row>
                </Container>
                <Container className='auctionContainer'>
                    <section class="top_catagory_area d-md-flex clearfix secondSection">
                        <div class="auction-catagory-area d-flex align-items-center bg-img" style={{ backgroundColor: '#7f1aff' }}>
                            <div class="catagory-content">
                                <h2 style={{ color: 'white', fontWeight: 'bold' }}>Аукционные товары 500 ₸</h2>
                                <h6 style={{ color: 'white', marginTop: '12px' }}>Участвуй в аукционе и получай лучшие <br /> товары прямо на сайте</h6>
                                <a href="#" class="btn karl-btn">Перейти</a>
                            </div>
                        </div>
                        <div class="auction-png-catagory-area d-flex align-items-center bg-img" style={{ backgroundImage: `url(${AuctionMain})` }} >

                        </div>
                    </section>
                    {/* <section class="top_catagory_area d-md-flex clearfix secondSection ">
                        <div class="single_catagory_area d-flex align-items-center bg-img" style={{ backgroundColor: '#7f1aff' }}>
                            <div class="catagory-content">
                                <h2 style={{ color: 'white', fontWeight: 'bold' }}>Уникальная одежда <br />для уникальных людей</h2>
                                <h6 style={{ color: 'white', marginTop: '12px' }}>Покупай винтажные вещи с любимых <br /> магазинов прямо на сайте</h6>
                                <a href="#" class="btn karl-btn">Начать покупки</a>
                            </div>
                        </div>

                        <div class="single_catagory_area d-flex align-items-center bg-img" style={{ backgroundImage: `url(${Background1})` }} >

                        </div>
                    </section> */}
                </Container>

            </div>
        </>
    )
}
