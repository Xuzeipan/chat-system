import {Outlet} from "react-router"
import TimeBar from '../Header/timeBar'

function AppLayout() {

  return (
    <>
        <TimeBar />
        <Outlet />
    </>
  )
}

export default AppLayout