import { createContext, FC, useContext, useState } from 'react'
import { AlertColor } from '@mui/material/Alert'

interface Message {
  text: string
  ttl: number
  variant: AlertColor
}

interface MessageContextType {
  message: Message
  showMessage: (options: Partial<Message>) => void
}

const MessageContext = createContext<MessageContextType>({
  message: { text: '', variant: 'info', ttl: 5000 },
  showMessage: () => {}
})

interface Props {
  children: React.ReactNode
}

const MessageProvider: FC<Props> = ({ children }) => {
  const [message, setMessage] = useState<Message>({ text: '', variant: 'info', ttl: 5000 })
  const showMessage = (messageOptions: Partial<Message>) => setMessage(prev => ({ ...prev, ...messageOptions }))
  
  return (
    <MessageContext.Provider value={{ message, showMessage }}>{children}</MessageContext.Provider>
  )
}

export default MessageProvider

export const useMessage = () => useContext(MessageContext)
