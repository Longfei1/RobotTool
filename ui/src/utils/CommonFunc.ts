import { v7 as uuidv7 } from "uuid"

export default class CommonFunc {
    static CloneObj(obj: any): any {
        return JSON.parse(JSON.stringify(obj))
    }

    static GenericUid(): string {
        return uuidv7()
    }

    static tryPromiseFunc(f: () => Promise<Boolean>, retryCount: number = 0, intervalMs: number = 1000) {
        let finalRet = new Promise((resolve, reject) => {
            let count = 0

            let wrapF = async () => {return true}
            wrapF = async () => {
                let ret = await f()
                if (!ret) {
                    if (count < retryCount) {
                        count++
                        setTimeout(wrapF, intervalMs)
                        return true
                    } 

                    resolve(false)
                    return true
                }

                resolve(true)
                return true
            }
            wrapF()
        });

        return finalRet
    }
}