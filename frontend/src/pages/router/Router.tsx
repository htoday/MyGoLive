import React from "react";
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {LoginPage} from "../LoginPage/LoginPage.tsx";

export class WebsiteRouter extends React.Component<any, any>{

    render() {
        return(
            <BrowserRouter>
                <Routes>
                    <Route path={"*"} element={<LoginPage/>}/>
                </Routes>
            </BrowserRouter>
        )
    }
}