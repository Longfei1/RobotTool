<template>
    <el-button-group>
      <el-button type="primary" size="small" @click="onClickExecute">{{data.name}}</el-button>
      <el-button type="primary" size="small" :icon="Edit"/>
    </el-button-group>
  </template>
  
  <script lang="ts" setup name="BtnBevRequest">
    import { ref } from 'vue';
    import { Edit } from '@element-plus/icons-vue'
  
    let props = defineProps(["data"]);
  
    let data = ref(props.data);
  
    function onClickExecute() {
        console.log(data.value);
    }
  </script>
  
  <style scoped>
  
  </style>