<template>
  <div
      class="my-2 container"
      :class="[isMe ? 'right': 'left']"
  >
    <div v-if="!isMe" class="name">{{ message?.from }}</div>
    <div
        class="flex justify-start"
        :class="[isMe ? 'flex-row-reverse': '']">
      <div class="message">
        {{ message?.content }}
      </div>
      <span class="time flex flex-col justify-end mx-2">{{
          new Date(message?.createdAt).toLocaleTimeString()
        }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ChatMessage} from "@/models/chat";
import {configStore} from "@/plugins/store";

const {message} = defineProps({message: ChatMessage})
const config = configStore()
const isMe = message?.from === config.username

</script>

<style scoped>
.message {
  position: relative;
  width: max-content;
  padding: 0.5rem;
  border-radius: 0.5rem;
}

.message::after {
  content: '';
  position: absolute;
  bottom: 0;
  width: 2rem;
  height: 2rem;
  background-color: white;
}

.message::before {
  content: '';
  position: absolute;
  bottom: 0;
  width: 1rem;
  height: 1rem;
  background-color: inherit;
}

.right .message {
  margin-right: 1rem;
  border-bottom-right-radius: 0;
  background-color: #9ccc65;
}

.right .message::before {
  right: -1rem;
}

.right .message::after {
  right: -2rem;
  border-radius: 0 0 0 50%;
}

.left .message {
  margin-left: 1rem;
  border-bottom-left-radius: 0;
  background-color: skyblue;
}

.left .message::before {
  left: -1rem;
}

.left .message::after {
  left: -2rem;
  border-radius: 0 0 50% 0;
}

.name {
  margin-left: 1.25rem;
}

.message:hover + .time {
  display: flex;
}

.time {
  color: #9e9e9e;
  display: none;
  vertical-align: bottom;
}

</style>