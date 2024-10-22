<template>
  <el-aside class="menu">
    <el-row :span="2" class="config"> 
      <MenuConfig/>
    </el-row>

    <el-row :span="3" class="bevtree">
      <MenuBevTree/>
    </el-row>

    <el-row :span="19" class="msgnode">
      <MenuRequest/>
    </el-row>
  </el-aside>
</template>

<script lang="ts" setup name="MainMenuView">
  import MenuConfig from "@/components/MenuConfig.vue"
  import MenuBevTree from '@/components/MenuBevTree.vue';
  import MenuRequest from '@/components/MenuRequest.vue';
</script>

<style scoped>
  .menu {
    background-color: #fafafa;
    box-shadow: var(--el-box-shadow-lighter);
    border: 1px solid var(--el-border-color);
    border-radius: var(--border-radius);
    padding-left: 5px;
  }
</style>