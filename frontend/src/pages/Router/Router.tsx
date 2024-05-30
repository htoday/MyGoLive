import React from "react";
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {LoginPage} from "../LoginPage/LoginPage.tsx";
import {Home} from "../Home/Home.tsx";

export class WebsiteRouter extends React.Component<any, any>{

    render() {
        return(
            <BrowserRouter>
                <Routes>
                    <Route path={"*"} element={<LoginPage/>}/>
                    <Route path={"/home"} element={<Home/>}></Route>
                </Routes>
            </BrowserRouter>
        )
    }
}