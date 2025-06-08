export interface User {
  id: number            // ID из PostgreSQL
  username: string      // Никнейм
  avatarUrl: string     // URL аватарки
}

export interface Chat {
  _id: string
  title?: string
  avatarUrl?: string
  members: {
    id: number
    username: string
    avatarUrl: string
  }[]
  lastMessage: string
  updatedAt: string

  partner?: {
    id: number
    username: string
    avatarUrl: string
  }
}



export interface Message {
  _id: string                // ID сообщения
  chatId: string             // ID чата, к которому принадлежит сообщение
  sender: User               // Кто отправил сообщение
  content: string            // Текст сообщения
  createdAt: string          // Время отправки
}