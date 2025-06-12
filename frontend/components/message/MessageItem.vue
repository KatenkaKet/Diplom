<!-- components/message/MessageItem.vue -->
<template>
  <div :class="['message-item', props.isMine ? 'mine' : 'theirs']">
    <div class="message-bubble">
      <div class="text-body-2">{{ props.message.content }}</div>
      <div class="timestamp">{{ new Date(props.message.createdAt).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }) }}</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { Message } from '@/types/chat'

const props = defineProps<{
  message: Message
  isMine: boolean
}>()

console.log('MessageItem', { senderId: props.message.sender.id, isMine: props.isMine })

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
  text-align: right;
}
.theirs {
  align-self: flex-start;
  text-align: left;
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
