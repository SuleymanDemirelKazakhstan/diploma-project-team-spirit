import React, { useState } from "react";
import '../../index.css'
import { Input, InputGroup, InputGroupText, Label } from "reactstrap";
import InputMask from 'react-input-mask';
import { BsTelephone, BsLock } from 'react-icons/bs';
export default function Login() {

    const emptyUsers = {
        phone: '+7',
        password: ''
    }
    const [loginDetails, setLoginDetails] = useState(emptyUsers)

    function handleChange(event) {
        setLoginDetails({ ...loginDetails, [event.target.name]: event.target.value })
    }

    function handleSubmit(event) {
        console.log(loginDetails)
    }
    return (
        <>
            <div className="auth-inner">
                <form>

                    <h3 className="authorization">Авторизация</h3>
                    <p className="s16">Введите данные ниже для входа</p>
                    <div className="form-group f1">
                        <p className="label">Номер телефона</p>
                        <InputGroup>
                            <InputGroupText addonType="prepend"> <BsTelephone /> </InputGroupText>
                            <InputMask mask="+7(999)-999-9999" value={loginDetails.phoneNumber}
                                onChange={handleChange} className={'telMask'} placeholder={'+7(___)-___-____'}>
                                {(inputProps) => (
                                    <input
                                        {...inputProps}
                                        type="tel"
                                        className="ant-input"
                                    />
                                )}
                            </InputMask>
                            {/* <Input placeholder={'+7(___)___ __ __'} type={"text"} name={"phone"} value={loginDetails.phoneNumber} onChange={handleChange}/> */}
                        </InputGroup>
                    </div>
                    <div className="form-group mt-4">
                        <p className="label">Пароль</p>
                        <InputGroup>
                            <InputGroupText addonType="prepend"> <BsLock /> </InputGroupText>
                            <Input placeholder={'Введите пароль'} type={"password"} name={"password"} value={loginDetails.password} onChange={handleChange} />
                        </InputGroup>
                    </div>
                    <div className="form-group mt-4">
                        <a className="forgot" href="#">Забыли пароль?</a>
                    </div>
                    <button type="submit" className="btn btn-primary btn-block submit" onClick={handleSubmit}>Войти</button>

                </form>
            </div>
        </>

    );

}