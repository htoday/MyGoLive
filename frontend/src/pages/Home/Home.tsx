import {BottomMenu} from "./components/BottomMenu/BottomMenu.tsx";
import {useState} from "react";
import styles from "./Home.module.less"
import {Channel} from "./pages/Channel/Channel.tsx";
export function Home() {
    const [index, setIndex] = useState(0)
    return (
        <>
            <div className={styles.background}>
                <div className={styles.page_area}>
                    {index===1&&<Channel/>}
                </div>
            </div>
            <BottomMenu
                onIndexChange={(index) => {
                    setIndex(() => index)
                }}
                menuList={["主 页","频 道","我 的","设 置"]}
                index={index}
            />
        </>

    )
}