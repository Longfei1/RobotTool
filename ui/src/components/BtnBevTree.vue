<template>
  <el-button type="primary" size="small" @click="onClickExecute">{{data.title}}</el-button>
</template>

<script lang="ts" setup name="BtnBevTree">
  import { ref } from 'vue';

  let props = defineProps(["data"]);

  let data = ref(props.data);

  async function onClickExecute() {
      console.log("click bevtree", data.value);
      
      try {
        let ret = await window.execBevTree(data.value.id)
        console.log(ret);
      } catch (e) {
        alert(e);
      }
  }
</script>

<style scoped>

</style>