export interface ShowMessage {
    id: number;
    protoId: number;
    protoType: string;
    msg: any;
    type: string;
    result: string;
}

export interface FilterMessage {
    protoId: number;
    protoType: string;
}