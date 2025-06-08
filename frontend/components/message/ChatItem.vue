
<template>
  <!-- @vue-ignore -->
  <v-list-item @click="$emit('selectChat', chat._id ? chat._id : chat.id)" class="chat-item" clickable>
    <v-row align="center" no-gutters class="w-100">
      <v-col cols="3" class="d-flex justify-center">
        <v-avatar size="40">
          <v-icon>mdi-account</v-icon>
        </v-avatar>
      </v-col>

      <v-col cols="9">
        <div class="d-flex flex-column">
          <span class="font-weight-medium">
            {{ username || 'Неизвестный' }}
          </span>
          <span class="text-caption text-grey">@{{ username || '-' }}</span>
          <span v-if="chat.lastMessage" class="text-truncate text-grey text-body-2 mt-1">
            {{ chat.lastMessage }}
          </span>
        </div>
      </v-col>
    </v-row>
  </v-list-item>
</template>

<script setup lang="ts">
import type { Chat } from '@/types/chat'

const props = defineProps<{ chat: Chat }>()

// @ts-ignore
const username = props.chat.partner?.username || props.chat.members?.[0]?.username  || props.chat.username
</script>
