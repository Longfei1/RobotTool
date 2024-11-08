<template>
  <el-dialog class="dialog" 
    v-model="show" 
    width="750"
    height="600"
    @close="close"
    :close-on-click-modal="false"
  >
    <template #header>
      <el-select
      v-model="selectProto"
      size="default"
      filterable
      placeholder="请选择一个消息"
      style="width: 400px"
      >
        <el-option
        v-for="item in protoList"
        :key="item.id"
        :label="item.name"
        :value="item.id"
        />
      </el-select>
    </template>
    <json-editor-vue class="editor" v-model="jsonData" 
      :modeList="['code', 'tree', 'form']"
      :currentMode="'tree'"
      :options="options"
      ref="editorRef">
    </json-editor-vue>
    <template #footer>
      <div class="dialog-footer">
        <el-input v-model="reqName" style="width: 400px; margin-right: 258px" placeholder="快捷消息名称" />
        <el-button type="primary" @click="save">保存</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup name="CreateRequestDialog">
  import JsonEditorVue from 'json-editor-vue3'
  import {ref, watch} from "vue"
  import type { RequestProto, RequestData } from '@/type/request';
  import CommonFunc from '@/utils/CommonFunc';

  let props = defineProps(["data", "onSave", "onClose", "options"])

  let show = ref(true)
  let options = ref(props.options)
  
  let protoList = ref<RequestProto[]>(props.data)
  let selectProto = ref(null)
  let jsonData = ref({})
  let reqName = ref("")

  let editorRef = ref()

  watch([selectProto], (oldVal, newVal) => {
    console.log("change sel proto, old:%s new:%s", oldVal, selectProto.value)
   
    for (let v of protoList.value) {
      if (v.id == selectProto.value) {
        jsonData.value = CommonFunc.CloneObj(v.defaultValue)
        //options.value.schema = v.schema
        //console.log(editorRef.value)
        editorRef.value.editor.setSchema(v.schema)
        reqName.value = v.name
        break
      }
    }
  })

  function close() {
    show.value = false
    let ret: RequestData = {
      uid: "invalid",
      id: selectProto.value,
      name: "",
      data: jsonData.value,
    }
    props.onClose?.(ret)
  }

  function save() {
    show.value = false
    let ret: RequestData = {
      uid: CommonFunc.GenericUid(),
      id: selectProto.value,
      name: reqName.value,
      data: jsonData.value,
    }
    props.onSave?.(ret)
  }
</script>

<style scoped>
  .editor {
    width: 700px;
    height: 450px;
  }
</style>