<template>
  <div class="container flex flex-col">
    <video ref="player" controls autoplay></video>
    <div class="toolbar flex flex-row justify-around">
      <span>{{ client.name }}</span>
      <n-button text @click="sdk.changeMediaStatus('audio')">
        <n-icon size="1.5rem" v-if="sdk.audioEnabled">
          <svg xmlns="http://www.w3.org/2000/svg"
               viewBox="0 0 1024 1024">
            <path
                d="M842 454c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8c0 140.3-113.7 254-254 254S258 594.3 258 454c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8c0 168.7 126.6 307.9 290 327.6V884H326.7c-13.7 0-24.7 14.3-24.7 32v36c0 4.4 2.8 8 6.2 8h407.6c3.4 0 6.2-3.6 6.2-8v-36c0-17.7-11-32-24.7-32H548V782.1c165.3-18 294-158 294-328.1zM512 624c93.9 0 170-75.2 170-168V232c0-92.8-76.1-168-170-168s-170 75.2-170 168v224c0 92.8 76.1 168 170 168zm-94-392c0-50.6 41.9-92 94-92s94 41.4 94 92v224c0 50.6-41.9 92-94 92s-94-41.4-94-92V232z"
                fill="currentColor"></path>
          </svg>
        </n-icon>
        <n-icon size="1.5rem" v-else>
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024">
            <defs></defs>
            <path
                d="M682 455V311l-76 76v68c-.1 50.7-42 92.1-94 92c-19.1.1-36.8-5.4-52-15l-54 55c29.1 22.4 65.9 36 106 36c93.8 0 170-75.1 170-168z"
                fill="currentColor"></path>
            <path
                d="M833 446h-60c-4.4 0-8 3.6-8 8c0 140.3-113.7 254-254 254c-63 0-120.7-23-165-61l-54 54c48.9 43.2 110.8 72.3 179 81v102H326c-13.9 0-24.9 14.3-25 32v36c.1 4.4 2.9 8 6 8h408c3.2 0 6-3.6 6-8v-36c0-17.7-11-32-25-32H547V782c165.3-17.9 294-157.9 294-328c0-4.4-3.6-8-8-8zm13.1-377.7l-43.5-41.9c-3.1-3-8.1-3-11.2.1l-129 129C634.3 101.2 577 64 511 64c-93.9 0-170 75.3-170 168v224c0 6.7.4 13.3 1.2 19.8l-68 68c-10.5-27.9-16.3-58.2-16.2-89.8c-.2-4.4-3.8-8-8-8h-60c-4.4 0-8 3.6-8 8c0 53 12.5 103 34.6 147.4l-137 137c-3.1 3.1-3.1 8.2 0 11.3l42.7 42.7c3.1 3.1 8.2 3.1 11.3 0L846.2 79.8l.1-.1c3.1-3.2 3-8.3-.2-11.4zM417 401V232c0-50.6 41.9-92 94-92c46 0 84.1 32.3 92.3 74.7L417 401z"
                fill="currentColor"></path>
          </svg>
        </n-icon>
      </n-button>
      <n-button text @click="sdk.changeMediaStatus('video')">
        <n-icon size="1.5rem" v-if="sdk.videoEnabled">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
            <path
                d="M15 8v8H5V8h10m1-2H4c-.55 0-1 .45-1 1v10c0 .55.45 1 1 1h12c.55 0 1-.45 1-1v-3.5l4 4v-11l-4 4V7c0-.55-.45-1-1-1z"
                fill="currentColor"></path>
          </svg>
        </n-icon>
        <n-icon size="1.5rem" v-else>
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
            <path
                d="M9.56 8l-2-2l-4.15-4.14L2 3.27L4.73 6H4c-.55 0-1 .45-1 1v10c0 .55.45 1 1 1h12c.21 0 .39-.08.55-.18L19.73 21l1.41-1.41l-8.86-8.86L9.56 8zM5 16V8h1.73l8 8H5zm10-8v2.61l6 6V6.5l-4 4V7c0-.55-.45-1-1-1h-5.61l2 2H15z"
                fill="currentColor"></path>
          </svg>
        </n-icon>
      </n-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import {Client} from "@/models/client";
import {onMounted, reactive, ref} from "vue";
import {configStore} from "@/plugins/store";
import {SrsRtcPlayer} from "@/plugins/srs.sdk";
import {useMessage} from "naive-ui";

const config = configStore()
const message = useMessage()

const {client} = defineProps<{ client: Client }>()
const isMe = ref(client.name === config.username)

const sdk = reactive(new SrsRtcPlayer())
const player = ref<HTMLVideoElement | null>(null)

onMounted(async () => {
  player.value!.srcObject = sdk.stream;
  const webrtcURL = `${config.webrtcURL}/${config.room}/${client.name}`
  try {
    if (!isMe.value) {
      await sdk.play(webrtcURL)
    } else {
      await sdk.publish(webrtcURL)
    }
  } catch (e) {
    let msg = ''
    if (e instanceof DOMException) {
      if (e.name === 'NotFoundError') {
        msg = '找不到麦克风和摄像头设备：'
      } else if (e.name === 'NotAllowedError') {
        msg = '你禁止了网页访问摄像头和麦克风：'
      }
      msg += `${e.name} ${e.message}`
    } else {
      msg = (e as Error).toString()
    }
    message.error(msg, {duration: 5000, closable: true})
  }
})

</script>

<style scoped>
.container {
  border: grey 1px solid;
}

.toolbar {
  font-size: 1.125rem;
}

video {
  max-width: 100%;
  max-height: 100%;
}
</style>