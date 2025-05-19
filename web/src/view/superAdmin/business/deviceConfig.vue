<template>
    <div>
      <!-- <warning-bar
        title="在资源权限中将此角色的资源权限清空 或者不包含创建者的角色 即可屏蔽此客户资源的显示"
      /> -->
      <div class="gva-table-box">
        <div class="gva-btn-list">
          <el-button type="primary" icon="plus" @click="openDrawer"
            >新增</el-button
          >
        </div>
        <el-table
          ref="multipleTable"
          :data="tableData"
          style="width: 100%"
          tooltip-effect="dark"
          row-key="ID"
        >
          <el-table-column type="selection" width="20" />
          <el-table-column
            align="left"
            label="ID"
            prop="id"
            width="80"
          />
          <el-table-column
            align="left"
            label="摄像头IP端口"
            prop="camera"
            width="180"
          />
          <el-table-column
            align="left"
            label="加密盒子IP"
            prop="encryptBox"
            width="150"
          />
          <el-table-column
            align="left"
            label="启用加密盒子"
            prop="enableEncryptBox"
            width="120"
          />
          <el-table-column
            align="left"
            label="解密盒子IP"
            prop="decryptBox"
            width="150"
          />
          <el-table-column
            align="left"
            label="启用解密盒子"
            prop="enableDecryptBox"
            width="120"
          />
          <el-table-column
            align="left"
            label="tcp端口"
            prop="tcpPort"
            width="120"
          />
          <el-table-column
            align="left"
            label="udp端口"
            prop="udpPort"
            width="100"
          />
          <el-table-column
            align="left"
            label="修改时间"
            prop="modifyTime"
            width="180"
          />
          <el-table-column
            align="left"
            label="创建时间"
            prop="createTime"
            width="180"
          />
          <el-table-column align="left" label="操作" min-width="200">
            <template #default="scope">
              <el-button
                type="primary"
                link
                icon="edit"
                @click="updateDevice(scope.row)"
                >变更</el-button
              >
              <el-button
                type="primary"
                link
                icon="delete"
                @click="deleteDevice(scope.row)"
                >删除</el-button
              >
              <el-button
                type="primary"
                link
                icon="video-camera"
                @click="handlePlayVideo(scope.row)"
                >播放</el-button
              >
            </template>
          </el-table-column>
        </el-table>
        <div class="gva-pagination">
          <el-pagination
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            layout="total, sizes, prev, pager, next, jumper"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
          />
        </div>
      </div>
      <el-drawer
        v-model="drawerFormVisible"
        :before-close="closeDrawer"
        :show-close="false"
      >
        <template #header>
          <div class="flex justify-between items-center">
            <span class="text-lg">设备配置</span>
            <div>
              <el-button @click="closeDrawer">取 消</el-button>
              <el-button type="primary" @click="enterDrawer">确 定</el-button>
            </div>
          </div>
        </template>
        <el-form :inline="true" :model="form" label-width="120px">
          <el-form-item label="摄像头IP端口">
            <el-input v-model="form.camera" autocomplete="off" />
          </el-form-item>
          <el-form-item label="加密盒子IP">
            <el-input v-model="form.encryptBox" autocomplete="off" />
          </el-form-item>
          <el-form-item label="启用加密盒子">
            <el-switch v-model="form.enableEncryptBox" />
          </el-form-item>
          <el-form-item label="解密盒子IP">
            <el-input v-model="form.decryptBox" autocomplete="off" />
          </el-form-item>
          <el-form-item label="启用解密盒子">
            <!-- <el-input v-model="form.enableDecryptBox" autocomplete="off" /> -->
            <el-switch v-model="form.enableDecryptBox" />
          </el-form-item>
          <el-form-item label="tcp端口">
            <el-input v-model="form.tcpPort" autocomplete="off" />
          </el-form-item>
          <el-form-item label="udp端口">
            <el-input v-model="form.udpPort" autocomplete="off" />
          </el-form-item>
        </el-form>
      </el-drawer>
    </div>
  </template>
  
  <script setup>
    import {
      createDeviceConfig,
      updateDeviceConfig,
      deleteDeviceConfig,
      getDeviceConfig,
      getDeviceConfigList,
      playVideo
    } from '@/api/deviceConfig'
    import WarningBar from '@/components/warningBar/warningBar.vue'
    import { ref } from 'vue'
    import { ElMessage, ElMessageBox } from 'element-plus'
    import { formatDate } from '@/utils/format'
  
    defineOptions({
      name: 'DeviceConfig'
    })
  
    const form = ref({
      camera: '',
      encryptBox: '',
      enableEncryptBox: true,
      decryptBox: '',
      enableDecryptBox: true,
      tcpPort: '',
      udpPort: ''
    })
  
    const page = ref(1)
    const total = ref(0)
    const pageSize = ref(10)
    const tableData = ref([])
  
    // 分页
    const handleSizeChange = (val) => {
      pageSize.value = val
      getTableData()
    }
  
    const handleCurrentChange = (val) => {
      page.value = val
      getTableData()
    }
  
    // 查询
    const getTableData = async () => {
      const table = await getDeviceConfigList({
        page: page.value,
        pageSize: pageSize.value
      })
      if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
      }
    }
  
    getTableData()
  
    const drawerFormVisible = ref(false)
    const type = ref('')
    const updateDevice = async (row) => {
      const res = await getDeviceConfig({ id: row.id })
      type.value = 'update'
      if (res.code === 0) {
        form.value = res.data
        drawerFormVisible.value = true
      }
    }
    const closeDrawer = () => {
      drawerFormVisible.value = false
      form.value = {
        camera: '',
        encryptBox: '',
        enableEncryptBox: true,
        decryptBox: '',
        enableDecryptBox: true,
        tcpPort: '',
        udpPort: ''
      }
    }
    const deleteDevice = async (row) => {
      ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        const res = await deleteDeviceConfig({ id: row.id })
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '删除成功'
          })
          if (tableData.value.length === 1 && page.value > 1) {
            page.value--
          }
          getTableData()
        }
      })
    }
    const handlePlayVideo = async (row) => {
      if (!row.decryptBox || !row.tcpPort) {
        ElMessage.error('解密盒子IP或TCP端口未配置');
        return;
      }

      // 调用 deviceConfig.js 中的 playVideo 方法
      await playVideo(row)
        .then(() => {
          ElMessage.success('播放设置成功');
        })
        .catch((error) => {
          ElMessage.error('播放设置失败');
          console.error('播放设置错误:', error);
        });

      // const rtspUrl = `rtsp://admin:1qa2ws3ed@${row.decryptBox}:${row.tcpPort}/Streaming/Channels/1`;
      // // 打开视频流
      // window.open(rtspUrl, '_blank');
    }
    const enterDrawer = async () => {
      let res
      switch (type.value) {
        case 'create':
          res = await createDeviceConfig(form.value)
          break
        case 'update':
          res = await updateDeviceConfig(form.value)
          break
        default:
          res = await createDeviceConfig(form.value)
          break
      }
  
      if (res.code === 0) {
        closeDrawer()
        getTableData()
      }
    }
    const openDrawer = () => {
      type.value = 'create'
      drawerFormVisible.value = true
    }
  </script>
  
  <style></style>
  