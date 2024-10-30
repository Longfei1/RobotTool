
export interface RequestProto {
    id: any;
    name: string;
    schema: object;
    defaultValue: object;
}

export interface RequestData {
    uid: string; // 存储时的唯一id
    id: any; // 对应消息号
    name: string; //名称（别名）
    data: object;
}

export interface RequestInfo {
    // 定义ts中需要引用的变量就行
    data: RequestData;
    proto: RequestProto;
}
