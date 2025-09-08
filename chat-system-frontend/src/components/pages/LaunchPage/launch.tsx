import { motion } from "motion/react"

function launch() {
    return (
        <>
            <div className="flex flex-col items-center justify-center min-h-screen">
                <motion.img
                    className="size-24 cursor-pointer "
                    src="src/assets/react.svg"
                    alt=""
                    whileHover={{ rotate: 360 }}
                    transition={{ duration: 1 }}
                />
                <h1>Launch</h1>
            </div>
        </>
    )
}

export default launch