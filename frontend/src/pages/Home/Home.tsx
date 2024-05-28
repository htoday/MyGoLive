import {BottomMenu} from "./components/BottomMenu/BottomMenu.tsx";
import {useState} from "react";
import styles from "./Home.module.less"
export function Home() {
    const [index, setIndex] = useState(0)
    return (
        <>
            <div className={styles.background}>
                Home
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