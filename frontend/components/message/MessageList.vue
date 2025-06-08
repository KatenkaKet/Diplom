<!-- components/message/MessageList.vue -->
<template>
  <div class="message-list">
    <MessageItem
      v-for="msg in props.messages"
      :key="msg._id"
      :message="msg"
      :isMine="msg.sender.id === props.currentUserId"
    />
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'
import MessageItem from './MessageItem.vue'
import type { Message } from '@/types/chat'

const props = defineProps<{
  messages: Message[]
  currentUserId: number
}>()
</script>

<style scoped>
.message-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
  padding: 16px;
  overflow-y: auto;
  height: 100%;
}
</style>
