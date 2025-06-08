<!-- components/message/DialogBox.vue -->
<template>
  <v-card class="fill-height d-flex flex-column">
    <v-card-title>
      Чат ID: {{ props.chatId }}
    </v-card-title>

    <v-divider></v-divider>

    <v-card-text class="flex-grow-1 overflow-auto">
      <MessageList :messages="messages" :currentUserId="1" />
    </v-card-text>

    <v-divider></v-divider>

    <v-card-actions>
      <v-text-field
        v-model="message"
        density="compact"
        placeholder="Введите сообщение..."
        hide-details
        class="flex-grow-1"
        @keyup.enter="sendMessage"
      />
    </v-card-actions>
  </v-card>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import MessageList from './MessageList.vue'
import type { Message } from '@/types/chat'

const props = defineProps<{ chatId: string }>()

const message = ref('')
const currentUserId = 1

const messages = ref<Message[]>([
  {
    _id: 'm1',
    chatId: '1',
    sender: { id: 1, username: 'me', avatarUrl: '' },
    content: 'Привет!',
    createdAt: new Date().toISOString()
  },
  {
    _id: 'm2',
    chatId: '1',
    sender: { id: 2, username: 'alice', avatarUrl: '/avatar1.png' },
    content: 'Привет-привет! Как дела?',
    createdAt: new Date().toISOString()
  }
])

function sendMessage() {
  if (!message.value.trim()) return

  messages.value.push({
    _id: Math.random().toString(36).substring(2),
    chatId: props.chatId,
    sender: { id: currentUserId, username: 'me', avatarUrl: '' },
    content: message.value,
    createdAt: new Date().toISOString()
  })

  message.value = ''
}
</script>

<style scoped>
.fill-height {
  height: 100%;
}
</style>