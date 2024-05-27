import styles from "./LoginPage.module.less"
import {useEffect, useState} from "react";
export function LoginPage() {


    return (
        <>
            <div className={styles.background}>
                <div className={styles.content}>
                    <div className={styles.row}>
                        <div className={styles.image}/>
                        <Login></Login>
                    </div>
                </div>
            </div>

        </>

    )
}
function Login(){
    const [remember, setRemember] = useState(false)
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    useEffect(() => {

        if (localStorage.getItem("remember")) {
            const isRemember = localStorage.getItem("remember") === "true"
            setRemember(remember)
            if (isRemember) {
                if (localStorage.getItem("rememberUsername")) {
                    setUsername(localStorage.getItem("rememberUsername") as string)
                }
                if (localStorage.getItem("rememberPassword")) {
                    setPassword(localStorage.getItem("rememberPassword") as string)
                }
            }
        } else {
            localStorage.setItem("remember", false.toString())
        }
    }, []);
    return (
        <div className={styles.main}>
            <h1 className={styles.title}>登录</h1>
            <input value={username} placeholder={"用户名"}/>
            <input value={password} placeholder={"密码"}/>
        </div>
    )
}