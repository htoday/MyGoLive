import styles from "./GiftLayer.module.less"
import React, {useEffect} from "react";

export enum GiftId{
    COMMON
}
export class Gift{
    constructor(
        public source:string,
        public target:string,
        public id:GiftId,
        public message:string,
    ){
    }
}
// @ts-ignore
export function GiftLayer(props:{
    gifts:Gift[],
    clearGifts:()=>void,
}){
    const layerRef=React.createRef() as React.RefObject<HTMLDivElement>;
    const onGift:(layer:React.RefObject<HTMLDivElement>,gift:Gift)=>void=onGiftFunction
    useEffect(() => {
        console.log(...props.gifts)
        for(let i=0;i<props.gifts.length;i++){
            onGift(layerRef,props.gifts[i])
            //console.log(props.gifts[i])


        }
    }, [props]);

    return(
        <div className={styles.layer} ref={layerRef}>
            这是礼物层 无法被点击 用于描述送礼物的特效
        </div>
    )
}
function onGiftFunction(layer:React.RefObject<HTMLDivElement>,gift:Gift){
    switch (gift.id){
        case GiftId.COMMON:{
            let newDiv=document.createElement("p")
            newDiv.style.position="absolute"
            let yPosition=Math.random()*100
            newDiv.style.top=yPosition+"%"
            newDiv.style.left="0%"
            newDiv.innerHTML=`用户 ${gift.source} 向主播 ${gift.target} 送出了一份普通礼物 并说 ${gift.message}`
            newDiv.className=styles.text
            layer.current?.appendChild(newDiv)
            let startLocation=0
            let moveTime=1000
            let speed=10
            let count=0

            const move=()=>{
                if(count>moveTime){
                    newDiv.remove()
                }else{
                    const nowLocation=startLocation+count*speed
                    newDiv.style.left=nowLocation/100+"%"
                    requestAnimationFrame(move)
                    count++
                }

            }
            requestAnimationFrame(move)
            break;
        }
        default:{

        }
    }

}