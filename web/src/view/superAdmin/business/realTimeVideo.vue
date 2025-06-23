<template>
  <div class="video-stream-container">
    <div class="video-wrapper">
      <video
        ref="videoPlayer"
        class="video-js vjs-default-skin vjs-big-play-centered"
        controls
        preload="auto"
        width="640"
        height="360"
      >
        <source :src="videoUrl" type="application/x-mpegURL" />
      </video>
    </div>
  </div>
</template>

<script>
import 'video.js/dist/video-js.css'
import videojs from 'video.js'

export default {
  name: 'VideoStream',
  data() {
    return {
      player: null,
      videoUrl: 'http://192.168.1.35:8080/stream/index.m3u8' // This will be the HLS stream URL
    }
  },
  mounted() {
    this.initializePlayer()
  },
  beforeDestroy() {
    if (this.player) {
      this.player.dispose()
    }
  },
  methods: {
    initializePlayer() {
      this.player = videojs(this.$refs.videoPlayer, {
        fluid: true,
        aspectRatio: '16:9',
        autoplay: true
      })
    }
  }
}
</script>

<style scoped>
.video-stream-container {
  padding: 20px;
  background: #f0f2f5;
  min-height: 100vh;
}

.video-wrapper {
  max-width: 1280px;
  margin: 0 auto;
  background: #fff;
  padding: 20px;
  border-radius: 4px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.video-js {
  width: 100%;
  height: auto;
}
</style> 