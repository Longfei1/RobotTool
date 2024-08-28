<template>
  <el-aside class="menu">
    <el-row :span="2" class="config"> 
      config
      <button @click="onClickConfig">config</button>
      <button @click="onClickConfig2">config2</button>
    </el-row>

    <el-row :span="3" class="bevtree">
      <h1>bevtree</h1>
    </el-row>

    <el-row :span="19" class="msgnode">
      <h1>msgnode</h1>
    </el-row>
  </el-aside>
</template>

<script lang="ts" setup name="MainMenuView">
  import JsonEditor from '@/utils/JsonEditor';

  function onClickConfig() {
    JsonEditor.OpenJsonEditDialog({name: "666"}, () => {
      console.log("onsave callback")
    }, () => {
      console.log("onclose callback")
    });
  }

  function onClickConfig2() {
    JsonEditor.OpenJsonViewDialog({name: "666"}, () => {
      console.log("onclose callback")
    });
  }
</script>"

<style scoped>
  .menu {
    background-color: #fafafa;
    box-shadow: var(--el-box-shadow-lighter);
    border: 1px solid var(--el-border-color);
    border-radius: var(--border-radius);
  }
</style>