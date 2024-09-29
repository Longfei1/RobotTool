<template>
  <el-aside class="menu">
    <el-row :span="2" class="config"> 
      <el-select
      v-model="selectServerStr"
      size="default"
      style="width: 100px"
      >
        <el-option
          v-for="item in serverData"
          :key="item.name"
          :value="item.name"
        />
      </el-select>
      <button @click="onClickEdit">编辑</button>
      <button @click="onClickCreate">新建</button>
      <button @click="onClickDelete">删除</button>
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
  import { ref, onBeforeMount } from 'vue';
  import type { ServerConfig } from '@/type/server';
  
  let serverData = ref<ServerConfig[]>([])
  let serverTemplates: ServerConfig[] = []
  let selectServerStr = ref("")

  onBeforeMount(async () => {
      try {
        const response = await fetch('conf/server.json')
        console.log(response)
        if (!response.ok) {
          throw new Error('Network response was not ok')
        }
        serverTemplates = await response.json()
      } catch (error) {
        console.error('There has been a problem with your fetch operation:', error)
      }

      selectServerStr.value = localStorage.getItem("selectServer") || ""
      let serverStr = localStorage.getItem("serverData")
      if (serverStr) {
        serverData.value = JSON.parse(serverStr)
      } else {
        serverData.value = serverTemplates
      }

      //没有配置
      if (serverData.value.length == 0) {
        localStorage.removeItem("selectServer")
        return
      }

      //查找选中配置
      for (const cfg of serverData.value) {
        if (cfg.name == selectServerStr.value) {
          selectServerStr.value = cfg.name
          return
        }
      }

      //默认选中第一个
      selectServerStr.value = serverData.value[0].name
      localStorage.setItem("selectServer", selectServerStr.value)
  });

  function onClickEdit() {
    let selectServerData = getSelectServerData()
    if (selectServerData) {
      let editName = selectServerData.name
      JsonEditor.OpenJsonEditDialog(selectServerData, (data: any) => {
        let oldCfg = serverData.value
        if (changeServerJson(editName, data)) {
          localStorage.setItem("serverData", JSON.stringify(serverData.value))
          localStorage.setItem("selectServer", selectServerStr.value)
          console.log("edit serverjson, target:", editName , "old:", oldCfg, "new:", serverData.value)
        }
      });
    } else {
      alert("no valid server config")
    }
  }

  function onClickCreate() {
    JsonEditor.OpenJsonEditDialog(serverTemplates[0], (data: any) => {
      if (changeServerJson("", data)) {
        localStorage.setItem("serverData", JSON.stringify(serverData.value))
        console.log("create serverjson ", serverData.value)
      }
    });
  }

  function onClickDelete() {
    let selectServerData = getSelectServerData()
    if (selectServerData) {
      for (let i = 0; i < serverData.value.length; i++) {
        if (serverData.value[i].name == selectServerData.name) {
          serverData.value.splice(i, 1)

          localStorage.setItem("serverData", JSON.stringify(serverData.value))
          if (serverData.value.length > 0) {
            selectServerStr.value = serverData.value[0].name
            localStorage.setItem("selectServer", selectServerStr.value)
          }
          return true
        }
      }
    } else {
      alert("no valid server config")
    }
  }

  function getSelectServerData(): ServerConfig | null {
    if (selectServerStr.value) {
        for (const i in serverData.value) {
        if (serverData.value[i].name == selectServerStr.value) {
          return serverData.value[i]
        }
      }
    }
    return null
  }

  function changeServerJson(oldName: string, data: any) : boolean {
    if (!data.name) {
      alert("name不能为空！")
      return false
    }

    let name = oldName    
    if (!name) {
      name = data.name
    }

    //修改指定配置
    for (const i in serverData.value) {
      if (serverData.value[i].name == name) {
        serverData.value[i] = data
        selectServerStr.value = name
        return true
      }
    }

    //插入配置
    serverData.value.push(data)

    return true
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