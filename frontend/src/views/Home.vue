<template>
  <main id="main">
    <Input/>
  </main>
</template>


<script setup lang="ts">
import {onMounted} from "vue"
import {SrsRtcSignaling} from "@/websocket"
import Input from "@/components/Input.vue";

async function ws(room: string, user: string) {
  const sig = new SrsRtcSignaling()
  await sig.connect(room, user)
}

onMounted(async () => {
  console.log('mounted')
})

</script>

<style scoped>
</style>
