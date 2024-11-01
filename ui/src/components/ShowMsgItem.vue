<template>
  <el-button-group>
    <el-button type="primary" size="small" @click="onClickFilter">
      <el-icon style="scale: 1.5"> <Hide/> </el-icon>
    </el-button>
    <el-button :type="data.type" size="small" @click="onClickView">{{data.protoType}}</el-button>
    <el-button :type="data.type" size="small" @click="onClickView">{{data.msg}}</el-button>
  </el-button-group>
</template>

<script lang="ts" setup name="ShowMsgItem">
  import { ref } from 'vue';
  import JsonEditor from '@/utils/JsonEditor';

  let props = defineProps(["data"]);

  let data = ref(props.data);

  function onClickFilter() {
    console.log(data.value);
  }

  function onClickView() {
    JsonEditor.OpenJsonViewDialog(data.value.msg);
  }
</script>

<style scoped>

</style>