<template>
  <el-dialog class="dialog" 
    v-model="show" 
    width="750"
    height="600"
    @close="close"
    :close-on-click-modal="false"
  >
    <json-editor-vue class="editor" v-model="jsonData" 
      :modeList="['view']"
      :currentMode="'view'">
    </json-editor-vue>
  </el-dialog>
</template>

<script lang="ts" setup name="JsonViewDialog">
  import JsonEditorVue from 'json-editor-vue3'
  import {ref} from "vue"

  let props = defineProps(["data", "onClose"])

  let jsonData = ref(props.data)
  let show = ref(true)

  function close() {
    show.value = false
    props.onClose?.(jsonData.value)
  }
</script>

<style scoped>
  .editor {
      width: 700px;
      height: 480px;
    }
</style>