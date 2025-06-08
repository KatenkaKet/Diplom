import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Chat, Message } from '@/types/chat'

export const useChatStore = defineStore('chat', () => {
  const chats = ref<Chat[]>([])
  const messages = ref<Message[]>([])
  const activeChatId = ref<string | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const fetchChats = async () => {
    try {
      isLoading.value = true
      error.value = null
      
      const token = localStorage.getItem('token')
      if (!token) {
        throw new Error('No authentication token found')
      }

      const { $chatApi } = useNuxtApp()
      console.log('Fetching chats...')

      const res = await $chatApi.get('/users/search', {
        params: { query: "*" },
        headers: {
          Authorization: `Bearer ${token}`
        }
      })

      if (!res.data) {
        throw new Error('No data received from server')
      }

      console.log('Chats received:', res.data)
      chats.value = res.data
    } catch (error: any) {
      console.error('Ошибка загрузки чатов:', error)
      error.value = error instanceof Error ? error.message : 'Failed to load chats'
      chats.value = []
    } finally {
      isLoading.value = false
    }
  }

  const fetchMessages = async (chatId: string) => {
    try {
      isLoading.value = true
      error.value = null

      const token = localStorage.getItem('token')
      if (!token) {
        throw new Error('No authentication token found')
      }

      const { $chatApi } = useNuxtApp()
      console.log('Fetching messages for chat:', chatId)

      const res = await $chatApi.get(`/chats/${chatId}/messages`, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      })

      if (!res.data) {
        throw new Error('No data received from server')
      }

      console.log('Messages received:', res.data)
      messages.value = res.data
      activeChatId.value = chatId
    } catch (error) {
      console.error('Ошибка загрузки сообщений:', error)
      error.value = error instanceof Error ? error.message : 'Failed to load messages'
      messages.value = []
    } finally {
      isLoading.value = false
    }
  }

  const sendMessage = async (content: string) => {
    if (!activeChatId.value) {
      error.value = 'No active chat selected'
      return
    }

    try {
      isLoading.value = true
      error.value = null

      const token = localStorage.getItem('token')
      if (!token) {
        throw new Error('No authentication token found')
      }

      const { $chatApi } = useNuxtApp()
      console.log('Sending message to chat:', activeChatId.value)

      const res = await $chatApi.post('/messages', {
        chat_id: activeChatId.value,
        content
      }, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      })

      if (!res.data) {
        throw new Error('No response data received')
      }

      console.log('Message sent successfully:', res.data)
      messages.value.push(res.data as Message)
    } catch (error) {
      console.error('Ошибка отправки сообщения:', error)
      error.value = error instanceof Error ? error.message : 'Failed to send message'
    } finally {
      isLoading.value = false
    }
  }

  return {
    chats,
    messages,
    activeChatId,
    isLoading,
    error,
    fetchChats,
    fetchMessages,
    sendMessage
  }
})
