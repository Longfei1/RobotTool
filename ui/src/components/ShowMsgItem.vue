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
  import DialogFactory from '@/utils/DialogFactory';
  import type { ShowMessage, FilterMessage} from '@/type/msg';

  let props = defineProps(["data", "onFilter"]);

  let data = ref<ShowMessage>(props.data);

  function onClickFilter() {
    let msg: FilterMessage = {
      protoId: data.value.protoId,
      protoType: data.value.protoType,
    } 
    props.onFilter?.(msg)
  }

  function onClickView() {
    DialogFactory.OpenJsonViewDialog(data.value.msg);
  }
</script>

<style scoped>

</style>