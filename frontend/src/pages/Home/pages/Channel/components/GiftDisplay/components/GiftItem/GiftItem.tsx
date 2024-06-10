import styles from "./GiftItem.module.less"
import {GiftType, SendGiftRequest} from "../../../../../../../../api/gift.ts";
import {baseData} from "../../../../../../../../data/BaseData.ts";
import {Room} from "../../../../../../../../api/room.ts";
import {oneRunningAsync} from "../../../../../../../../utils/Utils.ts";
export function GiftItem(props:{
    giftType:GiftType,
    giftName:string,
    giftPrice:number,
    giftImage:string,
    room:Room,

}){
    const handleSendGift=oneRunningAsync(
        async ()=>{
            const url=baseData.talkApiServer.getBaseUrl()+"/gift"
            //确定是否发送礼物
            if(!confirm(`确定要花费${props.giftPrice}元发送${props.giftName}吗?`)){
                return
            }
            await fetch(url,{
                method:"POST",
                headers:{
                    "Content-Type":"application/json"
                },
                body:JSON.stringify(new SendGiftRequest(
                    props.room.roomId,
                    localStorage.getItem("username")!,
                    props.giftType.valueOf(),
                ))
            }).then(res=> {
                if(!res.ok){
                    throw new Error("请求失败")
                }
                return res.json()
            }).then(()=>{
                alert("发送礼物成功!")
            }).catch(()=>{
                alert("发送礼物失败! 请检查你的金额或网络连接!")
            })
        },()=>{
            alert("请勿多次点击!")
        }
    )

    return (
        <div className={styles.content} onClick={handleSendGift}>
            <img src={props.giftImage} alt={"图片无法加载"}/>
            <div className={styles.row}>
                <h1>{props.giftName}</h1>
            </div>
            <div className={styles.row}>
                <p className={styles.price}> [{props.giftPrice}] 金币</p>

            </div>
            <div className={styles.row}>
                <p className={styles.tips}>
                    点击向主播发送礼物
                </p>
            </div>
        </div>
    )

}