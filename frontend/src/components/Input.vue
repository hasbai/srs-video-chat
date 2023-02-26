<template>
  <div id="input">
    <div class="flex justify-between top">
      <n-input placeholder="room" v-model:value="store.room" maxlength="16" show-count clearable>
        <template #prefix>
          <n-icon>
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024"><path d="M946.5 505L560.1 118.8l-25.9-25.9a31.5 31.5 0 0 0-44.4 0L77.5 505a63.9 63.9 0 0 0-18.8 46c.4 35.2 29.7 63.3 64.9 63.3h42.5V940h691.8V614.3h43.4c17.1 0 33.2-6.7 45.3-18.8a63.6 63.6 0 0 0 18.7-45.3c0-17-6.7-33.1-18.8-45.2zM568 868H456V664h112v204zm217.9-325.7V868H632V640c0-22.1-17.9-40-40-40H432c-22.1 0-40 17.9-40 40v228H238.1V542.3h-96l370-369.7l23.1 23.1L882 542.3h-96.1z" fill="currentColor"></path></svg>
          </n-icon>
        </template>
      </n-input>
      <div class="w-16"></div>
      <n-input placeholder="user" v-model:value="store.username" maxlength="16" show-count clearable>
        <template #prefix>
          <n-icon>
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024"><path d="M858.5 763.6a374 374 0 0 0-80.6-119.5a375.63 375.63 0 0 0-119.5-80.6c-.4-.2-.8-.3-1.2-.5C719.5 518 760 444.7 760 362c0-137-111-248-248-248S264 225 264 362c0 82.7 40.5 156 102.8 201.1c-.4.2-.8.3-1.2.5c-44.8 18.9-85 46-119.5 80.6a375.63 375.63 0 0 0-80.6 119.5A371.7 371.7 0 0 0 136 901.8a8 8 0 0 0 8 8.2h60c4.4 0 7.9-3.5 8-7.8c2-77.2 33-149.5 87.8-204.3c56.7-56.7 132-87.9 212.2-87.9s155.5 31.2 212.2 87.9C779 752.7 810 825 812 902.2c.1 4.4 3.6 7.8 8 7.8h60a8 8 0 0 0 8-8.2c-1-47.8-10.9-94.3-29.5-138.2zM512 534c-45.9 0-89.1-17.9-121.6-50.4S340 407.9 340 362c0-45.9 17.9-89.1 50.4-121.6S466.1 190 512 190s89.1 17.9 121.6 50.4S684 316.1 684 362c0 45.9-17.9 89.1-50.4 121.6S557.9 534 512 534z" fill="currentColor"></path></svg>
          </n-icon>
        </template>
      </n-input>
    </div>
    <div class="h-4"></div>
    <n-input placeholder="webrtc" v-model:value="webrtc" clearable>
      <template #prefix>
        <n-icon>
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path d="M16 7a7.66 7.66 0 0 1 1.51.15a8 8 0 0 1 6.35 6.34l.26 1.35l1.35.24a5.5 5.5 0 0 1-1 10.92H7.5a5.5 5.5 0 0 1-1-10.92l1.34-.24l.26-1.35A8 8 0 0 1 16 7m0-2a10 10 0 0 0-9.83 8.12A7.5 7.5 0 0 0 7.49 28h17a7.5 7.5 0 0 0 1.32-14.88a10 10 0 0 0-7.94-7.94A10.27 10.27 0 0 0 16 5z" fill="currentColor"></path></svg>
        </n-icon>
      </template>
      </n-input>
  </div>
</template>

<script setup lang="ts">
import {NIcon, NInput} from "naive-ui"
import {randomHex} from '@/utils'
import {useRoute} from 'vue-router'
import {configStore} from "@/plugins/store";
import {computed, ref} from "vue";

const route = useRoute()
const store = configStore()

store.room = route.params.room as string || store.room || randomHex()
store.username = store.username || randomHex()

let host = ref(window.location.host)
const webrtc = computed({
  get() {
    return `webrtc://${host.value}/${store.room}/${store.username}`
  },
  set(value){
    const values = value.split('/')
    if (values.length != 5) {
      return
    }
      host.value = values[values.length-3]
    store.room = values[values.length-2]
    store.username = values[values.length-1]
  }
})

</script>
<style scoped>
#input {

}

.top .n-input {
  min-width: 6rem;
}

</style>