
export type Message = {
    id: string;
    preacher: string;
    series: string;
    thumbnail: string;
    title: string;
    date: string;
    published: boolean;
    passages: Passage[];
}

export type MessageRequest = {
    preacher: string;
    series: string;
    thumbnail: string;
    title: string;
    date: string;
    passages: Passage[];
}

export type Passage = {
    id?: string;
    header?: string ;
    verse?: string;
    message?: string;
    type : string;
    order: number;
}