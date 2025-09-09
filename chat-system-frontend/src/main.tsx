import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import {useTheme} from './hooks/useTheme'

const { ThemeProvider } = useTheme()

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ThemeProvider>
        <App />
    </ThemeProvider>
  </StrictMode>,
)
