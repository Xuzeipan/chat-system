import { useEffect, useState } from "react"

function timeBar() {
    const [currentTime, setCurrentTime] = useState(new Date())

    useEffect(() => {
        const timer = setInterval(() => {
            setCurrentTime(new Date())
        }, 1000)

        return () => clearInterval(timer)
    }, [])

    // 格式化时间
    const formatTime = (date: Date): string => {
        return date.toLocaleString('en-US', {
            weekday: 'long',
            year: 'numeric',
            month: 'long',
            day: 'numeric',
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit',
        })
    }

    return (
        <>
            <div className="p-1 cursor-pointer">
                <p className="text-center">{formatTime(currentTime)}</p>
            </div>
        </>
    )
}

export default timeBar