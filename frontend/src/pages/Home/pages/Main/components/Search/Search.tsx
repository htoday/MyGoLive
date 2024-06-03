import styles from "./Search.module.less"
import iconSearch from "../../../../../../assets/icon/MaterialSymbolsSearch.svg"
export function Search(){
    return(
        <div className={styles.center}>
            <div className={styles.content}>

                <input placeholder={"搜索想看的内容吧"}/>
                <img src={iconSearch} alt={""}/>
            </div>
        </div>
    )
}