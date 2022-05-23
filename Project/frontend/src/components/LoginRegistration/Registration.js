import React, { useState } from "react";
import '../../index.css'
import { Input, InputGroup, InputGroupText, Label } from "reactstrap";
import { BsTelephone } from 'react-icons/bs';
import InputMask from 'react-input-mask';
/*import { Input, Label } from 'semantic-ui-react'*/
export default function Registration() {
    const [phoneNumber, setPhoneNumber] = useState(null)

    function handleChange(event) {
        setPhoneNumber(event)
    }
    return (
        <>
            <div className="auth-inner">
                <form>

                    <h3 className="authorization">Регистрация</h3>
                    <p className="s16">Введите номер телефона для регистрации</p>
                    <div className="form-group">
                        <p className="label">Номер телефона</p>
                        {/*<input type="text" className="form-control" placeholder="First name" />*/}
                        <InputGroup>
                            <InputGroupText addonType="prepend"> <BsTelephone /> </InputGroupText>
                            <InputMask mask="+7(999)-999-9999" value={phoneNumber}
                                onChange={handleChange} className={'telMask'} placeholder={'+7(___)-___-____'}>
                                {(inputProps) => (
                                    <input
                                        {...inputProps}
                                        type="tel"
                                        className="ant-input"
                                    />
                                )}
                            </InputMask>
                            {/* <Input placeholder={'+7(___)___ __ __'} type="text"name="phone"/> */}
                        </InputGroup>
                    </div>
                    <button type="submit" className="btn btn-primary btn-block submit">Войти</button>

                </form>
            </div>
        </>

    );

}