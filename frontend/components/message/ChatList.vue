<template>
  <div>
    <v-text-field
      v-model="search"
      density="compact"
      placeholder="Поиск по нику"
      prepend-inner-icon="mdi-magnify"
      class="mb-2"
      hide-details
    />

    <!-- Состояние загрузки -->
    <div v-if="chatStore.isLoading" class="d-flex justify-center align-center pa-4">
      <v-progress-circular indeterminate color="primary" />
    </div>

    <!-- Сообщение об ошибке -->
    <v-alert
      v-else-if="chatStore.error"
      type="error"
      class="ma-2"
      closable
    >
      {{ chatStore.error }}
    </v-alert>

    <!-- Сообщение о необходимости авторизации -->
    <v-alert
      v-else-if="!authStore.isAuthenticated()"
      type="warning"
      class="ma-2"
    >
      Пожалуйста, войдите в систему
    </v-alert>

    <!-- Список чатов -->
    <v-list v-else lines="one">
      <template v-if="displayedChats.length > 0">
        
        <ChatItem
          v-for="chat in displayedChats"
          :key="chat._id ? chat._id : chat.id"
          :chat="chat"
          :user-id="authStore.user?.id"
          @selectChat="handleSelectChat"
        />
      </template>
      <v-list-item v-else>
        <v-list-item-title class="text-center text-grey">
          {{ search.length >= 2 ? 'Ничего не найдено' : 'Нет доступных чатов' }}
        </v-list-item-title>
      </v-list-item>
    </v-list>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, computed, onMounted } from "vue";
import ChatItem from "./ChatItem.vue";
import type { Chat } from "@/types/chat";
import { useChatStore } from "@/stores/chat";
import { useAuthStore } from "@/stores/auth";

const { $chatApi } = useNuxtApp();
const emit = defineEmits<{
  (e: 'selectChat', data: { chatId: string, username: string }): void
}>();

const search = ref("");
const searchResults = ref<Chat[]>([]);
const chatStore = useChatStore();
const authStore = useAuthStore();

onMounted(async () => {
  if (process.client) {
    authStore.initAuth();
    if (authStore.isAuthenticated()) {
      await chatStore.fetchChats();
    }
  }
});

watch(search, async (value) => {
  if (value.length >= 2) {
    try {
      const res = await $chatApi.get("/users/search", {
        params: { query: value },
        headers: {
          Authorization: `Bearer ${authStore.token}`
        }
      });
      searchResults.value = res.data.map((user: any) => ({
        _id: user.id,
        members: [{ username: user.username, avatar_url: user.avatar_url }]
      }));
    } catch (e) {
      console.error("Ошибка поиска:", e);
      searchResults.value = [];
    }
  } else {
    searchResults.value = [];
  }
});

// Показываем searchResults если есть поисковый запрос, иначе показываем все чаты
const displayedChats = computed(() => {
  console.log(search.value.length >= 2 ? searchResults.value : chatStore.chats);
  return search.value.length >= 2 ? searchResults.value : chatStore.chats;
});

async function handleSelectChat(chatIdOrUserId: string, username: string) {
  console.log('handleSelectChat called with:', chatIdOrUserId, username)
  // Если это обычный чат (есть в chatStore.chats), просто эмитим
  const chat = chatStore.chats.find(c => c._id === chatIdOrUserId)
  console.log('Found chat in chatStore:', chat)
  if (chat) {
    emit('selectChat', { chatId: chat._id, username })
    console.log('Emitted selectChat for existing chat:', chat._id, username)
    return
  }

  // Если это пользователь из поиска — создаём/ищем приватный чат
  try {
    if (!authStore.user) {
      console.error('authStore.user is null!')
      return
    }
    console.log('Creating/finding private chat with members:', [authStore.user.id, chatIdOrUserId])
    const res = await $chatApi.post('/chats', {
      type: 'private',
      members: [authStore.user.id, chatIdOrUserId]
    }, {
      headers: {
        Authorization: `Bearer ${authStore.token}`
      }
    })
    const chatObj = res.data
    console.log('Received chatObj from API:', chatObj)
    emit('selectChat', { chatId: chatObj.id, username })
    console.log('Emitted selectChat for new chat:', chatObj.id, username)
  } catch (e) {
    console.error('Ошибка создания/поиска приватного чата:', e)
  }
}
</script>
