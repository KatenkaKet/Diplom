<!-- components/message/MessageItem.vue -->
<template>
  <div :class="['message-item', isMine ? 'mine' : 'theirs']">
    <v-avatar v-if="!isMine" size="24" class="mr-2">
      <v-img :src="message.sender.avatarUrl" @error="imgErr = true" v-if="!imgErr" />
      <v-icon v-else>mdi-account</v-icon>
    </v-avatar>

    <div class="message-bubble">
      <div class="text-body-2">{{ message.content }}</div>
      <div class="timestamp">{{ new Date(message.createdAt).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }) }}</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { Message } from '@/types/chat'

defineProps<{
  message: Message
  isMine: boolean
}>()

const imgErr = ref(false)
</script>

<style scoped>
.message-item {
  display: flex;
  align-items: flex-end;
  max-width: 75%;
}
.mine {
  align-self: flex-end;
  flex-direction: row-reverse;
  text-align: right;
}
.theirs {
  align-self: flex-start;
}
.message-bubble {
  background-color: #f0f0f0;
  border-radius: 12px;
  padding: 6px 10px;
  max-width: 100%;
  word-break: break-word;
}
.timestamp {
  font-size: 0.7rem;
  color: gray;
  margin-top: 2px;
}
</style>
