<template>
  <v-container>
    <!-- Список сообщений -->
    <v-list v-if="chatStore.messages.length">
      <v-list-item
        v-for="msg in chatStore.messages"
        :key="msg._id"
      >
        <v-avatar size="36" class="mr-2">
          <v-img :src="msg.sender.avatarUrl || undefined">
            <template #placeholder>
              <v-icon>mdi-account</v-icon>
            </template>
          </v-img>
        </v-avatar>
        <v-list-item-content>
          <v-list-item-title>{{ msg.sender.username }}</v-list-item-title>
          <v-list-item-subtitle>{{ msg.content }}</v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>
    </v-list>

    <p v-else class="text-muted">Нет сообщений</p>

    <!-- Поле ввода -->
    <v-form @submit.prevent="chatStore.sendMessage(input)">
      <v-text-field
        v-model="input"
        placeholder="Введите сообщение"
        append-icon="mdi-send"
        @click:append="chatStore.sendMessage(input); input = ''"
        density="compact"
      />
    </v-form>
  </v-container>
</template>


<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { useChatStore } from '@/stores/chat'

const input = ref('')
const route = useRoute()
const chatStore = useChatStore()

const chatId = route.params.chatId as string

onMounted(() => {
  chatStore.fetchMessages(chatId)
})

watch(() => route.params.chatId, (newId) => {
  if (newId) chatStore.fetchMessages(newId as string)
})
</script>

