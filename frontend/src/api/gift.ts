export enum GiftType{
    FLOWER="flower",
    HEART="heart",
    CAKE="cake",
    PLANE="plane",
    ROCKET="rocket"
}
export class Gift{
    constructor(
        public source:string,
        public target:string,
        public giftType:GiftType,
        public message:string,
    ){
    }
}
export function getGiftByName(name:string):GiftType | undefined{
    switch(name){
        case "flower":
            return GiftType.FLOWER
        case "heart":
            return GiftType.HEART
        case "cake":
            return GiftType.CAKE
        case "plane":
            return GiftType.PLANE
        case "rocket":
            return GiftType.ROCKET
        default:
            return undefined
    }
}
export class SendGiftRequest{
    constructor(
        public roomId:number,
        public name:string,
        public giftType:string,
    ) {
    }
}
export class SendGiftResponse{
    constructor(
        public status:number,
    ) {
    }
}