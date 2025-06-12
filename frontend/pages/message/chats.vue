<template>
  <v-container fluid class="fill-height d-flex pa-0">
    <v-row no-gutters class="flex-grow-1 fill-height">
      <!-- Левая колонка: список чатов -->
      <v-col cols="4" class="pa-2 left-panel fill-height">
        <ChatList @selectChat="handleChatSelect" class="fill-height" />
      </v-col>

      <!-- Правая колонка: диалог -->
      <v-col cols="8" class="pa-2 fill-height">
        <DialogBox
          v-if="activeChatId"
          :chat-id="activeChatId"
          :companion-name="activeChatUsername"
          class="fill-height"
        />

        <v-card
          v-else
          class="fill-height d-flex flex-column align-center justify-center"
        >
          <v-card-text class="text-grey">Выберите чат слева</v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import ChatList from '@/components/message/ChatList.vue'
import DialogBox from '@/components/message/DialogBox.vue'

const activeChatId = ref<string | null>(null)
const activeChatUsername = ref<string>('')

function handleChatSelect(data: { chatId: string, username: string }) {
  activeChatId.value = data.chatId
  activeChatUsername.value = data.username
}
</script>

<style scoped>
.fill-height {
  height: 100vh;
  overflow: hidden;
}

.left-panel {
  background-color: #f5f5f5; /* светло-серый фон */
  border-right: 1px solid #ddd; /* тонкая граница справа */
}
</style>
