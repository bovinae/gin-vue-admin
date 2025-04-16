<template>
    <div>
      <!-- <warning-bar
        title="在资源权限中将此角色的资源权限清空 或者不包含创建者的角色 即可屏蔽此客户资源的显示"
      /> -->
      <div class="gva-table-box">
        <el-table
          ref="multipleTable"
          :data="tableData"
          style="width: 100%"
          tooltip-effect="dark"
          row-key="ID"
        >
          <el-table-column type="selection" width="55" />
          <el-table-column
            align="left"
            label="ID"
            prop="id"
            width="80"
          />
          <el-table-column
            align="left"
            label="文件名"
            prop="name"
            width="270"
          />
          <el-table-column
            align="left"
            label="大小"
            prop="size"
            width="150"
          />
          <el-table-column
            align="left"
            label="修改时间"
            prop="modifyTime"
            width="250"
          />
          <el-table-column align="left" label="操作" min-width="160">
            <template #default="scope">
              <el-button
                type="primary"
                link
                icon="video-play"
                @click="playVideo(scope.row)"
                >播放</el-button
              >
              <el-button
                type="primary"
                link
                icon="video-pause"
                @click="stopVideo(scope.row)"
                >停止</el-button
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

      <!-- Video Player Dialog -->
      <el-dialog
        v-model="dialogVisible"
        title="视频播放"
        width="80%"
        :before-close="handleClose"
      >
        <div class="video-container">
          <video
            ref="videoPlayer"
            controls
            autoplay
            preload="auto"
            width="100%"
            height="500"
          >
          </video>
        </div>
      </el-dialog>
    </div>
  </template>
  
<script>
import { ref, onMounted, onBeforeUnmount, nextTick } from 'vue'
import axios from 'axios'
import { ElMessage } from 'element-plus'
// Remove video.js CSS import
// import 'video.js/dist/video-js.css'
// Remove video.js import
// import videojs from 'video.js'

// Import hls.js
import Hls from 'hls.js'

export default {
  name: 'HistoricalVideo',
  setup() {
    const tableData = ref([])
    const loading = ref(false)
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)
    const dialogVisible = ref(false)
    const videoPlayer = ref(null) // Ref for the <video> element
    // Remove video.js player ref
    // const player = ref(null)
    const videoUrl = ref('')
    const hlsInstance = ref(null) // Ref for hls.js instance

    const fetchPcapFiles = async () => {
      loading.value = true
      try {
        const response = await axios.get('http://192.168.1.39:5000/list_pcaps', {
          params: {
            page: page.value,
            pageSize: pageSize.value
          }
        })
        if (response.data.code === 0) {
          tableData.value = response.data.data.list
          total.value = response.data.data.total
        } else {
          ElMessage.error(response.data.msg || 'Failed to fetch pcap files')
        }
      } catch (error) {
        ElMessage.error('Failed to fetch pcap files: ' + error.message)
      } finally {
        loading.value = false
      }
    }

    const playVideo = async (row) => {
      loading.value = true
      try {
        // First, call the POST interface to start the replay
        const formData = new URLSearchParams()
        formData.append('pcap_file', row.name)
        
        // Send request and wait for response
        const replayResponse = await axios.post('http://192.168.1.39:5000/replay', formData, {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
          }
        })

        if (replayResponse.data.code !== 0) {
          throw new Error(replayResponse.data.msg || '启动回放失败')
        }

        // Wait for 60 seconds to ensure the server has processed the replay request
        await new Promise(resolve => setTimeout(resolve, 60000))

        // Construct the dynamic playlist URL
        const fileNameWithoutExtension = row.name.replace(/\.pcap$/, '');
        videoUrl.value = `http://192.168.1.39:5000/${fileNameWithoutExtension}.m3u8`
        console.log('Using HLS playlist URL:', videoUrl.value); // Log the URL for debugging
        dialogVisible.value = true

        // Initialize hls.js after the dialog is shown and DOM is updated
        await nextTick() // Wait for Vue to update the DOM

        const videoElement = videoPlayer.value
        if (!videoElement) {
          console.error('Video element not found!')
          return
        }

        // Cleanup previous hls instance if exists
        if (hlsInstance.value) {
          hlsInstance.value.destroy()
        }

        if (Hls.isSupported()) {
          console.log('HLS is supported, initializing hls.js')
          const hls = new Hls()
          hlsInstance.value = hls // Store the instance
          hls.loadSource(videoUrl.value)
          hls.attachMedia(videoElement)
          hls.on(Hls.Events.MANIFEST_PARSED, function () {
            console.log('Manifest parsed, playing video')
            videoElement.play().catch(e => console.error('Autoplay failed:', e))
          })
          hls.on(Hls.Events.ERROR, function (event, data) {
            console.error('HLS.js Error:', data)
            if (data.fatal) {
              switch (data.type) {
                case Hls.ErrorTypes.NETWORK_ERROR:
                  console.error('Fatal network error encountered, trying to recover')
                  hls.startLoad()
                  break
                case Hls.ErrorTypes.MEDIA_ERROR:
                  console.error('Fatal media error encountered, trying to recover')
                  hls.recoverMediaError()
                  break
                default:
                  // Cannot recover
                  ElMessage.error(`视频播放错误: ${data.details}`)
                  hls.destroy()
                  hlsInstance.value = null
                  break
              }
            }
          })
        } else if (videoElement.canPlayType('application/vnd.apple.mpegurl')) {
          // Native HLS support (e.g., Safari)
          console.log('Native HLS support detected')
          videoElement.src = videoUrl.value
          videoElement.addEventListener('loadedmetadata', function () {
            videoElement.play().catch(e => console.error('Autoplay failed:', e))
          })
        } else {
          ElMessage.error('浏览器不支持 HLS 播放')
        }

      } catch (error) {
        console.error('Play video error:', error)
        ElMessage.error('启动回放失败: ' + error.message)
      } finally {
        loading.value = false
      }
    }

    const stopVideo = async (row) => {
      loading.value = true
      try {
        // Call the stop API
        await axios.post('http://192.168.1.39:5000/stop')

        // Close the video player dialog if it's open
        if (dialogVisible.value) {
          handleClose()
        }

        ElMessage.success('已停止回放')
      } catch (error) {
        ElMessage.error('停止回放失败: ' + error.message)
      } finally {
        loading.value = false
      }
    }

    const handleClose = async () => {
      // Call the stop API when closing the dialog
      try {
        console.log('Closing dialog, calling stop API...')
        await axios.post('http://192.168.1.39:5000/stop')
        console.log('Stop API called successfully.')
      } catch (error) {
        console.error('Failed to call stop API on close:', error)
        // Optionally show a non-critical error message, or just log it
        // ElMessage.error('停止回放时出错: ' + error.message)
      }

      // Destroy hls.js instance on dialog close
      if (hlsInstance.value) {
        console.log('Destroying HLS instance')
        hlsInstance.value.destroy()
        hlsInstance.value = null
      }
      // Stop video playback if using native HLS
      const videoElement = videoPlayer.value
      if(videoElement) {
        videoElement.pause()
        videoElement.removeAttribute('src') // Prevent memory leak
        videoElement.load()
      }

      dialogVisible.value = false
    }

    const handleCurrentChange = (val) => {
      page.value = val
      fetchPcapFiles()
    }

    const handleSizeChange = (val) => {
      pageSize.value = val
      fetchPcapFiles()
    }

    onMounted(() => {
      fetchPcapFiles()
    })

    onBeforeUnmount(() => {
      // Ensure cleanup when component is destroyed
      handleClose()
    })

    return {
      tableData,
      loading,
      page,
      pageSize,
      total,
      dialogVisible,
      videoPlayer,
      videoUrl,
      playVideo,
      stopVideo,
      handleClose,
      handleCurrentChange,
      handleSizeChange
      // hlsInstance is managed internally, no need to return
    }
  }
}
</script>

<style scoped>
.pcap-list-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.box-card {
  margin-bottom: 20px;
}

.video-container {
  width: 100%;
  max-width: 100%;
  margin: 0 auto;
}

/* Remove video.js specific styles if any were added */
</style> 