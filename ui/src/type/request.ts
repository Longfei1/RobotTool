
export interface RequestProto {
    id: any;
    name: string;
    schema: object;
    defaultValue: object;
}

export interface RequestData {
    name: string;
    data: object;
}

export interface RequestInfo {
    // 定义ts中需要引用的变量就行
    data: RequestData;
    proto: RequestProto;
}
