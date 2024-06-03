import styles from "./LoginPage.module.less"
import {useEffect, useState} from "react";
import {RememberPasswordButton} from "./components/RememberPasswordButton/RememberPasswordButton.tsx";
import {baseData} from "../../data/BaseData.ts";
import {isPhone, oneRunningAsync} from "../../utils/Utils.ts";

export function LoginPage() {

    const [mode,setMode]=useState(0)
    return (
        <>
            <div className={styles.background}>
                <div className={styles.content}>
                    <div className={styles.row}>
                        {!isPhone() && <div className={styles.image}/>}
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
    const [remember, setRemember] = useState(localStorage.getItem("remember") === "true")
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const handleLogin = oneRunningAsync(async () => {
        class LoginRequest {
            constructor(
                public username: string,
                public password: string,
            ) {
            }
        }

        class LoginResponse {
            constructor(
                public status: boolean,
                public token: string,
                public expireTime: number,
                public refreshAfter: number,
            ) {
            }
        }

        if (!(username && password)) {
            alert("请输入用户名和密码")
            return
        }
        if (remember) {
            localStorage.setItem("rememberUsername", username)
            localStorage.setItem("rememberPassword", password)
        }
        const url = baseData.server.getBaseUrl() + "/user/login"
        return await fetch(url, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(new LoginRequest(username, password))
        }).then((res) => {

            if (res.status === 200) {
                return res.json()
            } else {
                throw new Error("登录失败")

            }
        }).then((res: LoginResponse) => {

            if (remember) {
                localStorage.setItem("remember", true.toString())
                localStorage.setItem("rememberUsername", username)
                localStorage.setItem("rememberPassword", password)
            }
            localStorage.setItem("token", res.token)
            localStorage.setItem("expireTime", res.expireTime.toString())
            localStorage.setItem("refreshAfter", res.refreshAfter.toString())
        }).catch((e) => {
            alert("登录失败:" + e.message)

        })

    }, () => {
        alert("正在登录中 请勿重复点击!")
    })
    useEffect(() => {

        if (localStorage.getItem("remember")) {
            //console.log(localStorage.getItem("remember"))
            const isRemember = localStorage.getItem("remember") === "true"
            //console.log(isRemember)
            setRemember(remember)
            if (isRemember) {
                if (localStorage.getItem("rememberUsername")) {
                    setUsername(()=>localStorage.getItem("rememberUsername") as string)
                }
                if (localStorage.getItem("rememberPassword")) {
                    setPassword(()=>localStorage.getItem("rememberPassword") as string)
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
            <RememberPasswordButton
                onClick={() => {
                    setRemember(!remember)
                    //console.log(remember.toString())
                    localStorage.setItem("remember", (!remember).toString())
                }}
                state={remember}></RememberPasswordButton>
            <a onClick={props.setMode}>还没有账号?点我</a>
            <button onClick={handleLogin}>
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
    const handleRegister = oneRunningAsync(async () => {
        class RegisterRequest {
            constructor(
                public mobile: string,
                public username: string,
                public password: string,
                public code: number,
            ) {
            }
        }

        class RegisterResponse {
            constructor(
                public status: number
            ) {
            }
        }
        const url =baseData.server.getBaseUrl()+"/user/register"
        return await fetch(url, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(new RegisterRequest(mobile, username, password, 0))
        }).then(res => {
            if (!res.ok) throw new Error(res.status.toString())
            return new RegisterResponse(res.status)
        }).then(() => {
            alert("注册成功!")
            props.setMode()
        }).catch(err=>{
            alert("注册失败:"+err)
        })
    }, () => {
        alert("正在注册中 请勿重复点击!")
    })
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
            <button onClick={handleRegister}>注 册</button>
        </div>
    )
}