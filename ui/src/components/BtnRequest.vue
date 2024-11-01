<template>
    <el-button-group>
      <el-button type="primary" size="small" @click="onClickExecute">{{data.data.name}}</el-button>
      <el-button type="primary" size="small" @click="onEdit">
        <el-icon style="scale: 1.5"> <Edit/> </el-icon>
      </el-button>
      <el-button type="primary" size="small" @click="onDelete">
        <el-icon style="scale: 1.5"> <Delete/> </el-icon>
      </el-button>
    </el-button-group>
  </template>
  
  <script lang="ts" setup name="BtnBevRequest">
    import { ref } from 'vue';
    import type { RequestInfo } from '@/type/request';
    import JsonEditor from '@/utils/JsonEditor';
  
    let props = defineProps(["data", "onEdit", "onDelete"]);
  
    let data = ref<RequestInfo>(props.data);
  
    async function onClickExecute() {
      console.log("execute request", data.value);
      try {
        let ret = await window.executeRequest({
            id: data.value.data.id, 
            jsonData: JSON.stringify(data.value.data.data),
          }
        );
        //console.log(ret)
      } catch (e) {
        alert(e);
      }
    }

    function onEdit() {
      JsonEditor.OpenJsonEditDialog(data.value.data.data, (jsonData: any) => {
        data.value.data.data = jsonData;
        props.onEdit?.(data.value.data) 
      }, undefined, {schema: data.value.proto.schema}); 
    }

    function onDelete() {
      props.onDelete?.(data.value.data.uid) 
    }
  </script>
  
  <style scoped>
  
  </style>