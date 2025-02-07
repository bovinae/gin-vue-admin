<template>
    <div>
      <el-table :data="configList" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80"></el-table-column>
        <el-table-column prop="camera" label="摄像头IP端口" width="180"></el-table-column>
        <el-table-column prop="encryptBox" label="加密盒子IP" width="150"></el-table-column>
        <el-table-column prop="decryptBox" label="解密盒子IP" width="150"></el-table-column>
        <el-table-column prop="proto" label="协议" width="120"></el-table-column>
        <el-table-column prop="port" label="端口" width="100"></el-table-column>
        <el-table-column prop="modifyTime" label="修改时间" width="120"></el-table-column>
        <el-table-column prop="createTime" label="创建时间" width="120"></el-table-column>
      </el-table>
    </div>
  </template>
  
  <script>
  import { getDeviceConfigList } from "@/api/deviceConfig";
  
  // const page = ref(1)
  // const pageSize = ref(100)

  export default {
    data() {
      return {
        configList: [],
      };
    },
    created() {
      this.fetchConfigList();
    },
    methods: {
      async fetchConfigList() {
        try {
          const res = await getDeviceConfigList({
            // page: page.value,
            // pageSize: pageSize.value
          });
          this.configList = res.data.list;
        } catch (error) {
          console.error("获取配置列表失败:", error);
        }
      },
    },
  };
  </script>