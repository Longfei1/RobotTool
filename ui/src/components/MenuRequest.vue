<template>
  <el-container>
      <el-header height="25">快捷请求</el-header>
      <el-main>
        <BtnRequest v-for="btn in cfgData.msgs" :data="btn"></BtnRequest>
      </el-main>
  </el-container>
</template>
  
<script lang="ts" setup name="MenuRequest">
  import { ref, onBeforeMount } from 'vue';
  import BtnRequest from './BtnRequest.vue';

  let cfgData = ref({msgs: [{name: "test1"}, {name: "test2"}]})
</script>
  
<style scoped>
  .el-container {
    margin-top: 10px;
  }

  .el-container .el-header {
    padding: 0;
    font-weight: bold;
  }
  .el-container .el-main {
    margin-top: 5px;
    padding: 0;
  }

  .el-container .el-main .el-button-group {
    margin: 1%;
  }
</style>