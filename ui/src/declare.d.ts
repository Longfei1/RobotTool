declare module "json-editor-vue3";
declare interface Window {
    addShowMsg: (s: string) => void;
    addTipTag: (s: string) => void;
    addTipMsg: (s: string) => void;
    showPopMsg: (s: string) => void;
    initServerConfig: (arg: any) => any;
    execBevTree: (id: string) => any;
    getProtoRequest: () => any;
    executeRequest: (arg: any) => any;
}