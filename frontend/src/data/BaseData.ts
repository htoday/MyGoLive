export const baseData={
    server:{
        hostname:"localhost",
        post:3000,
        protocol:"http",
        getBaseUrl:function(){
            return this.protocol+"://"+this.hostname+":"+this.post;
        }
    }
}