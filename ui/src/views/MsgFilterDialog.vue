<template>
<el-dialog class="dialog" 
  v-model="show" 
  width="500"
  height="350"
  @close="close"
  :close-on-click-modal="false"
>
  <el-table :data="msgList" height="300" style="width: 100%">
    <el-table-column prop="protoId" label="消息ID" width="140" />
    <el-table-column prop="protoType" label="消息名" width="250" />
    <el-table-column fixed="right" label="操作" min-width="50" width="60">
      <template #default="scope">
        <el-button
          link
          type="primary"
          size="small"
          @click.prevent="deleteItem(scope.$index)"
        >
          删除
        </el-button>
      </template>
    </el-table-column>
  </el-table>
</el-dialog>
</template>

<script lang="ts" setup name="CreateRequestDialog">
  import type { FilterMessage } from '@/type/msg';
  import JsonEditorVue from 'json-editor-vue3'
  import {ref, watch} from "vue"

  let props = defineProps(["data", "onSave", "onClose"])

  let show = ref(true)
  let msgList = props.data

  function close() {
    props.onClose?.()
  }

  function deleteItem(index: number) {
    msgList.value.splice(index, 1)
  }
</script>

<style scoped>

</style>