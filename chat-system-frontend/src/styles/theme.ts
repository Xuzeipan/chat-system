// 检查当前系统是否为暗黑模式
const isSystemDarkMode = (): boolean => {
    return window.matchMedia('(prefers-color-scheme: dark)').matches
}

// 初始化主题，跟随系统设置
export const initializeTheme = (): void => {
    if (isSystemDarkMode()) {
        document.documentElement.classList.add('dark')
    } else {
        document.documentElement.classList.remove('dark')
    }
}

// 监听系统主题变化
export const setupThemeListener = (): (() => void) => {
  const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
  
  const handleChange = (e: MediaQueryListEvent) => {
    if (e.matches) {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
  }
  
  mediaQuery.addEventListener('change', handleChange)
  
  // 返回清理函数
  return () => mediaQuery.removeEventListener('change', handleChange)
}