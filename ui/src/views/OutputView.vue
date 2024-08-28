<template>
  <el-container class="output">
    <el-header class="tip" height="35">
      <el-tag type="primary">Tag 1</el-tag>
      <el-tag type="success">Tag 2</el-tag>
      <el-tag type="info">Tag 3</el-tag>
      <el-tag type="warning">Tag 4</el-tag>
      <el-tag type="danger">Tag 5</el-tag>
      <el-text class="showmsg">feaggiurhuiwwwwwwwwwwwwwwwwwwgerhguienvuieruibheruighiuehruvbnerwu9ghiuerhbuihguibehbrui</el-text>
    </el-header>

    <el-main class="msglist"> 
      msglist
    </el-main>
  </el-container>
</template>

<script lang="ts" setup name="OutPutView">
</script>

<style scoped>
  .output {
    height: 100vh;
  }

  .tip {
    padding-left: 5px;
    box-shadow: var(--el-box-shadow-lighter);
    border: 1px solid var(--el-border-color);
    border-radius: var(--border-radius);
    
    white-space: nowrap;
    line-height: 30px;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .tip .el-tag {
    margin-left: 5px;
  }

  .showmsg {
    margin-left: 5px;
  }

  .msglist {
    padding-left: 5px;
    padding-left: 5px;
    padding-top: 5px;
    padding-bottom: 5px;
    box-shadow: var(--el-box-shadow-lighter);
    border: 1px solid var(--el-border-color);
    border-radius: var(--border-radius);
  }
</style>