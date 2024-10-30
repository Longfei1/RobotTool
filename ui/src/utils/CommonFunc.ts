import { v7 as uuidv7 } from "uuid"

export default class CommonFunc {
    static CloneObj(obj: any): any {
        return JSON.parse(JSON.stringify(obj))
    }

    static GenericUid(): string {
        return uuidv7()
    }
}