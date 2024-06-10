export enum MessageType{
    TEXT,
    GIFT,
    RAFFLE_START,
    RAFFLE_END,
    SYSTEM,
}
export class Message{
    constructor(
        public msgType:number,
        public content:string,
        public name:string,
    ) {
    }
}
