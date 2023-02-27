<template>
  <div class="chat mt-8">
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
import Message from "@/components/ChatMessage.vue";

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
  store.sig.send({
    event: 'chat',
    data: new ChatMessage(content.value)
  })
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
  height: 0;
  flex: 1;
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
  background: rgb(239, 239, 239);
  border-radius: 2px;
}

.scroll::-webkit-scrollbar-thumb {
  background: #bfbfbf;
  border-radius: 10px;
}

.scroll::-webkit-scrollbar-thumb:hover {
  background: #333;
}

.scroll::-webkit-scrollbar-corner {
  background: #179a16;
}

</style>