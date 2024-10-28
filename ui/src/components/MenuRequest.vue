<template>
  <el-container>
      <el-header height="25">快捷请求</el-header>
      <el-main>
        <BtnRequest v-for="btn in cfgData" :data="btn"></BtnRequest>
      </el-main>
  </el-container>
</template>
  
<script lang="ts" setup name="MenuRequest">
  import { ref, onBeforeMount } from 'vue';
  import BtnRequest from './BtnRequest.vue';
  import type { RequestInfo, RequestData, RequestProto } from '@/type/request';
  import CommonFunc from '@/utils/CommonFunc';
  import JsonEditor from '@/utils/JsonEditor';

  let reqList = ref<RequestInfo[]>([])
  let reqProto: RequestProto[] = []

  onBeforeMount(async () => {
    getProtoRequest()

    loadRequestList()
  })

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
    let list = []
    if (requestStr) {
      list = JSON.parse(requestStr)
    }

    reqList.value = []
    for (let msg of list) {
      let info = buildReqInfo(msg)
      if (info) {
        reqList.value.push(info)
      }
    }
  }

  function buildReqInfo(msg: RequestData): (RequestInfo | null) {
    for (let proto of reqProto) {
      if (proto.name == msg.name) {
        //合并一次新模版的值
        let defaultVal = CommonFunc.CloneObj(proto.defaultValue)
        msg.data = Object.assign(defaultVal, msg.data) 

        let info: RequestInfo = {
          data: msg,
          proto: proto,
        }

        return info
      }
    }
    return null
  }
</script>
  
<style scoped>
  .el-container {
    margin-top: 10px;
  }

  .el-container .el-header {
    padding: 0;
    font-weight: bold;
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