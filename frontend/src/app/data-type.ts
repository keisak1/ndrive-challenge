export interface register{
    name:string,
    email: string,
    password: string,
    passwordConfirm: string
}

export interface login{
    email:string,
    password:string
}