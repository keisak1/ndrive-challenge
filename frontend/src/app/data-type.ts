
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

export interface product{
    id:string,
    name:string,
    price:number,
    stock:number,
    categoryId:Array<string>,
    rating:number,
    image:string

}

export interface category{
    id:string,
    name:string,
    created_at: Date,
    updated_at: Date,
}
