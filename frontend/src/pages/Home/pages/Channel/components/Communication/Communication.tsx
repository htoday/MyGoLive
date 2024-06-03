import {useState} from "react";
import styles from "./Communication.module.less"
enum MessageType{
    TEXT,
    GIFT,
    IMAGE,
    SYSTEM
}
export class Message{
    constructor(
        public type:MessageType,
        public content:string,
        public senderUid:number,
    ) {
    }
}
export function Communication(){
    const [message,setMessage]=useState([] as Message[])
    return(
        <div className={styles.communication_window}>

        </div>
    )
}