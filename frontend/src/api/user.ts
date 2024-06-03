export class RegisterRequest {
    constructor(
        public mobile: string,
        public username: string,
        public password: string,
        public code: number,
    ) {
    }
}

export class RegisterResponse {
    constructor(
        public status: number
    ) {
    }
}
export class LoginRequest {
    constructor(
        public username: string,
        public password: string,
    ) {
    }
}

export class LoginResponse {
    constructor(
        public status: boolean,
        public token: string,
        public expireTime: number,
        public refreshAfter: number,
    ) {
    }
}