import type { EvOpenJsonEidtor } from "@/type/event";
import JsonEditDialog from "@/views/JsonEditDialog.vue";
import JsonViewDialog from "@/views/JsonViewDialog.vue";
import { h, createApp, render } from "vue";
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';

class JsonEditor {
    static OpenJsonEditDialog(data: object, onSave: (d: object) => void, onClose: (d: object) => void){
        let closed = false;
        let mountNode = document.createElement("div"); 
        let dlg = createApp(JsonEditDialog, {
            data: data,
            onSave: (data: object) => {
                if (!closed) {
                    dlg.unmount();
                    closed = true;
                    document.body.removeChild(mountNode)

                    if (onSave) {
                        onSave(data)
                    }
                }
            },
            onClose: (data: object) => {
                if (!closed) {
                    dlg.unmount();
                    closed = true;
                    mountNode.remove()

                    if (onClose) {
                        onClose(data)
                    }
                }
            },
        })
        dlg.use(ElementPlus)
        dlg.mount(mountNode)
        document.body.appendChild(mountNode)
    }

    static OpenJsonViewDialog(data: object, onClose: (d: object) => void){
        let closed = false;
        let mountNode = document.createElement("div"); 
        let dlg = createApp(JsonViewDialog, {
            data: data,
            onClose: (data: object) => {
                if (!closed) {
                    dlg.unmount();
                    closed = true;
                    mountNode.remove()

                    if (onClose) {
                        onClose(data)
                    }
                }
            },
        })
        dlg.use(ElementPlus)
        dlg.mount(mountNode)
        document.body.appendChild(mountNode)
    }
}

export default JsonEditor;