<template>
  <el-container>
      <el-header height="25">行为树</el-header>
      <el-main>
        <BtnBevTree v-for="btn in cfgData.trees" :key="btn.id" :data="btn"></BtnBevTree>
      </el-main>
  </el-container>
</template>
  
<script lang="ts" setup name="MenuBevTree">
  import { ref, onBeforeMount } from 'vue';
  import BtnBevTree from './BtnBevTree.vue';

  let cfgData = ref<any>({});

  onBeforeMount(loadConfig);

  async function loadConfig() {
    try {
      const response = await fetch('conf/behavior.json')
      console.log(response)
      if (!response.ok) {
        throw new Error('Network response was not ok')
      }
      cfgData.value = await response.json()
    } catch (error) {
      console.error('There has been a problem with your fetch operation:', error)
    }
  }
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

  .el-container .el-main .el-button {
    margin-top: 3px;
    margin-left: 3px;
  }
</style>