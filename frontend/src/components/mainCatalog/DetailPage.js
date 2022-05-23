import axios from 'axios';
import React, { useEffect } from 'react';
import { Row, Col, Container, Button } from 'reactstrap';
import { Link } from "react-router-dom";

export default function DetailPage() {
    const listPr = ["Платья", "Верхняя одежда", "Блузы и топы", "Блузы и топы"]
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
                <Link className="nav-link" to={"/detail-payments"}>
                    <Button className='buy-btn'>Купить сейчас</Button>
                </Link>
            </div>
        </Link>

    );
    return (
        <>
            <div className="container mb-4 detailPage">
                <div className="row no-gutters mt-4">
                    <div className="col-12 col-md-5">
                        <div className="row no-gutters w-100">
                            <div className=" col-12 col-lg-2 flex-lg-first mb-md-0 mb-4" style={{ display: "block" }}>
                                <div className="row no-gutters ">

                                    <div className="col-4 col-sm-3 col-lg-12 border-full-1px-solid border-color-a0 justify-content-center side-picture active" role="button" data-target="side-1">
                                        <div className="p-2 w-75">
                                            <img src="https://i.imgur.com/tXMRWjH.png" className="img-fluid " />
                                        </div>
                                    </div>

                                    <div className="col-4 col-sm-3 col-lg-12 border-full-1px-solid border-color-a0 justify-content-center side-picture" role="button" data-target="side-2">
                                        <div className="p-2 w-75">
                                            <img src="https://i.imgur.com/NwMV5GE.png" className="img-fluid " />
                                        </div>
                                    </div>
                                    <div className="col-4 col-sm-3 col-lg-12 border-full-1px-solid border-color-a0 justify-content-center side-picture" role="button" data-target="side-3">
                                        <div className="p-2 w-75">
                                            <img src="https://i.imgur.com/Ahn3E6G.png" className="img-fluid " />
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div className="col-12 col-lg-10 flex-lg-last border-full-1px-solid border-color-a0">
                                <div className="d-flex align-items-center w-100">
                                    <img src="https://i.imgur.com/bi6s0Wm.png" className="img-fluid" id="main-image" style={{ maxHeight: "600px" }} />
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="col-12 col-md-5">
                        <div className="pt-0 pr-2 pb-2 pl-1 pl-md-5">
                            <div className="mb-4 border-full-2px-solid border-top-0 border-left-0 border-right-0 border-color-inverse">
                                <h1 style={{ fontSize: '28px', fontWeight: "500" }}>Бархатное платье с открытой спинкой</h1>
                                <p style={{ fontSize: '28px', fontWeight: "400" }}>600 ₸</p>
                            </div>

                            {/* <div className="d-flex align-content-center py-2 mb-4 ">
                            <ul className="list-unstyled list-inline ">
                                <li className="list-inline-item "><img src="img/icon/facebook.png " width="25px " alt="fb" /></li>
                                <li className="list-inline-item "><img src="img/icon/twitter.png " width="25px " alt="twt" /></li>
                                <li className="list-inline-item "><img src="img/icon/google-plus.png " width="25px " alt="G+" /></li>
                                <li className="list-inline-item "><img src="img/icon/linkedin.png " width="25px " alt="Ln" /></li>
                                <li className="list-inline-item "><img src="img/icon/pinterest.png " width="25px " alt="P" /></li>
                            </ul>
                        </div> */}

                            <Button className='buy-btn-product'>Купить сейчас</Button>
                            <div class="characteristics">
                                <p class="razmer-xs-l">
                                    Размер:<strong class="razmer-xs-l-emphasis-1">
                                        XS-L</strong
                                    >
                                </p>
                                <p class="razmer-xs-l">
                                    Состояние:<strong class="razmer-xs-l-emphasis-1">
                                        Б\У</strong
                                    >
                                </p>
                                <p class="magazin-robin-clothes">
                                    <strong class="magazin-robin-clothes-emphasis-0"
                                    >Магазин:</strong
                                    >
                                    Robin.clothes
                                </p>
                            </div>
                            <div className="mb-5 ">
                                    <div className="my-5">
                                        <ul className="nav nav-tabs" role="tablist">
                                            <li className="nav-item">
                                                <a className="nav-link text-color-a1  active" data-toggle="tab" href="#deskripsi" role="tab">Описание</a>
                                            </li>
                                            <li className="nav-item">
                                                <a className="nav-link text-color-a1" data-toggle="tab" href="#detail" role="tab">О магазине</a>
                                            </li>
                                            <li className="nav-item ">
                                                <a className="nav-link text-color-a1" data-toggle="tab" href="#shipping" role="tab">Условия</a>
                                            </li>
                                        </ul>
                                        <div className="tab-content font-size-09" style={{ backgroundColor: "#fafafa" }}>
                                            <div className="tab-pane fade show active px-3 py-4 text-color-a2" id="deskripsi" role="tabpanel">
                                            Бархатное синее платье с открытой спинкой. Материал приятный телу. Состояние отличное . Состав: 100% полиэстер.
                                            </div>
                                            <div className="tab-pane fade px-3 py-4 text-color-a2" id="detail" role="tabpanel">Detail</div>
                                            <div className="tab-pane fade px-3 py-4 text-color-a2" id="shipping" role="tabpanel">,,,</div>
                                            <div className="tab-pane fade px-3 py-4 text-color-a2" id="other" role="tabpanel">...</div>
                                        </div>
                                    </div>
                                {/* <select className="custom-select ml-0 mb-1 font-weight-bold " style={{ borderRadius: "0", width: "100px", border: "1px solid #bfb6bb " }}>
                                    <option selected>SIZE</option>
                                    <option value="1 ">XS</option>
                                    <option value="2 ">S</option>
                                    <option value="3 ">M</option>
                                    <option value="4 ">L</option>
                                    <option value="5 ">XL</option>
                                    <option value="6 ">XXL</option>
                                </select>
                                <a className=" ">
                                    <h6 className="text-muted font-size-08 text-uppercase "><u>Size Chart</u></h6>
                                </a> */}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div className=" mb-4 ">
                <h4 className="text-center text-uppercase ">Другие предложения</h4>
                <div className="row no-gutters gutter-20 mb-20px">
                <div className='mainCatalog'>
                    <div className='listProducts'>
                        {listProduct}
                    </div>
                </div>
                    {/* <div className="col-6 col-md-3 ">
                        <div style={{ border: "2px solid #fafafa" }} className="w-100">
                            <div className="d-flex justify-content-center align-items-center">
                                <img className="img-fluid" src="https://dummyimage.com/400x600/e66a63/fff" alt=" " />
                            </div>
                            <div className="text-center py-4 ">
                                <h6>Crocodile Leather Jacket</h6>
                                <h6 className="text-muted font-weight-normal ">$ 299.00</h6>
                            </div>
                        </div>
                    </div>
                    <div className="col-6 col-md-3 ">
                        <div style={{ border: "2px solid #fafafa" }} className="w-100">
                            <div className="d-flex justify-content-center align-items-center">
                                <img className="img-fluid" src="https://dummyimage.com/400x600/e66a63/fff" alt=" " />
                            </div>
                            <div className="text-center py-4 ">
                                <h6>Crocodile Leather Jacket</h6>
                                <h6 className="text-muted font-weight-normal ">$ 299.00</h6>
                            </div>
                        </div>
                    </div>
                    <div className="col-6 col-md-3 ">
                        <div style={{ border: "2px solid #fafafa" }} className="w-100">
                            <div className="d-flex justify-content-center align-items-center">
                                <img className="img-fluid" src="https://dummyimage.com/400x600/e66a63/fff" alt=" " />
                            </div>
                            <div className="text-center py-4 ">
                                <h6>Crocodile Leather Jacket</h6>
                                <h6 className="text-muted font-weight-normal ">$ 299.00</h6>
                            </div>
                        </div>
                    </div>
                    <div className="col-6 col-md-3 ">
                        <div style={{ border: "2px solid #fafafa" }} className="w-100">
                            <div className="d-flex justify-content-center align-items-center">
                                <img className="img-fluid" src="https://dummyimage.com/400x600/e66a63/fff" alt=" " />
                            </div>
                            <div className="text-center py-4 ">
                                <h6>Crocodile Leather Jacket</h6>
                                <h6 className="text-muted font-weight-normal ">$ 299.00</h6>
                            </div>
                        </div>
                    </div> */}
                </div>
            </div>
        </>
    )
}