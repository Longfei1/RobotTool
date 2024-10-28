<template>
  <el-container class="output">
    <el-header class="tip" height="35">
      <el-tag v-for="tag in tipTag.values()" :type="tag.type">{{tag.name}}:{{tag.msg}}</el-tag>
      <el-text class="tipmsg">{{tipMsg}}</el-text>
    </el-header>

    <el-main class="msglist"> 
      <el-scrollbar>
        <ShowMsgItem v-for="msg in showMsg" :data="msg"></ShowMsgItem>
      </el-scrollbar>
    </el-main>
  </el-container>
</template>

<script lang="ts" setup name="OutPutView">
  import {ref, onBeforeMount} from "vue";
  import ShowMsgItem from '@/components/ShowMsgItem.vue';

  let tipMsg = ref("tip")
  let tipTag = ref(new Map())
  let showMsg = ref<any[]>([])

  onBeforeMount(async ()=> {
    window['addShowMsg'] = addShowMsg
    window['addTipTag'] = addTipTag
    window['addTipMsg'] = addTipMsg
    window['showPopMsg'] = showPopMsg
  })

  function addShowMsg(msg: string) {
    let obj = JSON.parse(msg)
    console.log("addShowMsg", obj)
    showMsg.value.push(obj)
  }

  function addTipTag(tag: string) { 
    let obj = JSON.parse(tag)
    console.log("addTipTag", obj)
    tipTag.value.set(obj.name, obj)
  }

  function addTipMsg(msg: string) { 
    console.log("addTipMsg", msg)
    tipMsg.value = msg
  }

  function showPopMsg(msg: string) { 
    alert(msg)
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
    margin-left: 5px;
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
</style>