<!-- components/message/DialogBox.vue -->
<template>
  <v-card class="fill-height d-flex flex-column">
    <v-card-title>
      {{ props.companionName }}
    </v-card-title>

    <v-divider></v-divider>

    <v-card-text class="flex-grow-1 overflow-auto">
      <div v-if="isLoading" class="d-flex justify-center align-center pa-4">
        <v-progress-circular indeterminate color="primary" />
      </div>
      <div v-else-if="error" class="d-flex justify-center align-center pa-4">
        <div class="w-100 text-center text-grey text-body-1" style="background: #f5f5f5; border-radius: 8px; padding: 16px;">
          {{ error }}
        </div>
      </div>
      <MessageList v-else :messages="messages" :currentUserId="currentUserId" />
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
        :disabled="isLoading"
      />
    </v-card-actions>
  </v-card>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, computed } from 'vue'
import MessageList from './MessageList.vue'
import type { Message } from '@/types/chat'
import { useAuthStore } from '@/stores/auth'

const props = defineProps<{ 
  chatId: string,
  companionName: string 
}>()

const { $chatApi } = useNuxtApp()
const authStore = useAuthStore()

const message = ref('')
const messages = ref<Message[]>([])
const isLoading = ref(false)
const error = ref<string | null>(null)
const currentUserId = computed(() => authStore.user?.id || 0)

// Загрузка сообщений
async function loadMessages() {
  if (!props.chatId) return
  
  isLoading.value = true
  error.value = null
  
  console.log('Loading messages for chat:', props.chatId)
  console.log('Auth token:', authStore.token)
  
  try {
    const response = await $chatApi.get(`/chats/${props.chatId}/messages`, {
      headers: {
        Authorization: `Bearer ${authStore.token}`
      }
    })
    console.log('API Response:', response.data)
    // Логируем структуру сообщений
    console.log('API messages:', response.data)
    messages.value = response.data.map(msg => ({
      ...msg,
      createdAt: msg.createdAt || msg.created_at,
      sender: {
        id: msg.sender_id,
        username: msg.sender?.username || '',
        avatarUrl: msg.sender?.avatar_url || ''
      }
    }))
  } catch (e: any) {
    console.error('Full error:', e)
    console.error('Error response:', e.response)
    console.error('Error status:', e.response?.status)
    console.error('Error data:', e.response?.data)
    error.value = 'Диалог ещё не начат'
    console.error('Ошибка загрузки сообщений:', e)
  } finally {
    isLoading.value = false
  }
}

// Отправка сообщения
async function sendMessage() {
  if (!message.value.trim() || !props.chatId) return
  
  try {
    const response = await $chatApi.post('/messages', {
      chat_id: props.chatId,
      content: message.value
    }, {
      headers: {
        Authorization: `Bearer ${authStore.token}`
      }
    })
    
    messages.value.push(response.data)
    message.value = ''
  } catch (e: any) {
    error.value = e.response?.data?.error || 'Ошибка при отправке сообщения'
    console.error('Ошибка отправки сообщения:', e)
  }
}

// Загружаем сообщения при монтировании и при изменении chatId
onMounted(loadMessages)
watch(() => props.chatId, loadMessages)
</script>

<style scoped>
.fill-height {
  height: 100%;
}
</style>