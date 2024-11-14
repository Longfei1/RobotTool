<template>
  <el-dropdown ref="dropdownIns" split-button type="primary" size="small" 
  :hide-on-click="false" 
  @click="onClickExecute" 
  @command="onCommand"
  @visible-change="onVisibleChange">
    {{data.data.name}}
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item command="edit" icon="Edit">编辑</el-dropdown-item>

        <el-dropdown-item command="moveUp" divided icon="Back">
          前移
          <div @click.stop> <!-- 屏蔽点击事件透传到下拉框中 -->
            <el-input-number size="small" v-model="moveUpNum" :min="1" controls-position="right" style="margin-left: 5px;">
              <template #decrease-icon>
                <el-icon>
                  <Minus />
                </el-icon>
              </template>
              <template #increase-icon>
                <el-icon>
                  <Plus />
                </el-icon>
              </template>
            </el-input-number>
          </div>
        </el-dropdown-item>

        <el-dropdown-item command="moveDown" icon="Right">
          后移
          <div @click.stop>  <!-- 屏蔽点击事件透传到下拉框中 -->
            <el-input-number size="small" v-model="moveDownNum" :min="1" controls-position="right" style="margin-left: 5px;">
              <template #decrease-icon>
                <el-icon>
                  <Minus />
                </el-icon>
              </template>
              <template #increase-icon>
                <el-icon>
                  <Plus />
                </el-icon>
              </template>
            </el-input-number>
          </div>
        </el-dropdown-item>

        <el-dropdown-item command="delete" divided icon="Delete">删除</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>
  
<script lang="ts" setup name="BtnBevRequest">
  import { ref } from 'vue';
  import type { RequestInfo } from '@/type/request';
  import DialogFactory from '@/utils/DialogFactory';
import type { DropdownInstance } from 'element-plus';

  let props = defineProps(["data"]);

  const emit = defineEmits(["edit", "move-up", "move-down", "delete"]);

  let data: RequestInfo = props.data;

  let moveUpNum = ref(1)
  let moveDownNum = ref(1)
  let hideOnClick = ref(true)

  const dropdownIns = ref<DropdownInstance>()

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

  function onCommand(command: string) {
    switch (command) {
      case "edit":
        DialogFactory.OpenJsonEditDialog(data.data.data, (jsonData: any) => {
          data.data.data = jsonData;
          emit("edit", data.data)
        }, undefined, {schema: data.proto.schema}); 
        break;
      case "moveUp":
        emit("move-up", data.data.uid, moveUpNum.value);
        break;
      case "moveDown":
        emit("move-down", data.data.uid, moveDownNum.value)
        break;
      case "delete":
        emit("delete", data.data.uid)
        break;
      default:
        break;
    }
    dropdownIns.value?.handleClose();
  }

  function onVisibleChange(visible: boolean) {
    if (visible) {
      moveUpNum.value = 1
      moveDownNum.value = 1
    }
  }
</script>

<style scoped>

</style>