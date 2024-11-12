<template>
  <el-select
  v-model="selectServerStr"
  size="default"
  style="width: 135px"
  >
      <el-option
      v-for="item in serverData"
      :key="item.name"
      :value="item.name"
      />
  </el-select>

  <el-button-group>
    <el-button type="primary" @click="onClickCreate">
      <el-icon style="scale: 1.5"> <DocumentAdd/> </el-icon>
    </el-button>
    <el-button type="primary" @click="onClickEdit">
      <el-icon style="scale: 1.5"> <Edit/> </el-icon>
    </el-button>
    <el-button type="primary" @click="onClickDelete">
      <el-icon style="scale: 1.5"> <Delete/> </el-icon>
    </el-button>
  </el-button-group>
</template>
  
<script lang="ts" setup name="MenuConfig">
  import DialogFactory from '@/utils/DialogFactory';
  import { ref, onBeforeMount, watch } from 'vue';
  import type { ServerConfig } from '@/type/server';
  import CommonFunc from '@/utils/CommonFunc';

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
          //selectServerStr.value = cfg.name
          return
        }
      }

      //默认选中第一个
      selectServerStr.value = serverData.value[0].name
  });

  watch([selectServerStr], (oldVal, newVal) => {
    console.log("change sel config, old:%s new:%s", oldVal, selectServerStr.value)
    localStorage.setItem("selectServer", selectServerStr.value)

    setGoServerConfig()
  })

  async function setGoServerConfig() {
    let ret = await CommonFunc.tryPromiseFunc(async () => {
      if (!window.initServerConfig) {
        console.log("initServerConfig is null, wait excute again")
        return false
      }

      let selectServerData = getSelectServerData()
      if (selectServerData) {
        try {
          await window.initServerConfig(selectServerData)
        } catch (e) {
          alert(e);
        }
      }
      return true
    }, 30, 300)

    if (!ret) {
      alert("setGoServerConfig failed, exceed maximum retry times");
    }
  }

  function onClickEdit() {
    let selectServerData = getSelectServerData()
    if (selectServerData) {
      let editName = selectServerData.name
      DialogFactory.OpenJsonEditDialog(selectServerData, (data: any) => {
        let oldCfg = serverData.value
        if (changeServerJson(editName, data)) {
          localStorage.setItem("serverData", JSON.stringify(serverData.value))
          console.log("edit serverjson, target:", editName , "old:", oldCfg, "new:", serverData.value)
          setGoServerConfig()
        }
      });
    } else {
      alert("no valid server config")
    }
  }

  function onClickCreate() {
    DialogFactory.OpenJsonEditDialog(serverTemplates[0], (data: any) => {
      if (changeServerJson("", data)) {
        localStorage.setItem("serverData", JSON.stringify(serverData.value))
        console.log("create serverjson ", serverData.value)
        setGoServerConfig()
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
</script>

<style scoped>
  .el-button-group {
    margin-left: 3px;
  }
</style>