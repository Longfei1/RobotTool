<template>
  <el-container class="mainFrame">
    <OutputView/>
    <MainMenuView/>
  </el-container>
</template>

<script setup lang="ts">
  import OutputView from "./views/OutputView.vue";
  import MainMenuView from "./views/MainMenuView.vue";

</script>

<style scoped>
  .mainFrame {
    height: 100vh;
    min-width: 1024px;
    min-height: 768px;
  }
</style>