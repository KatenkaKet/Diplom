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
          @selectChat="$emit('selectChat', $event)"
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
const emit = defineEmits(["selectChat"]);

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
</script>
