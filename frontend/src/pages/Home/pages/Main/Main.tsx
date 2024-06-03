import styles from "./Main.module.less"
import {TopMenu} from "./components/TopMenu/TopMenu.tsx";
import {useState} from "react";
import {Search} from "./components/Search/Search.tsx";
export function Main(){
    const [pageIndex,setPageIndex]=useState(0)
    return (
        <div className={styles.content}>
            <Search/>
            <TopMenu menuList={["推荐","游戏","动漫","电影"]}
                     onIndexChange={(index)=>{setPageIndex(index)}}
                     index={pageIndex}/>
        </div>
    )
}