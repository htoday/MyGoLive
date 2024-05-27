import styles from "./LoginPage.module.less"
import {useEffect, useState} from "react";
import {RememberPasswordButton} from "./components/RememberPasswordButton/RememberPasswordButton.tsx";
export function LoginPage() {

    const [mode,setMode]=useState(0)
    return (
        <>
            <div className={styles.background}>
                <div className={styles.content}>
                    <div className={styles.row}>
                        <div className={styles.image}/>
                        {mode==0&& <Login setMode={()=>{setMode(1)}}/>}
                        {mode==1&&<Register setMode={()=>{setMode(0)}}/>}
                    </div>
                </div>
            </div>

        </>

    )
}
function Login(props:{
    setMode:()=>void
}){
    useEffect(() => {
        document.title = "登录界面"
        return ()=>{
            document.title=""
        }
    }, []);
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
            <input
                type={"text"}
                value={username}
                placeholder={"用户名"}
                onChange={(e) => {
                    setUsername(() => e.target.value)
                }}/>
            <input
                type={"password"}
                value={password}
                placeholder={"密码"}
                onChange={(e)=>{
                    setPassword(() => e.target.value)
                }}
            />
            <RememberPasswordButton onClick={()=>{setRemember(!remember)}} state={remember}></RememberPasswordButton>
            <a onClick={props.setMode}>还没有账号?点我</a>
            <button>
                登 录
            </button>
        </div>
    )
}
export function Register(props:{
    setMode:()=>void
}){
    useEffect(() => {
        document.title = "注册界面"
        return ()=>{
            document.title=""
        }
    }, []);
    const [mobile,setMobile]=useState("")
    const [username,setUsername]=useState("")
    const [password,setPassword]=useState("")
    const [passwordConfirm,setPasswordConfirm]=useState("")
    return (
        <div className={styles.main}>
            <h1 className={styles.title}>注册</h1>
            <input
                type={"text"}
                placeholder={"手机号"}
                value={mobile}
                onChange={(e)=>{
                    setMobile(()=>e.target.value)
                }}
            />
            <input
                type={"text"}
                placeholder={"用户名"}
                value={username}
                onChange={(e)=>{
                    setUsername(()=>e.target.value)
                }}
            />
            <input
                type={"password"}
                placeholder={"密码"}
                value={password}
                onChange={(e)=>{
                    setPassword(() => e.target.value)
                }}
            />
            <input
                type={"password"}
                placeholder={"确认密码"}
                value={passwordConfirm}
                onChange={(e)=>{
                    setPasswordConfirm(()=>e.target.value)
                }}
            />

            <a onClick={props.setMode}>返回登录界面</a>
            <button>注 册</button>
        </div>
    )
}