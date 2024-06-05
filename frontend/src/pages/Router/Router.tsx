import React from "react";
import {BrowserRouter, Link, Route, Routes} from "react-router-dom";
import {LoginPage} from "../LoginPage/LoginPage.tsx";
import {Home} from "../Home/Home.tsx";

export class WebsiteRouter extends React.Component<any, any>{
    static homeRef:any = React.createRef()
    static loginRef:any = React.createRef()
    render() {
        return(
            <BrowserRouter>
                <Routes>
                    <Route path={"*"} element={<LoginPage/>}/>
                    <Route path={"/home"} element={<Home/>}></Route>
                </Routes>
                <Link to={"/home"} ref={WebsiteRouter.homeRef} style={{display:"none"}}></Link>
                <Link to={""} ref={WebsiteRouter.loginRef} style={{display:"none"}}></Link>
            </BrowserRouter>
        )
    }
}