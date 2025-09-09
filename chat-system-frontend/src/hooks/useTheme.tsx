import { useEffect } from "react"

export function useTheme(): {
    isSystemDarkMode: () => boolean,
    initializeTheme: () => void,
    setupThemeListener: () => (() => void),
    ThemeProvider: React.FC<{ children: React.ReactNode }>
} {
    // 检查当前系统是否为暗黑模式
    const isSystemDarkMode = (): boolean => {
        return window.matchMedia('(prefers-color-scheme: dark)').matches
    }

    // 初始化主题，跟随系统设置
    const initializeTheme = (): void => {
        if (isSystemDarkMode()) {
            document.documentElement.classList.add('dark')
        } else {
            document.documentElement.classList.remove('dark')
        }
    }

    // 监听系统主题变化
    const setupThemeListener = (): (() => void) => {
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

    // 创建一个组件来处理主题初始化和清理
    const ThemeProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
        useEffect(() => {
            initializeTheme()
            const cleanup = setupThemeListener()

            return cleanup
        }, [])

        return <>{children}</>
    }

    return {
        isSystemDarkMode,
        initializeTheme,
        setupThemeListener,
        ThemeProvider
    }
}
