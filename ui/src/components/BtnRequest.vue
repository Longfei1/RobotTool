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
    import DialogFactory from '@/utils/DialogFactory';
  
    let props = defineProps(["data", "onEdit", "onDelete"]);
  
    let data: RequestInfo = props.data;
  
    async function onClickExecute() {
      console.log("execute request", data);
      try {
        let ret = await window.executeRequest({
            id: data.data.id, 
            jsonData: JSON.stringify(data.data.data),
          }
        );
        //console.log(ret)
      } catch (e) {
        alert(e);
      }
    }

    function onEdit() {
      DialogFactory.OpenJsonEditDialog(data.data.data, (jsonData: any) => {
        data.data.data = jsonData;
        props.onEdit?.(data.data) 
      }, undefined, {schema: data.proto.schema}); 
    }

    function onDelete() {
      console.log("onDelete", data)
      props.onDelete?.(data.data.uid) 
    }
  </script>
  
  <style scoped>
  
  </style>