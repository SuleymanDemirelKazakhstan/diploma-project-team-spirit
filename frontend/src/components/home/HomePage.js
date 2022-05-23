import React from 'react';
import Background1 from "./img/bg-img/bg1.png"
import AuctionMain from "./img/bg-img/auctionMain.png"
import { Row, Col, Container } from 'reactstrap';
import CarouselSlideItem from '../details/CarouselSlideItem'
import { Link } from "react-router-dom";

const slideWidth = 30;
const _items = [
    {
        player: {
            title: 'Dala Bala',
            desc: 'Винтажная одежда и аксуссуары',
            image: 'https://www.openbusiness.ru/upload/iblock/8ab/shop_boutique_mannequin_clothes1.jpg',
        },
    },
    {
        player: {
            title: "ANR Vintage Shop",
            desc: "Селективный секондхэнд",
            image: 'https://uznayvse.ru/images/stories/uzn_1423790058.jpg',
        },
    },
    {
        player: {
            title: 'Second hand',
            desc: 'Самодельные изделия',
            image: 'https://vipidei.com/wp-content/uploads/2021/08/kak-otkryt-magazin-sekond-hend.jpeg',
        },
    },
    {
        player: {
            title: 'Vintage shop',
            desc: 'Одежды сэконд хэнд',
            image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSv4GWJB5HKWlgZxJqpQ_c2OTovkW_MgRP0dA&usqp=CAU',
        },
    },
    {
        player: {
            title: 'Second hand shop',
            desc: 'Одежды сэконд хэнд',
            image: 'https://kazved.ru/images/uploads/a511ab6b99304b0bdf634fec6478fb4f.jpg',
        },
    },
    {
        player: {
            title: 'Dala bala',
            desc: 'Одежды сэконд хэнд',
            image: 'https://megazigzag.ru/sites/default/files/b02689f81099a4fa7403_510x350crop.jpg',
        },
    },
];

const length = _items.length;
_items.push(..._items);

const sleep = (ms = 0) => {
    return new Promise((resolve) => setTimeout(resolve, ms));
};

const createItem = (position, idx) => {
    const item = {
        styles: {
            transform: `translateX(${position * 20}rem)`,
        },
        player: _items[idx].player,
    };
    console.log()
    switch (position) {
        case length - 1:
        case length + 1:
            item.styles = { ...item.styles };
            break;
        case length:
            break;
        default:
            item.styles = { ...item.styles };
            break;
    }

    return item;
};
const keys = Array.from(Array(_items.length).keys());

export default function HomePage() {


    const [items, setItems] = React.useState(keys);
    const [isTicking, setIsTicking] = React.useState(false);
    const [activeIdx, setActiveIdx] = React.useState(0);
    const bigLength = items.length;

    const prevClick = (jump = 1) => {
        if (!isTicking) {
            setIsTicking(true);
            setItems((prev) => {
                return prev.map((_, i) => prev[(i + jump) % bigLength]);
            });
        }
    };

    const nextClick = (jump = 1) => {
        if (!isTicking) {
            setIsTicking(true);
            setItems((prev) => {
                return prev.map(
                    (_, i) => prev[(i - jump + bigLength) % bigLength],
                );
            });
        }
    };

    const handleDotClick = (idx) => {
        if (idx < activeIdx) prevClick(activeIdx - idx);
        if (idx > activeIdx) nextClick(idx - activeIdx);
    };

    React.useEffect(() => {
        if (isTicking) sleep(300).then(() => setIsTicking(false));
    }, [isTicking]);

    React.useEffect(() => {
        setActiveIdx((length - (items[0] % length)) % length) // prettier-ignore
    }, [items]);

    return (
        <>
            <div>
                <section className="top_catagory_area d-md-flex clearfix">
                    <div className="single_catagory_area d-flex align-items-center bg-img" style={{ backgroundImage: `url(${Background1})` }} >

                    </div>

                    <div className="single_catagory_area d-flex align-items-center bg-img" style={{ backgroundColor: '#7f1aff' }}>
                        <div className="catagory-content">
                            <h2 style={{ color: 'white', fontWeight: 'bold' }}>Уникальная одежда <br />для уникальных людей</h2>
                            <h6 style={{ color: 'white', marginTop: '12px' }}>Покупай винтажные вещи с любимых <br /> магазинов прямо на сайте</h6>
                            <a href="#" className="btn karl-btn">Начать покупки</a>
                        </div>
                    </div>
                </section>
                <div className='saleContainer'>
                    <div className='sale'>Категории</div>
                    <Row xs="auto">
                        <Col className='column'>

                            <Link className="nav-link" to={"/main-catalog"}>
                                <div className="product-card">
                                    <img
                                        alt=""
                                        className="rectangle-2"
                                        src="https://i.imgur.com/1QWURYL.png"
                                    />
                                    <div className="description-product-card">
                                        <div className="descriptionSale">
                                            <p className="titleSale">Женское</p>
                                        </div>
                                    </div>
                                </div>
                            </Link>

                        </Col>
                        <Col className='column'>
                            <Link className="nav-link" to={"/main-catalog"}>
                                <div className="product-card">
                                    <img
                                        alt=""
                                        className="rectangle-2"
                                        src="https://i.imgur.com/V2tq6hA.png"
                                    />
                                    <div className="description-product-card">
                                        <div className="descriptionSale">
                                            <p className="titleSale">Мужское</p>
                                        </div>
                                    </div>
                                </div>
                            </Link>
                        </Col>
                        <Col className='column'>
                            <Link className="nav-link" to={"/main-catalog"}>
                                <div className="product-card">
                                    <img
                                        alt=""
                                        className="rectangle-2"
                                        src="https://i.imgur.com/uVoBmUp.png"
                                    />
                                    <div className="description-product-card">
                                        <div className="descriptionSale">
                                            <p className="titleSale">Аксессуары</p>
                                        </div>
                                    </div>
                                </div>
                            </Link>
                        </Col>

                    </Row>
                </div>
                <div className='saleContainer'>
                    <div className='sale'>Скидки дня</div>
                    <Row xs="auto">
                        <Col className='saleCol'>
                            <div className="product-card">
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
                            </div>
                        </Col>
                        <Col>
                            <div className="product-card">
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
                            </div>
                        </Col>
                        <Col>
                            <div className="product-card">
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
                            </div>
                        </Col>
                        <Col>
                            <div className="product-card">
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
                            </div>
                        </Col>
                    </Row>
                </div>

                <section className="top_catagory_area d-md-flex clearfix auctionContainer">

                    <div className="auction-catagory-area d-flex align-items-center bg-img" style={{ backgroundColor: '#7f1aff' }}>
                        <div className="catagory-content">
                            <h2 style={{ color: 'white', fontWeight: 'bold' }}>Аукционные товары 500 ₸</h2>
                            <h6 style={{ color: 'white', marginTop: '12px' }}>Участвуй в аукционе и получай лучшие <br /> товары прямо на сайте</h6>
                            <a href="#" className="btn karl-btn">Перейти</a>
                        </div>
                    </div>
                    <div className="auction-png-catagory-area d-flex align-items-center bg-img" style={{ backgroundImage: `url(${AuctionMain})` }} >

                    </div>
                </section>
                <div className='blockMagazines'>
                    <div className='magazines'>Магазины</div>
                    <div className="carousel__wrap">
                        <div className="carousel__inner">
                            <div className="carousel__container">
                                <ul className="carousel__slide-list">
                                    {items.map((pos, i) => (
                                        <CarouselSlideItem
                                            key={i}
                                            idx={i}
                                            pos={pos}
                                            activeIdx={activeIdx}
                                            createItem={createItem}
                                        />
                                    ))}
                                </ul>
                            </div>
                            <div className="carousel__dots">
                                {items.slice(0, length).map((pos, i) => (
                                    <button
                                        key={i}
                                        onClick={() => handleDotClick(i)}
                                        className={i === activeIdx ? 'dot active' : 'dot'}
                                    />
                                ))}
                            </div>
                        </div>
                    </div>
                </div>

            </div>
        </>
    )
}
