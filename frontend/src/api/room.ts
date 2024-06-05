export class Room{
    constructor(
        public roomId:number,
        public roomName:string,
        public roomOwner:string,
        public viewerNum:number,
        public displayImage:string="",
    ) {
    }
}
export class CreateRoomRequest{
    constructor(
        public username:string,
        public roomName:string,
    ) {
    }
}
export class CreateRoomResponse{
    constructor(
        public roomId:number,
        public status:number,
    ) {
    }
}
export class JoinRoomRequest{
    constructor(
        public username:string,
        public roomId:number,
    ) {
    }
}
export class JoinRoomResponse{
    constructor(
        public status:number,
    ) {
    }
}
export class CloseRoomRequest{
    constructor(
        public username:string,
        public roomId:number,
    ) {
    }
}
export class CloseRoomResponse{
    constructor(
        public status:number,
    ) {
    }
}
export class GetRoomListRequest{
    constructor(
        public page:string,
    ) {
    }
}
export class GetRoomListResponse{
    constructor(
        public roomList:Room[],
        public status:number,
    ) {
    }
}
export class GetRoomPushAddressRequest{
    constructor(
        public roomId:number,
        public username:string,
    ) {
    }
}
export class GetRoomPushAddressResponse{
    constructor(
        public pushAddress:string,
        public status:number,
        public channelKey:string,
    ) {
    }
}