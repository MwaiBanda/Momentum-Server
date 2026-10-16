export type Transaction = { 
    id: string,
    amount: number,
    date: string,
    createdOn: string,
    description: string,
    user: User,
}

export type User = {
    id: string,
    email: string,
    fullname: string,
    phone: string,
}
