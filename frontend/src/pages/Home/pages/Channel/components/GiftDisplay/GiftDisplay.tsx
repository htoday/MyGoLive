import styles from "./GiftDisplay.module.less"
import {GiftItem} from "./components/GiftItem/GiftItem.tsx";
import {GiftType} from "../../../../../../api/gift.ts";
import {Room} from "../../../../../../api/room.ts";


export function GiftDisplay(props:{
    room:Room
}){
    const initGiftListComponent=()=>{
        const result=[] as JSX.Element[]
        result.push(
            <GiftItem
                giftType={GiftType.FLOWER}
                giftName={"鲜花"}
                giftPrice={1}
                giftImage={""}
                room={props.room}
            />
        )
        result.push(
            <GiftItem
                giftType={GiftType.HEART}
                giftName={"红心"}
                giftPrice={2}
                giftImage={""}
                room={props.room}
            />
        )
        result.push(
            <GiftItem
                giftType={GiftType.CAKE}
                giftName={"蛋糕"}
                giftPrice={10}
                giftImage={""}
                room={props.room}
            />
        )
        result.push(
            <GiftItem
                giftType={GiftType.PLANE}
                giftName={"飞机"}
                giftPrice={50}
                giftImage={""}
                room={props.room}
            />
        )
        result.push(
            <GiftItem
                giftType={GiftType.ROCKET}
                giftName={"火箭"}
                giftPrice={100}
                giftImage={""}
                room={props.room}
            />
        )
        return result
    }
    const giftList=initGiftListComponent()
    return(
        <div className={styles.content}>
            <h1>礼物列表</h1>
            <div className={styles.gift_list}>
                {giftList}
            </div>
        </div>
    )
}