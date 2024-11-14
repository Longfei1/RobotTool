<template>
  <el-container>
    <el-header height="25">
      快捷请求
      <el-button type="primary" size="small" @click="onClickAdd">
        <el-icon style="scale: 2"> <DocumentAdd/> </el-icon>
      </el-button>
    </el-header>
    <el-main>
      <el-space :size="4" wrap>
        <BtnRequest v-for="v in reqList" :data="v" :key="v.data.uid" 
        @edit="onReqBtnEdit" 
        @move-up="onReqBtnMoveUp"
        @move-down="onReqBtnMoveDown"
        @delete="onReqBtnDelete"/>
      </el-space>
    </el-main>
  </el-container>
</template>
  
<script lang="ts" setup name="MenuRequest">
  import { ref, onBeforeMount, watch } from 'vue';
  import BtnRequest from './BtnRequest.vue';
  import type { RequestInfo, RequestData, RequestProto } from '@/type/request';
  import CommonFunc from '@/utils/CommonFunc';
  import DialogFactory from '@/utils/DialogFactory';
  import { } from '@element-plus/icons-vue'

  let reqList = ref<RequestInfo[]>([])
  let reqProto: RequestProto[] = []
  let reqData = ref<RequestData[]>([])

  onBeforeMount(async () => {
    let ret = await CommonFunc.tryPromiseFunc(async () => {
      let ret = await getProtoRequest()
      if (!ret) {
        return false
      }

      loadRequestList()
      return true
    }, 30, 300)

    if (!ret) {
      alert("setGoServerConfig failed, exceed maximum retry times");
    }
  })

  watch(reqData, (oldVal, newVal) => {
    localStorage.setItem("requestList", JSON.stringify(reqData.value))
  }, {deep: true})

  async function getProtoRequest() {
    if (!window.getProtoRequest) {
      console.log("getProtoRequest is null, wait excute again")
      return false
    }
    try {
      let ret = await window.getProtoRequest()
      reqProto = ret
      console.log(reqProto)
    } catch (e) {
      alert(e);
    }
    return true
  }

  function loadRequestList() {
    let requestStr = localStorage.getItem("requestList")
    if (requestStr) {
      reqData.value = JSON.parse(requestStr)
    }

    refreshReqList(reqData.value)
  }

  function refreshReqList(datas: RequestData[]) {
    reqList.value = []
    for (let msg of datas) {
      let info = buildReqInfo(msg)
      if (info) {
        reqList.value.push(info)
      }
    }
    //console.log("refreshReqList", reqList.value);
  }

  function buildReqInfo(msg: RequestData): (RequestInfo | null) {
    for (let proto of reqProto) {
      if (proto.id == msg.id) {
        //合并一次新模版的值
        //let defaultVal = CommonFunc.CloneObj(proto.defaultValue)
        //msg.data = Object.assign(defaultVal, msg.data) 

        let info: RequestInfo = {
          data: msg,
          proto: proto,
        }

        return info
      }
    }
    return null
  }

  function onClickAdd() {
    DialogFactory.OpenCreateRequestDialog(reqProto, (ret: any) => {
      let retData: RequestData = ret
      if (!retData || !retData.id) {
        alert("创建消息失败！")
        return
      }

      saveRequest(retData)
    })
  }

  function saveRequest(req: RequestData) {
    if (!req || !req.id) {
      alert("保存消息失败！")
      return
    }

    let find = false
    for (let info of reqData.value) {
      if (info.uid == req.uid) {
        info.data = req.data
        find = true
        break
      }
    }

    if (!find) {
      reqData.value.push(req)
    }

    //console.log("saveRequest", req)

    refreshReqList(reqData.value)
  }

  function onReqBtnEdit(data: RequestData) {
    saveRequest(data)
  }

  function onReqBtnMoveUp(uid: string, diff: number) {
    console.log("onReqBtnMoveUp", uid, diff)
    moveRequest(uid, -diff)
  }

  function onReqBtnMoveDown(uid: string, diff: number) {
    moveRequest(uid, diff)
  }

  function moveRequest(uid: string, diff: number) {
    if (diff == 0) {
      return
    }

    if (reqData.value.length <= 1) {
      return
    }

    let curIdx = -1
    for (let i = 0; i < reqData.value.length; i++) {
      if (reqData.value[i].uid == uid) {
        curIdx = i
        break
      }
    }

    if (diff > 0) {
      let tmp = reqData.value[curIdx]
      let lastIdx = curIdx + diff >= reqData.value.length - 1 ? reqData.value.length - 1 : curIdx + diff
      for (let i = curIdx + 1; i <= lastIdx; i++) {
        reqData.value[i-1] = reqData.value[i]
      }
      if (curIdx != lastIdx) {
        reqData.value[lastIdx]=tmp
      }
    } else if (diff < 0) {
      let tmp = reqData.value[curIdx]
      let lastIdx = curIdx + diff > 0 ? curIdx + diff : 0
      for (let i = curIdx - 1; i >= lastIdx; i--) {
        reqData.value[i+1] = reqData.value[i]
      }
      if (curIdx != lastIdx) {
        reqData.value[lastIdx]=tmp
      }
    }

    refreshReqList(reqData.value)
  }

  function onReqBtnDelete(uid: string) {
    for (let i = 0; i < reqData.value.length; i++) {
      if (reqData.value[i].uid == uid) {
        reqData.value.splice(i, 1)
        break
      }
    }

    //console.log("onReqBtnDelete reqData", reqData.value)
    refreshReqList(reqData.value)
  }
</script>
  
<style scoped>
  .el-container {
    margin-top: 15px;
  }

  .el-container .el-header {
    padding: 0;
    font-weight: bold;
  }

  .el-container .el-header .el-button {
    scale: 0.7;
  }

  .el-container .el-main {
    margin-top: 5px;
    padding: 0;
  }
</style>