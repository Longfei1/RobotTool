export default class CommonFunc {
    static CloneObj(obj: any): any {
        return JSON.parse(JSON.stringify(obj))
    }
}