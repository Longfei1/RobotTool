import type { EvOpenJsonEidtor } from "@/type/event";
import JsonEditDialog from "@/views/JsonEditDialog.vue";
import JsonViewDialog from "@/views/JsonViewDialog.vue";
import CreateRequestDialog from "@/views/CreateRequestDialog.vue";
import { h, createApp, render } from "vue";
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import MsgFilterDialog from "@/views/MsgFilterDialog.vue";

class DialogFactory {
    static OpenJsonEditDialog(data: object, onSave?: (d: object) => void, onClose?: (d: object) => void, options: object = {}){
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
            options: options,
        })
        dlg.use(ElementPlus)
        dlg.mount(mountNode)
        document.body.appendChild(mountNode)
    }

    static OpenJsonViewDialog(data: object, onClose?: (d: object) => void, options: object = {},){
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
            options: options,
        })
        dlg.use(ElementPlus)
        dlg.mount(mountNode)
        document.body.appendChild(mountNode)
    }

    static OpenCreateRequestDialog(data: object, onSave?: (d: object) => void, onClose?: (d: object) => void, options: object = {}){
        let closed = false;
        let mountNode = document.createElement("div"); 
        let dlg = createApp(CreateRequestDialog, {
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
            options: options,
        })
        dlg.use(ElementPlus)
        dlg.mount(mountNode)
        document.body.appendChild(mountNode)
    }

    static OpenMsgFilterDialog(data: object, onClose?: () => void){
        let closed = false;
        let mountNode = document.createElement("div"); 
        let dlg = createApp(MsgFilterDialog, {
            data: data,
            onClose: () => {
                if (!closed) {
                    dlg.unmount();
                    closed = true;
                    mountNode.remove()

                    if (onClose) {
                        onClose()
                    }
                }
            },
        })
        dlg.use(ElementPlus)
        dlg.mount(mountNode)
        document.body.appendChild(mountNode)
    }
}

export default DialogFactory;