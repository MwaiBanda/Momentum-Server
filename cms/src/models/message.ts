
export type Message = {
    id: string;
    preacher: string;
    series: string;
    thumbnail: string;
    title: string;
    date: string;
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
    header: string | null;
    verse: string | null;
    message: string | null;
    type : string | null;
}