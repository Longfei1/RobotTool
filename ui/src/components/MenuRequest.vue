<template>
  <el-container>
      <el-header height="25">
        快捷请求
        <el-button type="success" size="small" @click="onClickAdd">
          <el-icon style="scale: 2"> <DocumentAdd/> </el-icon>
        </el-button>
      </el-header>
      <el-main>
        <BtnRequest v-for="v in reqList" :data="v" :onEdit="onReqBtnEdit" :onDelete="onReqBtnDelete"></BtnRequest>
      </el-main>
  </el-container>
</template>
  
<script lang="ts" setup name="MenuRequest">
  import { ref, onBeforeMount, watch } from 'vue';
  import BtnRequest from './BtnRequest.vue';
  import type { RequestInfo, RequestData, RequestProto } from '@/type/request';
  import CommonFunc from '@/utils/CommonFunc';
  import JsonEditor from '@/utils/JsonEditor';
  import { } from '@element-plus/icons-vue'

  let reqList = ref<RequestInfo[]>([])
  let reqProto: RequestProto[] = []
  let reqData = ref<RequestData[]>([])

  onBeforeMount(async () => {
    await getProtoRequest()

    loadRequestList()
  })

  watch(reqData, (oldVal, newVal) => {
    localStorage.setItem("requestList", JSON.stringify(reqData.value))
  }, {deep: true})

  async function getProtoRequest() {
    try {
      let ret = await window.getProtoRequest()
      reqProto = ret
      console.log(reqProto)
    } catch (e) {
      alert(e);
    }
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
    console.log(reqList.value);
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
    JsonEditor.OpenCreateRequestDialog(reqProto, (ret: any) => {
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

    refreshReqList(reqData.value)
  }

  function onReqBtnEdit(data: RequestData) {
    saveRequest(data)
  }

  function onReqBtnDelete(uid: string) {
    for (let i = 0; i < reqData.value.length; i++) {
      if (reqData.value[i].uid == uid) {
        reqData.value.splice(i, 1)
        break
      }
    }

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

  .el-container .el-main .el-button-group {
    margin-top: 3px;
    margin-left: 3px;
  }
</style>