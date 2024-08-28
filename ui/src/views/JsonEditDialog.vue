<template>
  <el-dialog class="dialog" 
    v-model="show" 
    width="750"
    height="600"
    @close="close"
    :close-on-click-modal="false"
  >
    <json-editor-vue class="editor" v-model="jsonData" 
      :modeList="['code', 'tree', 'form']"
      :currentMode="'tree'">
    </json-editor-vue>
    <template #footer>
      <div class="dialog-footer">
        <el-button type="primary" @click="save">保存</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup name="JsonEditDialog">
  import JsonEditorVue from 'json-editor-vue3'
  import {ref} from "vue"

  let props = defineProps(["data", "onSave", "onClose"])

  let jsonData = ref(props.data)
  let show = ref(true)

  function close() {
    show.value = false
    props.onClose?.(jsonData)
  }

  function save() {
    show.value = false
    props.onSave?.(jsonData)
  }
</script>

<style scoped>
    .editor {
      width: 700px;
      height: 450px;
    }
</style>