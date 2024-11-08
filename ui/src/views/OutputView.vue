<template>
  <el-container class="output">
    <el-header class="tip" height="35px">
      <el-tag v-for="tag in tipTag.values()" :type="tag.type" :key="tag.name">{{tag.name}}:{{tag.msg}}</el-tag>
    </el-header>

    <el-main class="msglist"> 
      <el-scrollbar>
        <ShowMsgItem v-for="msg in showMsg" :data="msg" :key="msg.id" :onFilter="onMsgBtnFilter"></ShowMsgItem>
      </el-scrollbar>
    </el-main>

    <el-footer class="bottom" height="30px">
      <el-button type="warning" size="small" @click="onClickFilter">
        <el-icon style="scale: 1.5"> <Filter/> </el-icon>
      </el-button>
      <el-button type="warning" size="small" @click="onClickClear">
        <el-icon style="scale: 1.5"> <CircleClose/> </el-icon>
      </el-button>
      <el-text class="tipmsg" type="primary">{{tipMsg}}</el-text>
    </el-footer>
  </el-container>
</template>

<script lang="ts" setup name="OutPutView">
  import {ref, onBeforeMount, watch} from "vue";
  import ShowMsgItem from '@/components/ShowMsgItem.vue';

  import type { ShowMessage, FilterMessage } from '@/type/msg';
import DialogFactory from "@/utils/DialogFactory";

  let tipMsg = ref("tip")
  let tipTag = ref(new Map())
  let showMsg = ref<ShowMessage[]>([])
  let filterMsg = ref<FilterMessage[]>([])

  onBeforeMount(async ()=> {
    window['addShowMsg'] = addShowMsg
    window['addTipTag'] = addTipTag
    window['addTipMsg'] = addTipMsg
    window['showPopMsg'] = showPopMsg

    loadFilterMsg()
  })

  watch(filterMsg, (oldVal, newVal) => {
    localStorage.setItem("msgFilter", JSON.stringify(filterMsg.value))
  }, {deep: true})

  function loadFilterMsg() {
    let requestStr = localStorage.getItem("msgFilter")
    if (requestStr) {
      filterMsg.value = JSON.parse(requestStr)
    }

    refreshShowMsgList()
  }

  function refreshShowMsgList() {
    for (let i = showMsg.value.length - 1; i >= 0; i--) {
      let find = false
      for (let filter of filterMsg.value) {
        if (showMsg.value[i].protoId == filter.protoId) {
          find = true
          break
        }
      }

      if (find) {
        showMsg.value.splice(i, 1) //移除
      }
    }
  }

  function clearShowMsgList() {
    showMsg.value = []
  }

  function pushShowMsg(msg: ShowMessage) {
    for (let e of filterMsg.value) {
      if (e.protoId == msg.protoId) {
        return undefined
      }
    }
    showMsg.value.push(msg)
  }

  function addShowMsg(msg: string) {
    let obj = JSON.parse(msg)
    //console.log("addShowMsg", obj)
    pushShowMsg(obj)
  }

  function addTipTag(tag: string) { 
    let obj = JSON.parse(tag)
    //console.log("addTipTag", obj)
    tipTag.value.set(obj.name, obj)
  }

  function addTipMsg(msg: string) { 
    //console.log("addTipMsg", msg)
    tipMsg.value = msg
  }

  function showPopMsg(msg: string) { 
    alert(msg)
  }

  function onClickFilter() {
   DialogFactory.OpenMsgFilterDialog(filterMsg)
  }

  function onClickClear() {
    clearShowMsgList()
  }

  function onMsgBtnFilter(msg: FilterMessage) {
    filterMsg.value.push(msg)
    refreshShowMsgList()
  }
</script>

<style scoped>
  .tip {
    padding-left: 5px;
    box-shadow: var(--el-box-shadow-lighter);
    border: 1px solid var(--el-border-color);
    border-radius: var(--border-radius);
    
    white-space: nowrap;
    line-height: 30px;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .tip .el-tag {
    margin-left: 5px;
  }

  .tipmsg {
    margin-left: 15px;
  }

  .msglist {
    padding-left: 5px;
    padding-left: 5px;
    padding-top: 5px;
    padding-bottom: 5px;
    box-shadow: var(--el-box-shadow-lighter);
    border: 1px solid var(--el-border-color);
    border-radius: var(--border-radius);
  }

  .msglist .el-scrollbar .el-button-group {
    display: flex;
    margin-left: 10px;
    margin-top: 5px;
  }

  .bottom {
    padding: 0 5px 0 0;
  }

  .bottom .el-button {
    margin-top: 3px;
    margin-left: 5px;
  }
</style>