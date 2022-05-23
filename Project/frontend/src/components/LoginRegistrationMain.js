import React, {useState} from "react";
import '../index.css'
import Login from "./LoginRegistration/Login";
import Registration from "./LoginRegistration/Registration";
export default function LoginRegistrationMain(){
    const [activeTab, setActiveTab] = useState("tab1");
    const handleTab1 = () => {
        // update the state to tab1
        setActiveTab("tab1");
    };
    const handleTab2 = () => {
        // update the state to tab2
        setActiveTab("tab2");
    };
    return (
        <>
            <div className="Tabs">
                <ul className="nav nav-pills mb-3 tabs" id="pills-tab" role="tablist">
                    <li
                        className={"tab-item text-center"}
                        onClick={handleTab1}
                    >
                        <a className={activeTab === "tab1" ? "nav-link active btr" : 'nav-link btl'} id="pills-home-tab"
                           role="tab" aria-controls="pills-home" aria-selected={activeTab==='tab1'}>Регистрация</a>
                    </li>
                    <li
                        className={"tab-item text-center"}
                        onClick={handleTab2}
                    >
                        <a className={activeTab === "tab2" ? "nav-link active btr btlogin" : 'nav-link btl btlogin'} id="pills-home-tab"
                           role="tab" aria-controls="pills-home" aria-selected={activeTab==='tab2'} style={{padding: '10px 40px'}}>Вход</a>
                    </li>
                </ul>

                <div className="outlet">
                    {activeTab === "tab1" ?
                        <Registration/>
                        :
                        <Login/>
                    }
                </div>
            </div>
        </>

    );

}