<template>
  <div class="chat">
    <div class="scroll">
      <div class="messages">
        <Message
            v-for="(message, i) in store.messages"
            :key="i"
            :message="message">
        </Message>
      </div>
    </div>
    <n-input
        v-model:value="content"
        @keydown.enter.exact.prevent="sendMessage"
        placeholder="Type your message here..."
        type="textarea" autosize
        :status="status"
    >
    </n-input>
  </div>
</template>

<script setup lang="ts">
import {NInput} from "naive-ui";
import {mainStore} from "@/plugins/store";
import {nextTick, ref, watch} from "vue";
import {ChatMessage} from "@/models/chat";
import Message from "@/home/chat/ChatMessage.vue";

const store = mainStore()

const content = ref('')
const status = ref('')
const sendMessage = () => {
  content.value = content.value.trim()
  if (content.value === '') {
    status.value = 'error'
    return
  } else {
    status.value = ''
  }
  store.sig.send(JSON.stringify({
    event: 'chat',
    data: new ChatMessage(content.value)
  }))
  content.value = ''
}

watch(store.messages, async () => {
  const scroll = document.querySelector('.scroll')!
  await nextTick()
  scroll.scrollTo({top: scroll.scrollHeight, behavior: 'smooth'})
})

</script>

<style scoped>
.chat {
  width: 75%;
  min-width: 200px;
  min-height: 200px;
  display: flex;
  flex-direction: column;
}

.scroll {
  overflow-y: auto;
  padding: 0 1rem;
  margin: 0.5rem 0;
}

.scroll::-webkit-scrollbar {
  width: 8px;
}

.scroll::-webkit-scrollbar-track {
  background: white;
  border-radius: 2px;
}

.scroll::-webkit-scrollbar-thumb {
  background: rgb(220, 220, 220);
  border-radius: 8px;
}

.scroll::-webkit-scrollbar-thumb:hover {
  background: rgb(180, 180, 180);
}

.scroll::-webkit-scrollbar-corner {
  background: #179a16;
}

</style>